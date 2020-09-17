package internal

import (
	"time"

	"github.com/ardanlabs/conf"
)

//Config holds all configuration for this service
type Config struct {
	conf.Version
	API struct {
		Host    string        `conf:"default:localhost:3000"`
		Timeout time.Duration `conf:"default:5s"`
	}
}
