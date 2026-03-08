package process

import (
	"log"
	"os"
	"os/exec"
)

type Manager struct {
	cmd     *exec.Cmd
	execCmd string
}

func New(execCmd string) *Manager {
	return &Manager{
		execCmd: execCmd,
	}
}

func (p *Manager) Start() error {

	log.Println("Starting server...")

	cmd := exec.Command("cmd", "/C", p.execCmd)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	p.cmd = cmd

	go func() {
		cmd.Wait()
		log.Println("Server process exited")
	}()

	return nil
}

func (p *Manager) Stop() {

	if p.cmd == nil || p.cmd.Process == nil {
		return
	}

	log.Println("Stopping server...")

	err := p.cmd.Process.Kill()
	if err != nil {
		log.Println("Server already stopped")
	}

	p.cmd.Wait()

	p.cmd = nil
}

func (p *Manager) Restart() error {

	p.Stop()

	return p.Start()
}