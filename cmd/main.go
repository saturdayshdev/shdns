package main

import (
	"log"
	"os"
	lib "shdns/internal"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
)

func main() {
	log.Println("Starting shDNS...")

	docker, err := lib.InitDockerClient()
	if err != nil {
		panic(err)
	}

	cloudflare, err := lib.InitCloudflareClient(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_EMAIL"))
	if err != nil {
		panic(err)
	}

	filters := types.EventsOptions{
		Filters: filters.NewArgs(
			filters.Arg("type", "container"),
			filters.Arg("event", "start"),
		),
	}

	docker.Events(filters, func(event events.Message) {
		labels, err := docker.GetContainerLabels(event.Actor.ID)
		if err != nil {
			log.Println(err)
			return
		}

		recType, recName := labels["shdns.type"], labels["shdns.name"]
		recValue, recProxied := labels["shdns.value"], labels["shdns.proxied"]
		if recType == "" || recName == "" || recValue == "" || recProxied == "" {
			return
		}
		
		err = cloudflare.CreateDNSRecord(recType, recName, recValue, recProxied == "true")
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Created DNS record %s %s %s", recType, recName, recValue)
	})
}
