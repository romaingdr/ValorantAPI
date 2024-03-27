package backend

type Weapons struct {
	Status int    `json:"status"`
	Data   []Item `json:"data"`
}

type Item struct {
	UUID            string      `json:"uuid"`
	DisplayName     string      `json:"displayName"`
	Category        string      `json:"category"`
	DefaultSkinUUID string      `json:"defaultSkinUuid"`
	DisplayIcon     string      `json:"displayIcon"`
	KillStreamIcon  string      `json:"killStreamIcon"`
	AssetPath       string      `json:"assetPath"`
	WeaponStats     WeaponStats `json:"weaponStats"`
	ShopData        ShopData    `json:"shopData"`
	Skins           []Skin      `json:"skins"`
}

type WeaponStats struct {
	FireRate            float64       `json:"fireRate"`
	MagazineSize        int           `json:"magazineSize"`
	RunSpeedMultiplier  float64       `json:"runSpeedMultiplier"`
	EquipTimeSeconds    float64       `json:"equipTimeSeconds"`
	ReloadTimeSeconds   float64       `json:"reloadTimeSeconds"`
	FirstBulletAccuracy float64       `json:"firstBulletAccuracy"`
	ShotgunPelletCount  int           `json:"shotgunPelletCount"`
	WallPenetration     string        `json:"wallPenetration"`
	Feature             string        `json:"feature"`
	FireMode            string        `json:"fireMode"`
	AltFireType         string        `json:"altFireType"`
	ADSStats            ADSStats      `json:"adsStats"`
	AltShotgunStats     interface{}   `json:"altShotgunStats"`
	AirBurstStats       interface{}   `json:"airBurstStats"`
	DamageRanges        []DamageRange `json:"damageRanges"`
}

type ADSStats struct {
	ZoomMultiplier      float64 `json:"zoomMultiplier"`
	FireRate            float64 `json:"fireRate"`
	RunSpeedMultiplier  float64 `json:"runSpeedMultiplier"`
	BurstCount          int     `json:"burstCount"`
	FirstBulletAccuracy float64 `json:"firstBulletAccuracy"`
}

type DamageRange struct {
	RangeStartMeters int     `json:"rangeStartMeters"`
	RangeEndMeters   int     `json:"rangeEndMeters"`
	HeadDamage       float64 `json:"headDamage"`
	BodyDamage       float64 `json:"bodyDamage"`
	LegDamage        float64 `json:"legDamage"`
}

type ShopData struct {
	Cost              int          `json:"cost"`
	Category          string       `json:"category"`
	ShopOrderPriority int          `json:"shopOrderPriority"`
	CategoryText      string       `json:"categoryText"`
	GridPosition      GridPosition `json:"gridPosition"`
	CanBeTrashed      bool         `json:"canBeTrashed"`
	Image             interface{}  `json:"image"`
	NewImage          string       `json:"newImage"`
	NewImage2         interface{}  `json:"newImage2"`
	AssetPath         string       `json:"assetPath"`
}

type GridPosition struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type Skin struct {
	UUID            string      `json:"uuid"`
	DisplayName     string      `json:"displayName"`
	ThemeUUID       string      `json:"themeUuid"`
	ContentTierUUID string      `json:"contentTierUuid"`
	DisplayIcon     string      `json:"displayIcon"`
	Wallpaper       interface{} `json:"wallpaper"`
	AssetPath       string      `json:"assetPath"`
	Chromas         []Chroma    `json:"chromas"`
	Levels          []Level     `json:"levels"`
}

type Chroma struct {
	UUID          string      `json:"uuid"`
	DisplayName   string      `json:"displayName"`
	DisplayIcon   interface{} `json:"displayIcon"`
	FullRender    string      `json:"fullRender"`
	Swatch        interface{} `json:"swatch"`
	StreamedVideo interface{} `json:"streamedVideo"`
	AssetPath     string      `json:"assetPath"`
}

type Level struct {
	UUID          string      `json:"uuid"`
	DisplayName   string      `json:"displayName"`
	LevelItem     interface{} `json:"levelItem"`
	DisplayIcon   string      `json:"displayIcon"`
	StreamedVideo string      `json:"streamedVideo"`
	AssetPath     string      `json:"assetPath"`
}
