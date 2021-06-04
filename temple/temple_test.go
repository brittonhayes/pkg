package temple_test

import (
	"fmt"

	"github.com/brittonhayes/pkg/temple"
)

func ExampleRender() {
	// Setup a template string and
	// interface for rendering
	templateString := "{{.Value}}"
	data := map[string]interface{}{
		"Value": "foo",
	}

	// Render the template
	output, err := temple.Render(templateString, data)
	if err != nil {
		return
	}

	// Print it out
	fmt.Println(output)
}
