package responses

type Signalbox struct {
	Name                string  `json:"Name"`
	Prefix              string  `json:"Prefix"`
	DifficultyLevel     int     `json:"DifficultyLevel"`
	Latititude          float64 `json:"Latititude"`
	Longitude           float64 `json:"Longitude"`
	MainImageURL        string  `json:"MainImageURL"`
	AdditionalImage1URL string  `json:"AdditionalImage1URL"`
	AdditionalImage2URL string  `json:"AdditionalImage2URL"`
	DispatchedBy        []struct {
		ServerCode string `json:"ServerCode"`
		SteamId    string `json:"SteamId"`
	} `json:"DispatchedBy"`
	Id string `json:"id"`
}

type SignalboxesOpenForServerResponse struct {
	Result      bool        `json:"result"`
	Data        []Signalbox `json:"data"`
	Count       int         `json:"count"`
	Description string      `json:"description"`
}
