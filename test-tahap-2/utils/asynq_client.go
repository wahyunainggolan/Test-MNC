package utils

import (
    "encoding/json"
    "log"

    "github.com/hibiken/asynq"
    "wallet-api/internal/background"
)

var RedisClient = asynq.RedisClientOpt{
    Addr: "localhost:6379",
}

func EnqueueTransferTask(fromUserID, toUserID string, amount int64, remarks string) error {
    payload, err := json.Marshal(background.TransferPayload{
        FromUserID: fromUserID,
        ToUserID:   toUserID,
        Amount:     amount,
        Remarks:    remarks,
    })
    if err != nil {
        return err
    }

    task := asynq.NewTask(background.TypeTransfer, payload)
    client := asynq.NewClient(RedisClient)
    defer client.Close()

    info, err := client.Enqueue(task)
    if err != nil {
        return err
    }

    log.Printf("Enqueued task: id=%s queue=%s", info.ID, info.Queue)
    return nil
}