package main

import (
	"context"
	"fmt"
	"os"

	"clean-architecture-learning/drivers"
)

func main() {
	ctx := context.Background()
	userDriver, err := drivers.InitializeUserDriver(ctx)
	if err != nil {
		fmt.Printf("failed to create UserDriver: %s\n", err)
		os.Exit(2)
	}
	userDriver.ServeUsers(ctx, ":8000")
}
