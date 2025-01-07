package greet

import (
	"fmt"
	"time"
)

func Today() {
	now := time.Now()
	data_formatada := now.Format("02/01/2006 15:04:05")
	fmt.Printf("Hoje é %s", data_formatada)
}
