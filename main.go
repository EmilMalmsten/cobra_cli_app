/*
Copyright © 2023 Emil M
*/
package main

import (
	"fmt"

	"github.com/limesten/cobra_cli_app/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("version is: %s", version)
	cmd.Execute()
}
