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

	benchstatName := fmt.Sprintf("benchstat_%s.txt", runtime.Version())
	f, err := os.Create(benchstatName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("benchstats written to '%s'\n", benchstatName)
	}()

	err = performance.Check(os.Stdout, f, performance.ProfileNone, ld, "AUTO", true, "1m")
	if err != nil {
		log.Fatal(err)
	}
}
