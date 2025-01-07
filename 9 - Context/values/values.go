package main

import (
	"context"
	"fmt"
)

func printUserID(ctx context.Context) {
	userID := ctx.Value("userID")
	if userID != nil {
		fmt.Println(userID)
	} else {
		fmt.Println("Key doesn't exist in context")
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "userID", 5)
	printUserID(ctx)
}
