package opnsense

import (
	"errors"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "[opnsense] ", log.Ldate|log.Ltime)

// SetLogger is used to set the default logger for critical errors.
// The initial logger is os.Stderr.
func SetLogger(l *log.Logger) error {
	if l == nil {
		return errors.New("logger is nil")
	}
	logger = l
	return nil
}
