package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/tnoda78/posiserver"
)

func main() {
	fmt.Println("[posiserver] strat.")

	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error:%s\n", err)
			os.Exit(1)
		}
	}()

	os.Exit(_main())
}

func _main() int {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(1)
	}

	server := posiserver.NewPosiserver()
	server.Run()

	return 0
}
