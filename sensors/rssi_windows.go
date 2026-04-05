//go:build windows

package sensors

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetRSSI() (int, error) {
	output, err := exec.Command("netsh", "wlan", "show", "interfaces").Output()
	if err != nil {
		return 0, err
	}

	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "Signal") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				rssiStr := strings.TrimSuffix(parts[len(parts)-1], "%")
				rssi, err := strconv.Atoi(rssiStr)
				if err != nil {
					return 0, err
				}

				return rssi, nil
			}
		}
	}

	return 0, nil
}
