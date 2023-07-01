package src

import (
	"encoding/json"
	"fmt"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/logic"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/query/simrail-api"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types"
	"log"
	"os"
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

	for index := range servers {
		server := &servers[index]
		signalboxes, signalboxErr := simrail_api.SignalboxesForServerQuery(server.Shortname)
		if signalboxErr != nil {
			log.Fatal(openServerErr)
		}
		convertedSignalBoxes := logic.ConvertSignalboxesByResponseToSignalboxes(signalboxes)
		server.Signalboxes = convertedSignalBoxes
	}

	fileStructure := make(map[string][]types.Server)

	data, _ := os.ReadFile("database/db.json")

	if len(data) > 0 {
		setupExistingFilestructureErr := json.Unmarshal(data, &fileStructure)
		if setupExistingFilestructureErr != nil {
			return
		}
	}

	//fmt.Printf("%+v\n", servers)

	newEntry := map[string][]types.Server{
		"tomorrow": servers,
	}

	for date, servers := range fileStructure {
		newEntry[date] = servers
	}

	jsons, _ := json.Marshal(newEntry)
	writerError := os.WriteFile("database/db.json", jsons, 0644)
	if writerError != nil {
		log.Fatal(writerError)
	}
}
