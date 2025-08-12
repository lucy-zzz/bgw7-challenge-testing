package main

import (
	"app/internal/application"
	"fmt"
)

func runAppWithApplication(app application.Application) error {
	defer app.TearDown()

	if err := app.SetUp(); err != nil {
		return fmt.Errorf("failed to setup app: %w", err)
	}

	if err := app.Run(); err != nil {
		return fmt.Errorf("failed to run app: %w", err)
	}

	return nil
}

func runApp() error {
	// env
	// ...

	// app
	// - config
	cfg := &application.ConfigApplicationDefault{
		Addr: "127.0.0.1:8080",
	}
	app := application.NewApplicationDefault(cfg)

	return runAppWithApplication(app)
}

func main() {
	if err := runApp(); err != nil {
		fmt.Println(err)
		return
	}
}
