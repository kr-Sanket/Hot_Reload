package builder

import (
	"log"
	"os/exec"
)

type Builder struct {
	BuildCmd string
}

func New(buildCmd string) *Builder {
	return &Builder{
		BuildCmd: buildCmd,
	}
}

func (b *Builder) Build() error {

	log.Println("Starting build...")
	log.Println("Running:", b.BuildCmd)

	cmd := exec.Command("cmd", "/C", b.BuildCmd)

	output, err := cmd.CombinedOutput()

	if len(output) > 0 {
		log.Println("Build output:")
		log.Println(string(output))
	}

	if err != nil {
		log.Println("Build failed:", err)
		return err
	}

	log.Println("Build successful")

	return nil
}