package google

import "context"

type Service interface {
	FetchAudio(ctx context.Context, word string) (string, error)
}

type service struct{}

func New() Service {
	return &service{}
}
