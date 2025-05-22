package main

import (
	"context"
	"log"
	"os"
	"pass-gen/app"
)

func main() {
	app := app.Launch()

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
