package cards

import (
	"context"
	"database/sql"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/services/cards/dto"
	"github.com/lib/pq"
)

func (s *service) CreateCard(ctx context.Context, arg *dto.CreateCardParams) (*maindb.Card, error) {
	var imagePublicID, imageURL sql.NullString

	switch {
	case arg.Image != nil:
		// if image exists
		res, err := s.cld.UploadFile(ctx, arg.Image)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("image_upload_failed"))
		}

		imagePublicID = sql.NullString{String: res.PublicID, Valid: true}
		imageURL = sql.NullString{String: res.SecureURL, Valid: true}
	case arg.ImageURL.Valid:
		// if image url exists
		res, err := s.cld.UploadFileURL(ctx, arg.ImageURL.String)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("image_upload_failed"))
		}

		imagePublicID = sql.NullString{String: res.PublicID, Valid: true}
		imageURL = sql.NullString{String: res.SecureURL, Valid: true}
	}

	var audioPublicID, audioURL sql.NullString

	switch {
	case arg.Audio != nil:
		// if audio exists
		res, err := s.cld.UploadFile(ctx, arg.Audio)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("audio_upload_failed"))
		}

		audioPublicID = sql.NullString{String: res.PublicID, Valid: true}
		audioURL = sql.NullString{String: res.SecureURL, Valid: true}
	case arg.AudioURL.Valid:
		// if audio url exists
		res, err := s.cld.UploadFileURL(ctx, arg.AudioURL.String)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("audio_upload_failed"))
		}

		audioPublicID = sql.NullString{String: res.PublicID, Valid: true}
		audioURL = sql.NullString{String: res.SecureURL, Valid: true}
	}

	// if nothing is uploaded, value will be null

	createdCard, err := s.store.CreateCard(ctx, maindb.CreateCardParams{
		DeckID:        arg.DeckID,
		English:       arg.English,
		Russian:       arg.Russian,
		Association:   sql.NullString(arg.Association),
		Example:       sql.NullString(arg.Example),
		Transcription: sql.NullString(arg.Transcription),
		Image:         imagePublicID,
		ImageUrl:      imageURL,
		Audio:         audioPublicID,
		AudioUrl:      audioURL,
	})

	if err != nil {
		// Delete image if was created
		if imagePublicID.Valid {
			err = s.cld.DeleteFile(ctx, imagePublicID.String, "image")
			if err != nil {
				return nil, errs.E(err, errs.IO, errs.Code("delete_image_failed"))
			}
		}

		// Delete audio if was created
		if audioPublicID.Valid {
			err = s.cld.DeleteFile(ctx, audioPublicID.String, "video")
			if err != nil {
				return nil, errs.E(err, errs.IO, errs.Code("delete_audio_failed"))
			}
		}

		if err, ok := err.(*pq.Error); ok {
			if err.Constraint == "cards_deck_id_fkey" {
				return nil, errs.E(err, errs.NotExist, errs.Code("deck_not_found"))
			}
		}

		return nil, errs.E(err, errs.Database, errs.Code("create_card_failed"))
	}

	return createdCard, nil
}
