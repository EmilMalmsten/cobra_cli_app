/*
Copyright © 2023 Emil M
*/
package main

import (
	"fmt"

	"github.com/limesten/cobra_cli_app/cmd"
)

var GoBlueprintVersion string

func main() {
	cmd.Execute()
	fmt.Printf("version iz: %s", GoBlueprintVersion)
}
