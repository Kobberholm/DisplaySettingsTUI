package settings

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/charmbracelet/log"
)

// VCPCode represents VCP feature codes
type VCPCode int

// VCP code enums
const (
	VCPBrightness VCPCode = 0x10
	VCPContrast   VCPCode = 0x12
)

// VCPValue represents a VCP feature value
type VCPValue struct {
	Current int
	Max     int
}

// GetBrightness gets the brightness value for a display
func GetBrightness(displayNum int) (*VCPValue, error) {
	return getVCP(displayNum, VCPBrightness)
}

// GetContrast gets the contrast value for a display
func GetContrast(displayNum int) (*VCPValue, error) {
	return getVCP(displayNum, VCPContrast)
}

// SetBrightness sets the brightness value for a display
func SetBrightness(displayNum int, value int) error {
	return setVCP(displayNum, VCPBrightness, value)
}

// SetContrast sets the contrast value for a display
func SetContrast(displayNum int, value int) error {
	return setVCP(displayNum, VCPContrast, value)
}

// getVCP gets a VCP value from the display
func getVCP(displayNum int, vcpCode VCPCode) (*VCPValue, error) {
	codeStr := fmt.Sprintf("%02x", vcpCode)

	cmd := exec.Command("ddcutil", "getvcp", "-d", strconv.Itoa(displayNum), codeStr)
	output, err := cmd.Output()
	if err != nil {
		log.Error("VCP command failed", "cmd", cmd)
		return nil, fmt.Errorf("failed to get VCP 0x%02x: %w", vcpCode, err)
	}

	// Parse output: "current value = XX, maximum value = YY"
	re := regexp.MustCompile(`current value\s*=\s*(\d+),\s*maximum value\s*=\s*(\d+)`)
	matches := re.FindStringSubmatch(string(output))

	if len(matches) != 3 {
		log.Error("failed to parse VCP 0x%02x", vcpCode)
		return nil, fmt.Errorf("could not parse output")
	}

	current, _ := strconv.Atoi(matches[1])
	maximum, _ := strconv.Atoi(matches[2])

	return &VCPValue{
		Current: current,
		Max:     maximum,
	}, nil
}

// setVCP sets a VCP value on the display
func setVCP(displayNum int, vcpCode VCPCode, value int) error {
	codeStr := fmt.Sprintf("%02x", vcpCode)

	cmd := exec.Command("ddcutil", "setvcp", "-d", strconv.Itoa(displayNum), codeStr, strconv.Itoa(value))

	if err := cmd.Run(); err != nil {
		log.Error("failed to set VCP 0x%02x to %d: %w", vcpCode, value, err)
		return fmt.Errorf("failed to set VCP 0x%02x to %d: %w", vcpCode, value, err)
	}

	return nil
}
