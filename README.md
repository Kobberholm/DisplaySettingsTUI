# DisplaySettingsTUI

A terminal user interface (TUI) application for adjusting display brightness and contrast settings on Linux systems using DDC/CI protocol.

![Go Version](https://img.shields.io/badge/go-%3E%3D1.25-blue)
![License](https://img.shields.io/badge/license-MIT-green)
[![GitHub](https://img.shields.io/badge/GitHub-navinreddy23%2FDisplaySettingsTUI-blue)](https://github.com/navinreddy23/DisplaySettingsTUI)

![Demo](demo.gif)

## Features

- 🖥️ Auto-detect connected displays
- 🔆 Adjust brightness and contrast with arrow keys
- 🎨 Beautiful terminal interface with progress bars
- ⚡ Real-time display adjustments with debouncing
- 🎯 Navigate between multiple displays easily

## Prerequisites

### 1. Install ddcutil

This application requires `ddcutil` to communicate with displays via DDC/CI protocol.

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install ddcutil
```

**Fedora:**
```bash
sudo dnf install ddcutil
```

**Arch Linux:**
```bash
sudo pacman -S ddcutil
```

### 2. Configure I2C Permissions

To use DDC/CI without root privileges, you need to configure I2C permissions properly.

#### Step 1: Add your user to the i2c group
```bash
sudo usermod -aG i2c $USER
```

#### Step 2: Create udev rule for I2C devices
Create a new udev rule file:
```bash
sudo nano /etc/udev/rules.d/45-ddcutil-i2c.rules
```

Add the following content:
```
# Rules for display detection using ddcutil
SUBSYSTEM=="i2c-dev", MODE="0660", GROUP="i2c"
```

#### Step 3: Load i2c-dev module
Ensure the i2c-dev module loads at boot:
```bash
sudo sh -c 'echo "i2c-dev" >> /etc/modules-load.d/modules.conf'
```

Load the module immediately:
```bash
sudo modprobe i2c-dev
```

#### Step 4: Apply changes
Reload udev rules and log out/in for group changes to take effect:
```bash
sudo udevadm control --reload-rules
sudo udevadm trigger
```

**Important:** Log out and log back in for the group membership to take effect!

#### Verify Setup
After logging back in, verify everything is working:
```bash
# Check if you're in the i2c group
groups | grep i2c

# Test ddcutil
ddcutil detect

# Check I2C permissions
ls -l /dev/i2c-*
```

For more detailed configuration steps, see: [DDC/CI Configuration Guide](https://www.ddcutil.com/config_steps/)

## Installation

### Option 1: Install from source

Clone the repository:
```bash
git clone https://github.com/navinreddy23/DisplaySettingsTUI.git
cd DisplaySettingsTUI
```

Build and install:
```bash
# First build the binary
make build

# Then install system-wide (requires sudo)
sudo make install

# OR install for current user only
make install-user
```

### Option 2: Download binary (Coming Soon)

Binary releases will be available on the [releases page](https://github.com/navinreddy23/DisplaySettingsTUI/releases) in the future.

Once available, you'll be able to:
```bash
wget https://github.com/navinreddy23/DisplaySettingsTUI/releases/latest/download/display-settings-tui
chmod +x display-settings-tui
sudo mv display-settings-tui /usr/local/bin/
```

## Usage

Simply run the application:
```bash
display-settings-tui
```

### Controls

#### Main Screen (Display Selection)
- `←/→` - Navigate between displays
- `Enter` - Select display to adjust
- `q` - Quit application

#### Settings Screen
- `↑/↓` - Switch between Brightness and Contrast
- `←/→` - Decrease/Increase selected setting (5% steps)
- `ESC` - Return to display selection
- `q` - Quit application

## Building from Source

### Requirements
- Go 1.25 or higher
- Make (optional, for using Makefile)

### Build Commands

```bash
# Build the binary
make build

# Run directly without installing
make run

# Clean build artifacts
make clean
```

## Troubleshooting

### "No displays found"
1. Ensure your monitor supports DDC/CI (most modern monitors do)
2. Enable DDC/CI in your monitor's OSD settings
3. Check cable connection (use DisplayPort or HDMI, VGA doesn't support DDC/CI)
4. Verify ddcutil is working: `ddcutil detect`

### "Permission denied" errors
1. Ensure you've followed all I2C permission configuration steps
2. Make sure you've logged out and back in after adding yourself to the i2c group
3. Check I2C device permissions: `ls -l /dev/i2c-*`

### Display not responding to adjustments
1. Some displays have DDC/CI disabled by default - check monitor settings
2. Try using ddcutil directly: `ddcutil setvcp 10 50` (sets brightness to 50%)
3. Some displays require specific cables or ports for DDC/CI support

## How It Works

DisplaySettingsTUI uses:
- **ddcutil** - For DDC/CI communication with displays
- **Bubble Tea** - For the terminal UI framework
- **VCP (Virtual Control Panel)** codes - Standard display control codes
  - Code 0x10: Brightness
  - Code 0x12: Contrast

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Acknowledgments

- [ddcutil](https://www.ddcutil.com/) - DDC/CI control utility
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling