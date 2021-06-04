package color

type Color int

const (

	// Info is a light blue color
	// to represent informational level text
	Info Color = iota

	// Error is a bright red color
	// for error messages
	Error

	// Success is a bright green color
	// for success messages
	Success

	// Warn is a bright yellow color
	// for indicating a potential issue
	Warn
)

func (c Color) Hex() string {
	return [...]string{"#00FFFF", "#B20000", "#198C19", "#FFFF00"}[c]
}

type RGB struct {
	Red   string
	Green string
	Blue  string
}

func (c Color) RGB() RGB {
	return [...]RGB{
		{
			Red:   "0",
			Green: "255",
			Blue:  "255",
		},
		{
			Red:   "178",
			Green: "0",
			Blue:  "0",
		},
		{
			Red:   "178",
			Green: "0",
			Blue:  "0",
		},
		{
			Red:   "25",
			Green: "140",
			Blue:  "25",
		},
		{
			Red:   "255",
			Green: "255",
			Blue:  "0",
		},
	}[c]
}
