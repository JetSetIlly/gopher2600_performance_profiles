package main

import (
	"log"
	"os"

	"github.com/jetsetilly/gopher2600/cartridgeloader"
	"github.com/jetsetilly/gopher2600/performance"

	_ "embed"
)

//go:embed "cartridge.bin"
var cartridge []byte

func main() {
	ld, err := cartridgeloader.NewLoaderFromEmbed("cartridge.bin", cartridge, "AUTO")
	if err != nil {
		log.Fatal(err)
	}

	err = performance.Check(os.Stdout, performance.ProfileCPU, ld, "AUTO", true, "5m")
	if err != nil {
		log.Fatal(err)
	}
}
