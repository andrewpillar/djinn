package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/andrewpillar/djinn/errors"
	"github.com/andrewpillar/djinn/version"
	"github.com/andrewpillar/djinn/workerutil"
)

func main() {
	config, driver, showversion := workerutil.ParseFlags(os.Args)

	if showversion {
		fmt.Printf("%s %s %s\n", os.Args[0], version.Tag, version.Ref)
		return
	}

	w, _, close_, err := workerutil.Init(config, driver)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], errors.Cause(err))
		os.Exit(1)
	}

	defer close_()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	workerutil.Start(ctx, w)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	sig := <-c

	cancel()
	w.Log.Info.Println("signal:", sig, "received, shutting down")
}
