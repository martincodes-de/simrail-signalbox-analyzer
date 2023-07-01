package src

import (
	"fmt"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/query/simrail-api"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types"
	"log"
)

type Application struct {
}

func (a Application) Run() {
	var servers []types.Server
	//fmt.Printf("%+v\n", servers)

	fmt.Println("Application started.")

	serversFromServersOpenResponse, err := simrail_api.OpenServersQuery()
	if err != nil {
		log.Fatal(err)
	}

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
