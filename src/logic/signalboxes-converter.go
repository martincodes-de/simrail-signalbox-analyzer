package logic

import (
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types/simrail-api/responses"
)

func ConvertSignalboxesByResponseToSignalboxes(signalboxesFromResponse []responses.Signalbox) []types.Signalbox {
	var signalboxes []types.Signalbox

	for _, signalbox := range signalboxesFromResponse {
		isDispatchedByPlayer := false
		dispatcherSteamId := ""

		if len(signalbox.DispatchedBy) == 1 {
			isDispatchedByPlayer = true
			dispatcherSteamId = signalbox.DispatchedBy[0].SteamId
		}

		signalboxes = append(signalboxes, types.Signalbox{
			Id:                   signalbox.Id,
			IsDispatchedByPlayer: isDispatchedByPlayer,
			DispatcherSteamId:    dispatcherSteamId,
			DifficultyLevel:      signalbox.DifficultyLevel,
			Name:                 signalbox.Name,
			Shortname:            signalbox.Prefix,
		})
	}

	return signalboxes
}
