package main

import (
	"fmt"
	"time"

	"wifi-presence-system/sensors"
)

const (
	windowSize  = 20
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorCyan   = "\033[36m"
	colorBlue   = "\033[34m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
)

var rssiWindow []int

func getRSSI() (int, error) {
	return sensors.GetRSSI()
}

func addToWindow(val int) {
	if len(rssiWindow) >= windowSize {
		rssiWindow = rssiWindow[1:]
	}
	rssiWindow = append(rssiWindow, val)
}

func mean() float64 {
	if len(rssiWindow) == 0 {
		return 0
	}

	sum := 0
	for _, v := range rssiWindow {
		sum += v
	}

	return float64(sum) / float64(len(rssiWindow))
}

func variance() float64 {
	if len(rssiWindow) < windowSize {
		return 0
	}

	m := mean()
	var sum float64
	for _, v := range rssiWindow {
		diff := float64(v) - m
		sum += diff * diff
	}

	return sum / float64(len(rssiWindow))
}

func detectState(varVal float64) string {
	if varVal > 2.0 {
		return "MOVING"
	} else if varVal > 1.0 {
		return "STILL"
	}
	return "EMPTY"
}

func colorForState(state string) string {
	switch state {
	case "MOVING":
		return colorRed
	case "STILL":
		return colorYellow
	default:
		return colorGreen
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	for {
		rssi, err := getRSSI()
		if err != nil {
			fmt.Println("Error:", err)
			time.Sleep(time.Second)
			continue
		}

		addToWindow(rssi)

		varVal := variance()
		state := detectState(varVal)

		clearScreen()

		fmt.Printf("%s%sWiFi Invisible Presence System%s\n", colorBold, colorCyan, colorReset)
		fmt.Printf("%s------------------------------%s\n", colorCyan, colorReset)
		fmt.Printf("%sRSSI:%s %s%d dBm%s\n", colorBlue, colorReset, colorGreen, rssi, colorReset)
		fmt.Printf("%sVariance:%s %.2f\n", colorBlue, colorReset, varVal)
		fmt.Printf("%sState:%s %s%s%s\n", colorBlue, colorReset, colorForState(state), state, colorReset)
		fmt.Println()

		bars := int(varVal * 2)
		fmt.Printf("%sActivity:%s ", colorBlue, colorReset)
		for i := 0; i < bars; i++ {
			fmt.Printf("%s█%s", colorForState(state), colorReset)
		}
		fmt.Println()

		time.Sleep(5 * time.Millisecond)
	}

}
