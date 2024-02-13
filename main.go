package main

import (
	"flag"
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
	// arguments and definitions
	var optBenchstat bool
	var optProfile string
	var optDuration string
	var optUncapped bool

	flgs := flag.NewFlagSet("gopher2600_performance_profiles", flag.ExitOnError)
	flgs.BoolVar(&optBenchstat, "benchstat", false, "record benchstats to file")
	flgs.StringVar(&optProfile, "profile", "none", "run with profiling: CPU, MEM, TRACE, ALL")
	flgs.StringVar(&optDuration, "duration", "1m", "duration of execution")
	flgs.BoolVar(&optUncapped, "uncapped", true, "run emulation with no FPS cap")

	// parse arguments
	err := flgs.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	args := flgs.Args()

	// how we load cartridge depends on how many arguments are left over on the
	// command line. if no argument has been specified then we use the embedded
	// cartridge
	var loader cartridgeloader.Loader

	switch len(args) {
	case 0:
		loader, err = cartridgeloader.NewLoaderFromEmbed("cartridge.bin", cartridge, "AUTO")
		if err != nil {
			log.Fatal(err)
		}
	case 1:
		loader, err = cartridgeloader.NewLoader(args[0], "AUTO")
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("too many arguments")
	}

	// output some basic system information before proceeding
	fmt.Println(runtime.Version())
	fmt.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("running for %v", optDuration)
	if optUncapped {
		fmt.Println(" (uncapped)")
	} else {
		fmt.Println(" (capped)")
	}

	// benchstatFile is nil if the benchstat option has not been enabled. that's
	// okay because performance.Check() expects that the argument might be nil
	var benchstatFile *os.File

	if optBenchstat {
		benchstatName := fmt.Sprintf("benchstat_%s.txt", runtime.Version())
		benchstatFile, err = os.Create(benchstatName)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err := benchstatFile.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
		fmt.Printf("writing benchstats to '%s'\n", benchstatName)
	}

	// parse profile option
	profile, err := performance.ParseProfileString(optProfile)
	if err != nil {
		log.Fatal(err)
	}
	if profile != performance.ProfileNone {
		fmt.Println("creating profile")
	}

	// performance check with os.Stdout for user feedback
	err = performance.Check(os.Stdout, benchstatFile, profile, loader, "AUTO", optUncapped, optDuration)
	if err != nil {
		log.Fatal(err)
	}
}
