package command

import (
	"encoding/json"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types"
	"log"
	"os"
	"time"
)

func SaveNewEntries(servers []types.Server) {
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
