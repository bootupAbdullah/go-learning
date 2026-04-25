package timer

import (
	"time"
)

func SetTimer() *time.Ticker{
	ticker := time.NewTicker(time.Second)
	return ticker
}
