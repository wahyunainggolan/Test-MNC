package main

import (
    "log"
    "wallet-api/internal/background"
)

func main() {
    srv := background.NewAsynqServer("localhost:6379")
    mux := background.NewTransferMux()

    if err := srv.Run(mux); err != nil {
        log.Fatalf("Could not run background worker: %v", err)
    }
}