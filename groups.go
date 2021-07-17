package rivers

// GaugeGroup represents a group of measuring stations
// located in the geographical area of Ireland.
type GaugeGroup int

const (
	Nore GaugeGroup = iota + 1
	Shannon
	Turlough
	Barrow
	MunsterBlackwater
	SuirBackUp
	_
	Erne
	Corrib
	Moy
	Fergus
	Maigue
	Slaney
	ShannonLRee
	Suck
	Tidal
	Boyne
	MunsterBlackwaterMallow
	MunsterBlackwaterFermoy
	Inny
	Brosna
	Foyle
	Bandon
	Laune
	Ballysadare
	Suir
	WaterfordCity
	SouthGalway
)

// StationGroups represents measurements stations grouped by geographical area.
type StationGroups struct {
	Groups []StationGroup `json:""`
}

// StationGroup represents a group of stations located in a geographical area.
type StationGroup struct {
	Name string `json:"group_name"`
	ID   int    `json:"group_id"`
}
