package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {
	log := log.New(os.Stdout, "ZIP ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(log); err != nil {
		log.Println("main : ", err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {
	// =========================================================================
	// Configuration

	var cfg struct {
		conf.Version
		API struct {
			Host    string        `conf:"default:localhost:3000"`
			Timeout time.Duration `conf:"default:5s"`
		}
	}
	cfg.Version.SVN = build
	cfg.Version.Desc = "copyright information here"

	if err := conf.Parse(os.Args[1:], "ZIP", &cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, err := conf.Usage("ZIP", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		case conf.ErrVersionWanted:
			version, err := conf.VersionString("ZIP", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config version")
			}
			fmt.Println(version)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	// =========================================================================
	// App Starting

	log.Println("main : starting : initializing application")
	defer log.Println("main : completed")

	// =========================================================================
	// Start API Service

	log.Println("main : initializing api")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	http.HandleFunc("/", hello)

	go func() {
		http.ListenAndServe("localhost:3000", nil)
	}()

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case sig := <-shutdown:
		log.Printf("main : %v : starting shutdown", sig)

		// Give outstanding requests a deadline for completion.
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Log the status of this shutdown.
		if sig == syscall.SIGSTOP {
			return errors.New("main : integrity issue caused shutdown")
		}
	}

	return nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello zipcode")
}
