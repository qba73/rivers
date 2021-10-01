package rivers

// GaugeGroup represents a group of measuring stations
// located in the geographical area of Ireland.
type GaugeGroup int

/*
var groupID = map[string]int{
	"Nore":                    1,
	"Shannon":                 2,
	"Turlough":                3,
	"Barrow":                  4,
	"MunsterBlackwater":       5,
	"SuirBackUp":              6,
	"Erne":                    8,
	"Corrib":                  9,
	"Moy":                     10,
	"Fergus":                  11,
	"Maigue":                  12,
	"Slaney":                  13,
	"ShannonLRee":             14,
	"Suck":                    15,
	"Tidal":                   16,
	"Boyne":                   17,
	"MunsterBlackwaterMallow": 18,
	"MunsterBlackwaterFermoy": 19,
	"Inny":                    20,
	"Brosna":                  21,
	"Foyle":                   22,
	"Bandon":                  23,
	"Laune":                   24,
	"Ballysadare":             25,
	"Suir":                    26,
	"WaterfordCity":           27,
	"SouthGalway":             28,
}
*/

// Group represents a group of stations located in a geographical area.
type Group struct {
	Name     string    `json:"group_name"`
	ID       int       `json:"group_id"`
	Readings []Reading `json:"stations"`
}

// Reading represents a sensor reading for
// the given station name located in the
// group of sensors idenfified by name.
type Reading struct {
	Name  string
	Value float64
}
