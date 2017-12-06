package concourse

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mitchellh/colorstring"
)

func Fatal(doing string, err error) {
	Sayf(colorstring.Color("[red]error %s: %s\n"), doing, err)
	os.Exit(1)
}

func Sayf(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
}

func ReadRequest(request *OutRequest) {
	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		Fatal("Error reading request: %v\n", err)
	}
}

func WriteResponse(response OutResponse) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		Fatal("Error writing response: %v\n", err)
	}
}
