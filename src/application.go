package src

import (
	"encoding/json"
	"fmt"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/logic"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/query/simrail-api"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types"
	"log"
	"os"
	"time"
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

	existingEntries := make(map[string][]types.Server)
	existingEntriesInBytes, readExistingDatabaseErr := os.ReadFile("database/db.json")
	if readExistingDatabaseErr != nil {
		log.Fatal("Cant read existing database", readExistingDatabaseErr)
	}

	if len(existingEntriesInBytes) > 0 {
		decodeExistingEntriesErr := json.Unmarshal(existingEntriesInBytes, &existingEntries)
		if decodeExistingEntriesErr != nil {
			log.Fatal("Cant decode existing database", decodeExistingEntriesErr)
		}
	}

	newEntry := map[string][]types.Server{
		time.Now().String(): servers,
	}

	for date, servers := range existingEntries {
		newEntry[date] = servers
	}

	newFileContent, encodeEntriesErr := json.Marshal(newEntry)
	if encodeEntriesErr != nil {
		log.Fatal("Cant encode changed database", encodeEntriesErr)
	}

	writerError := os.WriteFile("database/db.json", newFileContent, 0644)
	if writerError != nil {
		log.Fatal("Cant save changed database", writerError)
	}
}
