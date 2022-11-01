package setup

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

func setupCLD(cloud, key, secret string) *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(cloud, key, secret)
	if err != nil {
		panic(err)
	}

	return cld
}
