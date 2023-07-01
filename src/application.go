package src

import (
	"fmt"
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

	fmt.Println("Application started.")

	openServers, openServerErr := simrail_api.OpenServersQuery()
	if openServerErr != nil {
		log.Fatal(openServerErr)
	}
	servers = logic.ConvertOpenServersToServer(openServers)

	var firstServer = servers[0]
	signalboxes, signalboxErr := simrail_api.SignalboxesForServerQuery(firstServer.Shortname)
	if signalboxErr != nil {
		log.Fatal(openServerErr)
	}

	convertedSignalBoxes := logic.ConvertSignalboxesByResponseToSignalboxes(signalboxes)

	fmt.Printf("%+v\n", convertedSignalBoxes)
}
