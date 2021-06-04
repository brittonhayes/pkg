package color_test

import (
	"fmt"

	"github.com/brittonhayes/pkg/color"
)

// Get a hex code for info color
func ExampleColor_Hex_info() {

	// Setup the HEX code of the cyan info
	// color
	info := color.Info.Hex()

	// Print out the string
	// of the color
	fmt.Println(info)
}

// Get an RGB format for the error color
func ExampleColor_RGB_error() {

	// Setup the RGB of the red error
	// color
	errorColor := color.Error.RGB()

	// Print out the struct
	// of the color
	fmt.Println(errorColor)
}

