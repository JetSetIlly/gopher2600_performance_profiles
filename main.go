package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/jetsetilly/gopher2600/cartridgeloader"
	"github.com/jetsetilly/gopher2600_performance_profiles/performance"

	_ "embed"
)

//go:embed "cartridge.bin"
var cartridge []byte

func main() {
	ld, err := cartridgeloader.NewLoaderFromEmbed("cartridge.bin", cartridge, "AUTO")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(runtime.Version())
	fmt.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(0))

	err = performance.Check(os.Stdout, performance.ProfileCPU, ld, "AUTO", true, "1m")
	if err != nil {
		log.Fatal(err)
	}
}
