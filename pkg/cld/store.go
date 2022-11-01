package cld

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Cld interface {
	UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*uploader.UploadResult, error)
	UploadFileURL(ctx context.Context, url string) (*uploader.UploadResult, error)
	DeleteFile(ctx context.Context, publicID, resourceType string) error
}

type cld struct {
	*cloudinary.Cloudinary
}

func New(c *cloudinary.Cloudinary) Cld {
	return &cld{
		Cloudinary: c,
	}
}

func (cld *cld) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*uploader.UploadResult, error) {
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

func (cld *cld) UploadFileURL(ctx context.Context, url string) (*uploader.UploadResult, error) {
	resp, err := http.Get(url) //nolint:gosec // url is required to be a function parameter
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

func (cld *cld) DeleteFile(ctx context.Context, publicID, resourceType string) error {
	_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicID, ResourceType: resourceType})

	return err
}
