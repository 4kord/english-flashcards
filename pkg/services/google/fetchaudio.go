package google

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/4kord/english-flashcards/pkg/errs"
)

func (s *service) FetchAudio(ctx context.Context, word string) (string, error) {
	url := fmt.Sprintf("https://ssl.gstatic.com/dictionary/static/pronunciation/2022-03-02/audio/%v/%v_en_us_1.mp3", strings.ToLower(word)[0:2], strings.ToLower(word))

	resp, err := http.Get(url) //nolint:gosec // it is required to build url from user's input
	if err != nil {
		return "", errs.E(err, errs.IO, errs.Code("get_audio_failed"))
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return "", errs.E("Audio response status code != 200", errs.NotExist, errs.Code("audio_not_found"))
	}

	return url, nil
}
