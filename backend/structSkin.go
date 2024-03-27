package backend

type WeaponSkinData struct {
	Status int        `json:"status"`
	Data   WeaponSkin `json:"data"`
}

type WeaponSkin struct {
	UUID            string       `json:"uuid"`
	DisplayName     string       `json:"displayName"`
	ThemeUUID       string       `json:"themeUuid"`
	ContentTierUUID string       `json:"contentTierUuid"`
	DisplayIcon     string       `json:"displayIcon"`
	Wallpaper       string       `json:"wallpaper"`
	AssetPath       string       `json:"assetPath"`
	Chromas         []SkinChroma `json:"chromas"`
	Levels          []SkinLevel  `json:"levels"`
}

type SkinChroma struct {
	UUID          string  `json:"uuid"`
	DisplayName   string  `json:"displayName"`
	DisplayIcon   string  `json:"displayIcon"`
	FullRender    string  `json:"fullRender"`
	Swatch        *string `json:"swatch,omitempty"`
	StreamedVideo *string `json:"streamedVideo,omitempty"`
	AssetPath     string  `json:"assetPath"`
}

type SkinLevel struct {
	UUID          string  `json:"uuid"`
	DisplayName   string  `json:"displayName"`
	LevelItem     *string `json:"levelItem,omitempty"`
	DisplayIcon   string  `json:"displayIcon"`
	StreamedVideo *string `json:"streamedVideo,omitempty"`
	AssetPath     string  `json:"assetPath"`
}
