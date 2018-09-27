package utils

import (
	"fmt"
	"log"

	"github.com/getsentry/raven-go"
)

// LogError is a central function for sending error messages to Sentry and the local console
func LogError(err error, tags map[string]string) {

	// Log to Sentry
	raven.CaptureError(err, tags)

	// Log to console
	log.Println(fmt.Sprintf("%s -- %s", tags, err))
}
