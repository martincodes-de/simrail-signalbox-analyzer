package src

import (
	"fmt"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/command"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/logic"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/query/simrail-api"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types"
	"log"
)

type Application struct {
}

func (a Application) Run() {
	var servers []types.Server
	//fmt.Printf("%+v\n", servers)

	fmt.Println("Application started. Fetching servers ...")

	openServers, openServerErr := simrail_api.OpenServersQuery()
	if openServerErr != nil {
		log.Fatal(openServerErr)
	}
	servers = logic.ConvertOpenServersToServer(openServers)

	fmt.Println("Add signalboxes to server ...")

	for index := range servers {
		server := &servers[index]
		signalboxes, signalboxErr := simrail_api.SignalboxesForServerQuery(server.Shortname)
		if signalboxErr != nil {
			log.Fatal(openServerErr)
		}
		convertedSignalBoxes := logic.ConvertSignalboxesByResponseToSignalboxes(signalboxes)
		server.Signalboxes = convertedSignalBoxes
	}

	fmt.Println("Update database ...")

	command.SaveNewEntries(servers)

	fmt.Println("Done")
}
