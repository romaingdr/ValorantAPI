package backend

type UniqueMap struct {
	Datas UniqueMapData `json:"data"`
}

type UniqueMapData struct {
	UUID                 string          `json:"uuid"`
	DisplayName          string          `json:"displayName"`
	NarrativeDescription string          `json:"narrativeDescription"`
	TacticalDescription  string          `json:"tacticalDescription"`
	Coordinates          string          `json:"coordinates"`
	DisplayIcon          string          `json:"displayIcon"`
	Callouts             []UniqueCallout `json:"callouts"`
}

type UniqueCallout struct {
	RegionName      string         `json:"regionName"`
	SuperRegionName string         `json:"superRegionName"`
	Location        UniqueLocation `json:"location"`
}

// Location représente les coordonnées d'un point sur la carte du jeu Valorant
type UniqueLocation struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
