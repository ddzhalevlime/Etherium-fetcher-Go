package main

import (
	"context"
	"log"
)

func (app *application) StartEventListener() {
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		log.Println("Starting event listener in background...")
		app.contractInteractor.ListenForEvents(ctx, app.personInfoEvents)
	}()
}
