package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	practice "github.com/maruware/gstreamer-practice"
	"golang.org/x/sync/errgroup"
)

func main() {
	port := 60000

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sig
		cancel()
	}()

	eg, ctx := errgroup.WithContext(ctx)

	// Publisher
	eg.Go(func() error {
		cmd := practice.Publisher(port)
		log.Printf(cmd.String())

		return cmd.Run()
	})

	// Receiver
	eg.Go(func() error {
		cmd := practice.Receiver(port)
		log.Printf(cmd.String())

		return cmd.Run()
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
