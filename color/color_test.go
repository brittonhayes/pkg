package color_test

import (
	"fmt"

	"github.com/brittonhayes/pkg/color"
)

// Get a hex code for info color
func ExampleColor_String() {
	// Setup the HEX code of the cyan info
	// color
	info := color.Info.String()

	// Print out the string
	// of the color
	fmt.Println(info)
}
