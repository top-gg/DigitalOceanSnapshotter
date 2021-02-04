package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/digitalocean/godo"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

const createdAtFormat = "2006-01-02T15:04:05Z"

type snapshotterContext struct {
	DoContext    *DigitalOceanContext
	SlackContext *SlackContext
}

func initLogging() {
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
	})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)
}

func main() {
	initLogging()

	DOToken, present := os.LookupEnv("DO_TOKEN")

	if present == false {
		log.Fatal("Missing enviroment variable \"DO_TOKEN\"")
	}

	volumesEnv, present := os.LookupEnv("DO_VOLUMES")

	if present == false {
		log.Fatal("Missing enviroment variable \"DO_VOLUMES\"")
	}

	snapshotCountEnv, present := os.LookupEnv("DO_SNAPSHOT_COUNT")

	if present == false {
		log.Fatal("Missing enviroment variable \"DO_SNAPSHOT_COUNT\"")
	}

	snapshotCount, err := strconv.Atoi(snapshotCountEnv)

	if err != nil {
		log.Fatal("Enviroment variable \"DO_SNAPSHOT_COUNT\" is not an integer")
	}

	slackEnv := os.Getenv("SLACK_TOKEN")

	var slackContext *SlackContext = nil

	if slackEnv != "" {
		channelID, present := os.LookupEnv("SLACK_CHANNEL_ID")

		if present == false {
			log.Fatal("Missing enviroment variable \"SLACK_CHANNEL_ID\"")
		}

		slackContext = &SlackContext{
			client:    slack.New(slackEnv),
			channelID: channelID,
		}
	}

	ctx := snapshotterContext{
		DoContext: &DigitalOceanContext{
			client: godo.NewFromToken(DOToken),
			ctx:    context.TODO(),
		},
		SlackContext: slackContext,
	}

	volumeIDs := strings.Split(volumesEnv, ",")

	for _, volumeID := range volumeIDs {
		volume, _, err := ctx.DoContext.GetVolume(volumeID)
		if err != nil {
			handleError(ctx, err, true)
			return
		}

		_, _, err = ctx.DoContext.CreateSnapshot(&godo.SnapshotCreateRequest{
			VolumeID: volume.ID,
			Name:     time.Now().Format("2006-01-02T15:04:05"),
		})
		if err != nil {
			handleError(ctx, err, true)
			return
		}

		snapshots, _, err := ctx.DoContext.ListSnapshots(volume.ID, nil)

		snapshotLength := len(snapshots)

		if snapshotLength > snapshotCount {
			sort.SliceStable(snapshots, func(firstIndex, secondIndex int) bool {
				firstTime, err := time.Parse(snapshots[firstIndex].Created, createdAtFormat)
				if err != nil {
					handleError(ctx, err, true)
				}

				secondTime, err := time.Parse(snapshots[firstIndex].Created, createdAtFormat)
				if err != nil {
					handleError(ctx, err, true)
				}

				return firstTime.Before(secondTime)
			})

			snapshotsToDelete := snapshotLength - snapshotCount

			for i := 0; i < snapshotsToDelete; i++ {
				_, err := ctx.DoContext.DeleteSnapshot(snapshots[i].ID)

				if err != nil {
					handleError(ctx, err, false)
					return
				}
			}
		}
	}
}

func handleError(ctx snapshotterContext, err error, fatal bool) {
	errString := err.Error()

	if ctx.SlackContext != nil {
		err = ctx.SlackContext.SendEvent(errString, log.ErrorLevel)
		if err != nil {
			log.Error(fmt.Sprintf("Error while trying to send error to Slack %s", err.Error()))
		}
	}

	if fatal {
		log.Fatal(errString)
	}

	log.Error(errString)
}
