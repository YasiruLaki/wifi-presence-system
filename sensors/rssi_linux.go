//go:build linux

package sensors

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetRSSI() (int, error) {
	out, err := exec.Command("bash", "-c", "nmcli -t -f IN-USE,SIGNAL dev wifi | grep '^*' | cut -d: -f2").Output()
	if err != nil {
		return 0, err
	}

	rssi, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return 0, err
	}

	return rssi, nil
}
