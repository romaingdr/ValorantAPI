package backend

type GameModeData struct {
	Status int        `json:"status"`
	Data   []GameMode `json:"data"`
}

type GameMode struct {
	UUID                  string                 `json:"uuid"`
	DisplayName           string                 `json:"displayName"`
	Duration              string                 `json:"duration"`
	EconomyType           *string                `json:"economyType,omitempty"`
	AllowsMatchTimeouts   bool                   `json:"allowsMatchTimeouts"`
	IsTeamVoiceAllowed    bool                   `json:"isTeamVoiceAllowed"`
	IsMinimapHidden       bool                   `json:"isMinimapHidden"`
	OrbCount              int                    `json:"orbCount"`
	RoundsPerHalf         int                    `json:"roundsPerHalf"`
	TeamRoles             []string               `json:"teamRoles,omitempty"`
	GameFeatureOverrides  []GameFeatureOverride  `json:"gameFeatureOverrides,omitempty"`
	GameRuleBoolOverrides []GameRuleBoolOverride `json:"gameRuleBoolOverrides,omitempty"`
	DisplayIcon           string                 `json:"displayIcon"`
	ListViewIconTall      *string                `json:"listViewIconTall,omitempty"`
	AssetPath             string                 `json:"assetPath"`
}

type GameFeatureOverride struct {
	FeatureName string `json:"featureName"`
	State       bool   `json:"state"`
}

type GameRuleBoolOverride struct {
	RuleName string `json:"ruleName"`
	State    bool   `json:"state"`
}
