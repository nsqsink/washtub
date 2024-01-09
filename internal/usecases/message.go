package usecases

import (
	"context"

	"github.com/nsqsink/washtub/internal/models"
)

type MessageUsecase struct {
	MessageStore models.MessageStore
}

func NewMessageUsecase(store models.MessageStore) models.MessageUsecase {
	return &MessageUsecase{
		MessageStore: store,
	}
}

// Fetch implements models.MessageUsecase.
func (w *MessageUsecase) Fetch(ctx context.Context, request models.FetchRequest, workerID string) (res []models.Message, err error) {
	res, err = w.MessageStore.Fetch(ctx, request, workerID)
	if err != nil {
		return nil, err
	}
	return
}

// GetByID implements models.MessageUsecase.
func (w *MessageUsecase) GetByID(ctx context.Context, id string) (res models.Message, err error) {
	res, err = w.MessageStore.GetByID(ctx, id)
	return
}

// Pulse implements models.MessageUsecase.
func (w *MessageUsecase) Store(ctx context.Context, message models.Message) (err error) {
	err = w.MessageStore.Store(ctx, message)
	return
}
