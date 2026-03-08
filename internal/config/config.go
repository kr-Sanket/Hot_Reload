package config

import (
	"flag"
	"fmt"
)

type Config struct {
	Root  string
	Build string
	Exec  string
}

func Parse() (*Config, error) {

	root := flag.String("root", "", "project root directory")
	build := flag.String("build", "", "build command")
	execCmd := flag.String("exec", "", "run command")

	flag.Parse()

	if *root == "" || *build == "" || *execCmd == "" {
		return nil, fmt.Errorf("all flags --root, --build, and --exec are required")
	}

	cfg := &Config{
		Root:  *root,
		Build: *build,
		Exec:  *execCmd,
	}

	return cfg, nil
}