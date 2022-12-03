package cards

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/4kord/english-flashcards/pkg/errs"
	"github.com/4kord/english-flashcards/pkg/maindb"
	"github.com/4kord/english-flashcards/pkg/null"
	"github.com/lib/pq"
)

type CreateCardParams struct {
	DeckID        int32
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

func (s *service) CreateCard(ctx context.Context, arg *CreateCardParams) (*maindb.Card, error) {
	var imagePublicID, imageURL null.String

	switch {
	case arg.Image != nil:
		// if image exists
		res, err := s.cld.UploadFile(ctx, arg.Image)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("image_upload_failed"))
		}

		fmt.Println(res)

		imagePublicID = null.String(res.PublicID)
		imageURL = null.String(res.SecureURL)
	case arg.ImageURL != "":
		// if image url exists
		res, err := s.cld.UploadFileURL(ctx, string(arg.ImageURL))
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("image_upload_failed"))
		}

		imagePublicID = null.String(res.PublicID)
		imageURL = null.String(res.SecureURL)
	}

	var audioPublicID, audioURL null.String

	switch {
	case arg.Audio != nil:
		// if audio exists
		res, err := s.cld.UploadFile(ctx, arg.Audio)
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("audio_upload_failed"))
		}

		audioPublicID = null.String(res.PublicID)
		audioURL = null.String(res.SecureURL)
	case arg.AudioURL != "":
		// if audio url exists
		res, err := s.cld.UploadFileURL(ctx, string(arg.AudioURL))
		if err != nil {
			return nil, errs.E(err, errs.IO, errs.Code("audio_upload_failed"))
		}

		audioPublicID = null.String(res.PublicID)
		audioURL = null.String(res.SecureURL)
	}

	// if nothing is uploaded, value will be null

	var createdCard *maindb.Card

	err := s.store.ExecTx(ctx, func(q maindb.Querier) (bool, error) {
		var err error

		err = q.DeckAmountUp(ctx, arg.DeckID)
		if err != nil {
			return false, errs.E(err, errs.Internal, errs.Code("deck_amount_up_failed"))
		}

		createdCard, err = q.CreateCard(ctx, maindb.CreateCardParams{
			DeckID:        arg.DeckID,
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
			return false, errs.E(err, errs.Internal, errs.Code("create_card_failed"))
		}

		return true, nil
	})

	if err != nil {
		// Delete image if was created
		if imagePublicID != "" {
			err = s.cld.DeleteFile(ctx, string(imagePublicID), "image")
			if err != nil {
				return nil, errs.E(err, errs.IO, errs.Code("delete_image_failed"))
			}
		}

		// Delete audio if was created
		if audioPublicID != "" {
			err = s.cld.DeleteFile(ctx, string(audioPublicID), "video")
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
