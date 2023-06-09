package responses

type OpenServers struct {
	ServerCode   string `json:"ServerCode"`
	ServerName   string `json:"ServerName"`
	ServerRegion string `json:"ServerRegion"`
	IsActive     bool   `json:"IsActive"`
	Id           string `json:"id"`
}

type ServersOpenResponse struct {
	Result      bool          `json:"result"`
	Data        []OpenServers `json:"data"`
	Count       int           `json:"count"`
	Description string        `json:"description"`
}
