package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	clientID := uuid.New().String()
	fmt.Println("Generated Client ID:", clientID)
}
