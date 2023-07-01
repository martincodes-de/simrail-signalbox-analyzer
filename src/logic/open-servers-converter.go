package logic

import (
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types/simrail-api/responses"
)

func ConvertOpenServersToServer(openServers []responses.OpenServers) []types.Server {
	var servers []types.Server

	for _, openServer := range openServers {
		//fmt.Printf("%s | %s | %s | active: %t | ID: %s\n", openServer.ServerCode, openServer.ServerName, openServer.ServerRegion, openServer.IsActive, openServer.Id)
		servers = append(servers, types.Server{
			Id:        openServer.Id,
			IsActive:  openServer.IsActive,
			Name:      openServer.ServerName,
			Region:    openServer.ServerRegion,
			Shortname: openServer.ServerCode,
		})
	}

	return servers
}
