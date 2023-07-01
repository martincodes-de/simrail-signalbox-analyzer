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

	openServers, err := simrail_api.OpenServersQuery()
	if err != nil {
		log.Fatal(err)
	}
	servers = logic.ConvertOpenServersToServer(openServers)

	var firstServer = servers[0]

}
