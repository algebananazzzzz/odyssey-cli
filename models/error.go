package models

import (
	"fmt"
)

func PrintErrorAndExit(err error) {
	fmt.Printf(`╷
│ Error: Failed to read variables file
│ 
│ %v
╵
`, err)
}
