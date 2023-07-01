package logic

import (
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types"
	"github.com/martincodes-de/simrail-signalbox-analyzer/src/types/simrail-api/responses"
)

func ConvertSignalboxesByResponseToSignalboxes(signalboxesFromResponse []responses.Signalbox) []types.Signalbox {
	var signalboxes []types.Signalbox

	for _, signalbox := range signalboxesFromResponse {
		var dispatcherSteamId = signalbox.DispatchedBy[0].SteamId
		isDispatchedByPlayer := false

		if dispatcherSteamId != "" {
			isDispatchedByPlayer = true
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
