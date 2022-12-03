package cards

import (
	"context"
	"database/sql"
	"errors"
	"mime/multipart"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/null"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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

func (s *service) EditCard(ctx context.Context, arg *EditCardParams) (*maindb.Card, error) {
	card, err := s.store.GetCard(ctx, arg.CardID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.E(err, errs.NotExist, errs.Code("card_not_found"))
		}

		return nil, errs.E(err, errs.Database, errs.Code("get_card_failed"))
	}

	var imagePublicID, imageURL null.String

	switch { //nolint:dupl // .
	case arg.Image != nil:
		// if image exists
		var res *uploader.UploadResult

		res, err = s.cld.UploadFile(ctx, arg.Image)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("image_upload_failed"))
		}

		imagePublicID = null.String(res.PublicID)
		imageURL = null.String(res.SecureURL)
	case card.ImageUrl == arg.ImageURL:
		// if image url is the same
		imagePublicID = card.Image
		imageURL = card.ImageUrl
	case arg.ImageURL != "":
		// if image url exists
		var res *uploader.UploadResult

		res, err = s.cld.UploadFileURL(ctx, string(arg.ImageURL))
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("image_upload_failed"))
		}

		imagePublicID = null.String(res.PublicID)
		imageURL = null.String(res.SecureURL)
	}

	var audioPublicID, audioURL null.String

	switch { //nolint:dupl // .
	case arg.Audio != nil:
		// if audio exists
		var res *uploader.UploadResult

		res, err = s.cld.UploadFile(ctx, arg.Audio)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("audio_upload_failed"))
		}

		audioPublicID = null.String(res.PublicID)
		audioURL = null.String(res.SecureURL)
	case card.AudioUrl == arg.AudioURL:
		// if audio url is the same
		audioPublicID = card.Audio
		audioURL = card.AudioUrl
	case arg.AudioURL != "":
		// if audio url exists
		var res *uploader.UploadResult

		res, err = s.cld.UploadFileURL(ctx, string(arg.AudioURL))
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("audio_upload_failed"))
		}

		audioPublicID = null.String(res.PublicID)
		audioURL = null.String(res.SecureURL)
	}

	// if nothing is uploaded - value will be null (by default url is the current url in db)
	editedCard, err := s.store.EditCard(ctx, maindb.EditCardParams{
		ID:            arg.CardID,
		English:       arg.English,
		Russian:       arg.Russian,
		Association:   arg.Association,
		Example:       arg.Example,
		Transcription: arg.Transcription,
		Image:         imagePublicID,
		ImageUrl:      imageURL,
		Audio:         audioPublicID,
		AudioUrl:      audioURL,
	})
	if err != nil {
		return nil, errs.E(err, errs.Database, errs.Code("edit_card_failed"))
	}

	// if everything is successful, remove previous files
	if card.Image != "" {
		err = s.cld.DeleteFile(ctx, string(card.Image), "image")
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("delete_image_failed"))
		}
	}

	if card.Audio != "" {
		err = s.cld.DeleteFile(ctx, string(card.Audio), "video")
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("delete_audio_failed"))
		}
	}

	// if there was error deleting previous file, edited card will still be committed

	return editedCard, nil
}
