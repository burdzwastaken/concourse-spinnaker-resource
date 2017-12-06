package main

import (
	"fmt"
	"os"
)

// this resource only supports put for now

func main() {
	fmt.Fprintf(os.Stdout, `{"version":{"ref":"none"}}`)
}
