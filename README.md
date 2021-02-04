# DigitalOceanSnapshotter
[![Docker Pulls](https://img.shields.io/docker/pulls/topgg/digitaloceansnapshotter.svg)](https://hub.docker.com/r/topgg/digitaloceansnapshotter/)

Simple Digital Ocean Volumes Backup Service using Snapshots

## Getting Started

This Service is supposed to be run as a CronJob every x hours, The recommended way is using the Docker image from `topgg/digitaloceansnapshotter`.

It requires following enviroment variables to be set:

`DO_TOKEN` - Digital Ocean Access token

`DO_VOLUMES` - List of Digital Ocean volume ids seperated by a comma

`DO_SNAPSHOT_COUNT` - Amount of snapshots to keep for each volume

### Prerequisites

Go 1.15

## Deployment

This was designed for usage in Kubernetes but it can technically be run anywhere, an example Kubernetes CronJob manifest running every day could look like this

```yml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: DigitalOceanSnapshotter
spec:
  schedule: "0 0 * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: DigitalOceanSnapshotter
              image: topgg/digitaloceansnapshotter
              envFrom:
                - configMapRef:
                    name: your-config-map-name
          restartPolicy: OnFailure
```

**Note that this uses the latest image which you shouldn't do in prod, use a tag version or commit hash instead**
## Built With

* [Godo](https://github.com/digitalocean/godo) - DigitalOcean V2 API client library
* [Logrus](https://github.com/sirupsen/logrus) - Logging framework used
* [Slack-Go](https://github.com/slack-go/slack) - Slack API client library

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/top-gg/DigitalOceanSnapshotter/tags). 

## Authors

* **DevYukine @ Top.gg** - *Initial work* - [DevYukine](https://github.com/DevYukine)

See also the list of [contributors](https://github.com/top-gg/DigitalOceanSnapshotter/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details