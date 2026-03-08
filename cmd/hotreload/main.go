package main

import (
	"fmt"
	"os"

	"github.com/kr-Sanket/hotreload/internal/config"
)

func main() {

	cfg, err := config.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("HotReload starting with config:")
	fmt.Println("Root:", cfg.Root)
	fmt.Println("Build:", cfg.Build)
	fmt.Println("Exec:", cfg.Exec)
}