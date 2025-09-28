package display

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

// Display represents information about a detected display
// Claude generated code. Model Opus 4.1
type Display struct {
	Index     int    // Display number (1, 2, etc.)
	connector string // DRM connector (e.g., "card1-DP-1")
	serial    string // serial number (e.g., "93QSX34")
}

func (d *Display) FilterValue() string {
	return d.connector
}

func (d *Display) Title() string {
	return d.connector
}

func (d *Display) Description() string {
	return d.serial
}

// DetectDisplays runs ddcutil detect and returns parsed display information
func DetectDisplays() ([]Display, error) {
	// Execute ddcutil detect command
	cmd := exec.Command("ddcutil", "detect")
	output, err := cmd.Output()
	if err != nil {
		log.Error("Failed to detect displays: %v", err)
		return nil, fmt.Errorf("failed to run ddcutil detect: %w", err)
	}

	return parseDisplayOutput(output)
}

// parseDisplayOutput parses the ddcutil output into Display structs
func parseDisplayOutput(output []byte) ([]Display, error) {
	var displays []Display
	scanner := bufio.NewScanner(bytes.NewReader(output))

	// Regular expressions for parsing
	displayRe := regexp.MustCompile(`^Display\s+(\d+)`)
	connectorRe := regexp.MustCompile(`^\s+DRM connector:\s+(.+)`)
	serialRe := regexp.MustCompile(`^\s+Serial number:\s+(.+)`)

	var currentDisplay *Display

	for scanner.Scan() {
		line := scanner.Text()

		// Check for a new display section
		if matches := displayRe.FindStringSubmatch(line); matches != nil {
			// Save previous display if exists
			if currentDisplay != nil {
				displays = append(displays, *currentDisplay)
			}

			// Start a new display
			index, _ := strconv.Atoi(matches[1])
			currentDisplay = &Display{
				Index: index,
			}
		} else if currentDisplay != nil {
			// Parse connector
			if matches = connectorRe.FindStringSubmatch(line); matches != nil {
				currentDisplay.connector = strings.TrimSpace(matches[1])
			}
			// Parse serial number
			if matches = serialRe.FindStringSubmatch(line); matches != nil {
				currentDisplay.serial = strings.TrimSpace(matches[1])
			}
		}
	}

	// Remember the last display
	if currentDisplay != nil {
		displays = append(displays, *currentDisplay)
	}

	if err := scanner.Err(); err != nil {
		log.Error("Scanner error: %v", err)
		return nil, fmt.Errorf("error scanning output: %w", err)
	}

	return displays, nil
}
