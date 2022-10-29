package dto

import (
	"mime/multipart"

	"github.com/4kord/english-flashcards/pkg/null"
)

type EditCardParams struct {
	CardID        int32
	English       string
	Russian       string
	Association   null.String
	Example       null.String
	Transcription null.String
	Image         *multipart.FileHeader
	ImageURL      null.String
	Audio         *multipart.FileHeader
	AudioURL      null.String
}
