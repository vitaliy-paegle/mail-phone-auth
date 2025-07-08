package main

import (
	"fmt"
	"mail-phone-auth/internal/app"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	application := app.NewApp()
	application.Run()
	stopSignalHandler(application.Stop)
}

func stopSignalHandler(handler func()) {

	const pauseTime = 2

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	sig := <- signalChannel

	fmt.Printf("\n Stop signal: %v \n", sig)

	handler()

	time.Sleep(pauseTime * time.Second)

	fmt.Println("Program stoped")
}



