package src

import (
	"fmt"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/simrail-api/fetcher"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types"
	"log"
)

type Application struct {
}

func (a Application) Run() {
	fmt.Println("Application started.")

	serversFromServersOpenResponse, err := fetcher.MakeServersOpenHttpRequest()
	if err != nil {
		log.Fatal(err)
	}

	var servers []types.Server

	for _, server := range serversFromServersOpenResponse {
		//fmt.Printf("%s | %s | %s | active: %t | ID: %s\n", server.ServerCode, server.ServerName, server.ServerRegion, server.IsActive, server.Id)
		servers = append(servers, types.Server{
			Id:        server.Id,
			IsActive:  server.IsActive,
			Name:      server.ServerName,
			Region:    server.ServerRegion,
			Shortname: server.ServerCode,
		})
	}
}
