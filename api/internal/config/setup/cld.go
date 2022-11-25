package setup

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

func setupCLD(cloudinaryURL string) *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		panic(err)
	}

	return cld
}
