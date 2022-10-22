package cards

import (
	"context"
	"database/sql"
	"mime/multipart"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/lib/pq"
)

func (s *service) CreateCard(ctx context.Context, card maindb.Card, image *multipart.FileHeader, audio *multipart.FileHeader) (*maindb.Card, error) {
	var imagePublicID, imageURL sql.NullString

	if image != nil {
		imageRes, err := s.cld.UploadFile(ctx, image)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("image_upload_failed"))
		}

		imagePublicID = sql.NullString{String: imageRes.PublicID, Valid: true}
		imageURL = sql.NullString{String: imageRes.SecureURL, Valid: true}
	}

	var audioPublicID, audioURL sql.NullString
	if audio != nil {
		audioRes, err := s.cld.UploadFile(ctx, audio)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("audio_upload_failed"))
		}

		audioPublicID = sql.NullString{String: audioRes.PublicID, Valid: true}
		audioURL = sql.NullString{String: audioRes.SecureURL, Valid: true}
	}

	createdCard, err := s.store.CreateCard(ctx, maindb.CreateCardParams{
		DeckID:        card.DeckID,
		English:       card.English,
		Russian:       card.Russian,
		Association:   card.Association,
		Example:       card.Example,
		Transcription: card.Transcription,
		Image:         imagePublicID,
		ImageUrl:      imageURL,
		Audio:         audioPublicID,
		AudioUrl:      audioURL,
	})

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Constraint == "cards_deck_id_fkey" {
				return nil, errs.E(err, errs.InvalidRequest, errs.Code("deck_not_found"))
			}
		}

		if imagePublicID.Valid {
			err := s.cld.DeleteFile(ctx, imagePublicID.String, "image")
			if err != nil {
				return nil, errs.E(err, errs.IO, errs.Code("delete_image_failed"))
			}
		}

		if audioPublicID.Valid {
			err = s.cld.DeleteFile(ctx, audioPublicID.String, "video")
			if err != nil {
				return nil, errs.E(err, errs.IO, errs.Code("delete_audio_failed"))
			}
		}

		return nil, errs.E(err, errs.Database, errs.Code("create_card_failed"))
	}

	return createdCard, nil
}
