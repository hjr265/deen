package model

import (
	"time"

	"github.com/hjr265/deen/aladhan"
)

type Adhan struct {
	Name    string
	When    time.Time
	City    string
	Country string
	Method  aladhan.Method
}
