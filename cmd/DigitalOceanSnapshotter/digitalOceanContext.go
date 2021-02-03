package main

import (
	"context"

	"github.com/digitalocean/godo"
)

// DigitalOceanContext is an helper struct to acess Digital Ocean actions
type DigitalOceanContext struct {
	client *godo.Client
	ctx    context.Context
}

// GetVolume gets Volume by na
func (d DigitalOceanContext) GetVolume(id string) (*godo.Volume, *godo.Response, error) {
	return d.client.Storage.GetVolume(d.ctx, id)
}

// CreateSnapshot creates a snapshot from a Volume
func (d DigitalOceanContext) CreateSnapshot(options *godo.SnapshotCreateRequest) (*godo.Snapshot, *godo.Response, error) {
	return d.client.Storage.CreateSnapshot(d.ctx, options)
}

// ListSnapshots lists all snapshots from a volume
func (d DigitalOceanContext) ListSnapshots(volumeID string, opts *godo.ListOptions) ([]godo.Snapshot, *godo.Response, error) {
	return d.client.Storage.ListSnapshots(d.ctx, volumeID, opts)
}

// DeleteSnapshot Deletes a snaphot by id
func (d DigitalOceanContext) DeleteSnapshot(id string) (*godo.Response, error) {
	return d.client.Storage.DeleteSnapshot(d.ctx, id)
}
