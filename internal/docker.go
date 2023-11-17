package lib

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	API *client.Client
}

func (c* DockerClient) Events (opts types.EventsOptions, callback func(event events.Message)) error {
	eventsCh, errorCh := c.API.Events(context.Background(), opts)
	for {
		select {
		case event := <-eventsCh:
			go callback(event)
		case err := <-errorCh:
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (c* DockerClient) GetContainerLabels (containerId string) (map[string]string, error) {
	container, err := c.API.ContainerInspect(context.Background(), containerId)
	if err != nil {
		return nil, err
	}
	
	return container.Config.Labels, nil
}

func InitDockerClient() (*DockerClient, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	return &DockerClient{API: client}, nil
}