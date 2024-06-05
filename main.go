package main

import (
	"fmt"

	"github.com/gleich/tdjrl/src/config"
)

func main() {
	loadedConfig := config.Load()
	fmt.Printf("%+v\n", loadedConfig)
}
