package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"errors"

	"github.com/joho/godotenv"
	"github.com/mkhuda/go-arduino-feeder/handlers"
	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

func main() {

	portConnected, err := findPort()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(portConnected)

	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.TwoStopBits,
	}
	openedSerial, err := serial.Open(portConnected, mode)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(openedSerial)
	for scanner.Scan() {
		payload := scanner.Text()
		isDHT := strings.Contains(payload, "dht-")
		if isDHT {
			// slice
			sliced := strings.Split(payload, "-")

			// _, _ := strconv.Atoi(sliced[1])            // temperature from DHT11 (unused)
			humData, _ := strconv.Atoi(sliced[2])      // humidity from DHT11
			heatData, _ := strconv.Atoi(sliced[3])     // heat calculate DHT11 Humidity and LM35 Temperature Data
			tempLm35Data, _ := strconv.Atoi(sliced[4]) // temperature from LM35

			handlers.Api_Post_Temperature(tempLm35Data, humData, heatData)
		}

		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findPort() (string, error) {
	var connectedPort string
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading ENV")
	}

	registeredPID := os.Getenv("PID")
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return "", err
	}
	if len(ports) == 0 {
		return "", errors.New("no serial ports found")
	}
	for _, port := range ports {
		if port.SerialNumber == registeredPID {
			connectedPort = port.Name
		}
	}

	return connectedPort, nil
}
