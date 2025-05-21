package background

import (
    "context"
    "encoding/json"
    "fmt"
    "log"

    "github.com/hibiken/asynq"
)

type TransferPayload struct {
    FromUserID string
    ToUserID   string
    Amount     int64
    Remarks    string
}

const TypeTransfer = "transfer:process"

func NewAsynqServer(redisAddr string) *asynq.Server {
    return asynq.NewServer(
        asynq.RedisClientOpt{Addr: redisAddr},
        asynq.Config{
            Concurrency: 10,
        },
    )
}

func NewTransferMux() *asynq.ServeMux {
    mux := asynq.NewServeMux()
    mux.HandleFunc(TypeTransfer, HandleTransferTask)
    return mux
}

func HandleTransferTask(ctx context.Context, t *asynq.Task) error {
    var payload TransferPayload
    if err := json.Unmarshal(t.Payload(), &payload); err != nil {
        return fmt.Errorf("failed to unmarshal payload: %v", err)
    }

    log.Printf("Processing transfer: from %s to %s, amount: %d, remarks: %s",
        payload.FromUserID, payload.ToUserID, payload.Amount, payload.Remarks)

    // TODO: Implement actual DB logic to perform transfer
    return nil
}