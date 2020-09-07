package main

import (
	"flag"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	var gpioPin int
	var onThreshold int
	var offThreshold int
	var sleepInterval int

	flag.IntVar(&gpioPin, "pin", 17, "Which GPIO pin you're using to control the fan.")
	flag.IntVar(&onThreshold, "on", 65, "(degrees Celsius) Fan kicks on at this temperature.")
	flag.IntVar(&offThreshold, "off", 55, "(degrees Celsius) Fan shuts off at this temperature.")
	flag.IntVar(&sleepInterval, "sleep", 5, "(seconds) How often we check the core temperature.")

	flag.Parse()
	if offThreshold >= onThreshold {
		log.Fatal("off threshold must be less than on threshold")
	}

	err := rpio.Open()
	defer rpio.Close()

	if err != nil {
		log.Fatal("Cannot open GPIO try run as root")
	}

	gpio := rpio.Pin(gpioPin)
	fanOn := false
	gpio.Output()

	for {
		temp := getTemp()

		if temp > onThreshold && !fanOn {
			gpio.High()
			fanOn = true
		} else if fanOn && temp < offThreshold {
			gpio.Low()
			fanOn = false
		}

		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}

func getTemp() int {
	cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("cat /sys/class/thermal/thermal_zone0/temp failed with %s\n", err)
	}

	val := string(out)
	val = strings.TrimSuffix(val, "\n")
	read, err := strconv.Atoi(val)

	if err != nil {
		log.Fatalf("failed to convert output\n%s\n", err)
	}

	return int(read / 1000)
}
