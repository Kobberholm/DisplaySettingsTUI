package vcs

import (
	"runtime/debug"

	"github.com/charmbracelet/log"
)

func GetCommitHash() string {
	gitCommit := ""

	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Error("could not read build info")
		return ""
	}
	modified := false
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			gitCommit = setting.Value[0:5]
		case "vcs.modified":
			if setting.Value == "true" {
				modified = true
			}
		}
	}
	if modified {
		gitCommit += "+DIRTY"
	}

	return gitCommit
}
