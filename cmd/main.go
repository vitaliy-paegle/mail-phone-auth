package main

import (
	"fmt"
	"mail-phone-auth/internal/app"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	application := app.New()
	application.Run()

	stopSignalHandler(application.Stop)
}

func stopSignalHandler(handler func()) {

	const pauseTime = 3

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	sig := <- signalChannel

	fmt.Printf("\n Stop signal: %v \n", sig)

	handler()
	time.Sleep(pauseTime * time.Second)

	fmt.Println("Program is stoped")
}



