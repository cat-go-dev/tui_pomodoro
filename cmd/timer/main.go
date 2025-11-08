package main

import (
	"fmt"
	"os"
	"pomodoro/internal/app"
)

func main() {
	app := app.NewApp()

	err := app.Run()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}
