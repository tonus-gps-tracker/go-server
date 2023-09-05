package utils

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func GetEnv(env string) string {
	value, isSet := os.LookupEnv(env)
	if !isSet {
		log.Fatalf("[ERROR] Environment variable \"%s\" not set\n", env)
	}

	return value
}

func WaitInterruption() {
	var waitGroup sync.WaitGroup

	osInterrupt := make(chan os.Signal, 1)
	signal.Notify(osInterrupt, os.Interrupt)

	syscallSigterm := make(chan os.Signal, 1)
	signal.Notify(syscallSigterm, syscall.SIGTERM)

	waitGroup.Add(1)

	go func() {
		<-osInterrupt
		defer waitGroup.Done()
	}()

	go func() {
		<-syscallSigterm
		defer waitGroup.Done()
	}()

	waitGroup.Wait()
}

func StringToInt(stringValue string) int {
	intValue, err := strconv.Atoi(stringValue)
	if err != nil {
		log.Printf("[ERROR] StringToInt, strconv.Atoi: %s\n", err)
		return 0
	}

	return intValue
}

func StringToFloat32(stringValue string) float32 {
	floatValue, err := strconv.ParseFloat(stringValue, 32)

	if err != nil {
		log.Printf("[ERROR] StringToFloat32, strconv.ParseFloat: %s\n", err)
		return 0
	}

	return float32(floatValue)
}
