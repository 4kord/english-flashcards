package cld

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Cld struct {
	*cloudinary.Cloudinary
}

func New(cld *cloudinary.Cloudinary) *Cld {
	return &Cld{
		Cloudinary: cld,
	}
}

func (cld *Cld) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*uploader.UploadResult, error) {
	content, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer content.Close()

	res, err := cld.Upload.Upload(ctx, content, uploader.UploadParams{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (cld *Cld) UploadFileUrl(ctx context.Context, url string) (*uploader.UploadResult, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := cld.Upload.Upload(ctx, resp.Body, uploader.UploadParams{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cld *Cld) DeleteFile(ctx context.Context, publicId, resourceType string) error {
	_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicId, ResourceType: resourceType})

	return err
}
