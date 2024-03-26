package backend

type Role struct {
	UUID        string `json:"uuid"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	DisplayIcon string `json:"displayIcon"`
	AssetPath   string `json:"assetPath"`
}

type Ability struct {
	Slot        string `json:"slot"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	DisplayIcon string `json:"displayIcon"`
}

type Character struct {
	UUID          string    `json:"uuid"`
	DisplayName   string    `json:"displayName"`
	Description   string    `json:"description"`
	DeveloperName string    `json:"developerName"`
	DisplayIcon   string    `json:"displayIcon"`
	BustPortrait  string    `json:"bustPortrait"`
	FullPortrait  string    `json:"fullPortrait"`
	Role          Role      `json:"role"`
	Abilities     []Ability `json:"abilities"`
}

type AgentResponse struct {
	Status int         `json:"status"`
	Data   []Character `json:"data"`
}
