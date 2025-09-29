package main

import (
	"DisplaySettingsTUI/root"
	"DisplaySettingsTUI/vcs"
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func main() {
	v := flag.Bool("v", false, "Show version information")

	flag.Parse()

	if *v {
		fmt.Printf("Version: %s", vcs.GetCommitHash())
		os.Exit(0)
	}

	f, err := tea.LogToFile("/tmp/debug-display-settings-tui.log", "Tea.Debug->")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Info("starting display settings program...")

	m := root.New()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
