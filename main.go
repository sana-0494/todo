package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"todo/cmd"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	root := cmd.New()

	_, err := root.ExecuteContextC(ctx)
	if err != nil {
		log.Panic(err)
	}
}
