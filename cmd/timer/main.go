package main

import (
	"context"
	"fmt"
	"os"
	"pomodoro/internal/app"
)

func main() {
	ctx := context.Background()
	app := app.NewApp()

	err := app.Run(ctx)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}
