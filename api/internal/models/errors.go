package models

import (
	"fmt"
)

type modelsError string

func (m modelsError) Error() string {
	return fmt.Sprintf("models: %s", m)
}

const (
	ErrNoRecord = modelsError("no matching record found")
)
