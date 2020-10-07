package main

import (
	"context"
	"log"

	"fmt"
	
    firebase "firebase.google.com/go"
    "google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
    sa := option.WithCredentialsFile("path/to/serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
		log.Fatalln(err)
	}

	// データ追加
    _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
        "first":  "Ada",
        "middle": "Mathison",
        "last":   "Lovelace",
        "born":   1815,
    })
    if err != nil {
        log.Fatalf("Failed adding alovelace: %v", err)
    }
 
    // 切断
    defer client.Close()
}