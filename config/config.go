package config

import (
	raven "github.com/getsentry/raven-go"
)

// Init to set the DSN for Sentry logger
func Init() {

	// Init Sentry
	raven.SetDSN("https://sentry")

}
