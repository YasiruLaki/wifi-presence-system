//go:build darwin

package sensors

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetRSSI() (int, error) {
	output, err := exec.Command("sudo", "wdutil", "info").Output()
	if err != nil {
		return 0, err
	}

	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "RSSI") {
			parts := strings.Fields(line)
			for i := len(parts) - 1; i >= 0; i-- {
				token := strings.TrimSuffix(parts[i], ":")
				rssi, err := strconv.Atoi(token)
				if err == nil {
					return rssi, nil
				}
			}
		}
	}

	return 0, nil
}
