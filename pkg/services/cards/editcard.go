package cards

import (
	"context"
	"database/sql"
	"errors"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/services/cards/dto"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (s *service) EditCard(ctx context.Context, arg *dto.EditCardParams) (*maindb.Card, error) {
	card, err := s.store.GetCard(ctx, arg.CardID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.E(err, errs.NotExist, errs.Code("card_not_found"))
		}

		return nil, errs.E(err, errs.Database, errs.Code("get_card_failed"))
	}

	var imagePublicID, imageURL sql.NullString

	switch { //nolint:dupl // .
	case arg.Image != nil:
		// if image exists
		var res *uploader.UploadResult

		res, err = s.cld.UploadFile(ctx, arg.Image)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("image_upload_failed"))
		}

		imagePublicID = sql.NullString{String: res.PublicID, Valid: true}
		imageURL = sql.NullString{String: res.SecureURL, Valid: true}
	case card.ImageUrl.String == arg.ImageURL.String:
		// if image url is the same
		imagePublicID = card.Image
		imageURL = card.ImageUrl
	case arg.ImageURL.Valid:
		// if image url exists
		var res *uploader.UploadResult

		res, err = s.cld.UploadFileURL(ctx, arg.ImageURL.String)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("image_upload_failed"))
		}

		imagePublicID = sql.NullString{String: res.PublicID, Valid: true}
		imageURL = sql.NullString{String: res.SecureURL, Valid: true}
	}

	var audioPublicID, audioURL sql.NullString

	switch { //nolint:dupl // .
	case arg.Audio != nil:
		// if audio exists
		var res *uploader.UploadResult

		res, err = s.cld.UploadFile(ctx, arg.Audio)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("audio_upload_failed"))
		}

		audioPublicID = sql.NullString{String: res.PublicID, Valid: true}
		audioURL = sql.NullString{String: res.SecureURL, Valid: true}
	case card.AudioUrl.String == arg.AudioURL.String:
		// if audio url is the same
		audioPublicID = card.Audio
		audioURL = card.AudioUrl
	case arg.AudioURL.Valid:
		// if audio url exists
		var res *uploader.UploadResult

		res, err = s.cld.UploadFileURL(ctx, arg.AudioURL.String)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("audio_upload_failed"))
		}

		audioPublicID = sql.NullString{String: res.PublicID, Valid: true}
		audioURL = sql.NullString{String: res.SecureURL, Valid: true}
	}

	// if nothing is uploaded - value will be null (by default url is the current url in db)
	editedCard, err := s.store.EditCard(ctx, maindb.EditCardParams{
		ID:            arg.CardID,
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
		return nil, errs.E(err, errs.Database, errs.Code("edit_card_failed"))
	}

	// if everything is successful, remove previous files
	if card.Image.Valid {
		err = s.cld.DeleteFile(ctx, card.Image.String, "image")
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("delete_image_failed"))
		}
	}

	if card.Audio.Valid {
		err = s.cld.DeleteFile(ctx, card.Audio.String, "video")
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("delete_audio_failed"))
		}
	}

	// if there was error deleting previous file, edited card will still be committed

	return editedCard, nil
}
