package models

var SceneModel = map[int64]string{
	1:    "Ocean",
	2:    "Romance",
	3:    "Sunset",
	4:    "Party",
	5:    "Fireplace",
	6:    "Cozy",
	7:    "Forest",
	8:    "Pastel Colors",
	9:    "Wake up",
	10:   "Bedtime",
	11:   "Warm White",
	12:   "Daylight",
	13:   "Cool white",
	14:   "Night light",
	15:   "Focus",
	16:   "Relax",
	17:   "True colors",
	18:   "TV time",
	19:   "Plantgrowth",
	20:   "Spring",
	21:   "Summer",
	22:   "Fall",
	23:   "Deepdive",
	24:   "Jungle",
	25:   "Mojito",
	26:   "Club",
	27:   "Christmas",
	28:   "Halloween",
	29:   "Candlelight",
	30:   "Golden white",
	31:   "Pulse",
	32:   "Steampunk",
	1000: "Rhythm",
}

type ParamPayload struct {
	State     bool    `json:"state"`
	WarmWhite float64 `json:"w,omitempty"`       // value > 0 < 256
	ColdWhite float64 `json:"c,omitempty"`       // value > 0 < 256
	Speed     int64   `json:"speed,omitempty"`   // value > 0 < 101
	SceneId   int64   `json:"sceneId,omitempty"` // SceneModel
	R         float64 `json:"r,omitempty"`
	G         float64 `json:"g,omitempty"`
	B         float64 `json:"b,omitempty"`
	Dimming   int64   `json:"dimming,omitempty"`
	ColorTemp float64 `json:"temp,omitempty"`
}

type ResultPayload struct {
	Success     bool    `json:"success,omitempty"`
	Mac         string  `json:"mac,omitempty"`
	Rssi        int64   `json:"rssi,omitempty"`
	Src         string  `json:"src,omitempty"`
	State       bool    `json:"state,omitempty"`
	SceneId     int64   `json:"sceneId,omitempty"`
	Speed       int64   `json:"speed,omitempty"`
	Temp        int64   `json:"temp,omitempty"`
	Dimming     int64   `json:"dimming,omitempty"`
	HomeId      int64   `json:"homeId,omitempty"`
	RoomId      int64   `json:"roomId,omitempty"`
	HomeLock    bool    `json:"homeLock,omitempty"`
	PairingLock bool    `json:"pairingLock,omitempty"`
	TypeId      int64   `json:"typeId,omitempty"`
	ModuleName  string  `json:"module_name,omitempty"`
	FwVersion   string  `json:"fwVersion,omitempty"`
	GroupId     int64   `json:"groupId,omitempty"`
	DrvConf     []int64 `json:"drvConf,omitempty"`
}

type RequestPayload struct {
	Method string       `json:"method,omitempty"`
	Params ParamPayload `json:"params,omitempty"`
}

type ResponsePayload struct {
	Method string        `json:"method,omitempty"`
	Env    string        `json:"env,omitempty"`
	Result ResultPayload `json:"result,omitempty"`
}
