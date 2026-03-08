package process

import (
	"log"
	"os"
	"os/exec"
	"time"
)

type Manager struct {
	cmd       *exec.Cmd
	execCmd   string
	lastStart time.Time
}

func New(execCmd string) *Manager {
	return &Manager{
		execCmd: execCmd,
	}
}

func (p *Manager) Start() error {

	// Prevent rapid restart loops
	if time.Since(p.lastStart) < 2*time.Second {
		log.Println("Restart cooldown triggered")
		time.Sleep(2 * time.Second)
	}

	log.Println("Starting server...")

	cmd := exec.Command(p.execCmd)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	p.cmd = cmd
	p.lastStart = time.Now()

	return nil
}

func (p *Manager) Stop() {

	if p.cmd == nil {
		return
	}

	log.Println("Stopping server...")

	err := p.cmd.Process.Kill()
	if err != nil {
		log.Println("Server already stopped")
	}

	// Wait for process termination
	p.cmd.Wait()

	p.cmd = nil
}

func (p *Manager) Restart() error {

	p.Stop()

	return p.Start()
}