package main

import (
	"context"
	"docsign/api"
	"docsign/utils/server"
	"log"

	"docsign/utils/di"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cxt := context.Background()
	app := di.Init(di.Module{
		Imports: di.Modules{
			server.Module(),
			api.Module(),
		},
		Provides: di.Deps{
			func() *context.Context {
				return &cxt
			},
		},
	})
	if err := app.Start(cxt); err != nil {
		log.Fatal(err)
	}
}
