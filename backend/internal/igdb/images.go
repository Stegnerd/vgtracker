package igdb

import "strings"

const baseString = "https://images.igdb.com/igdb/image/upload/t_{size}/{hash}.jpg"

type ImageSize string

const (
	CoverSmall       ImageSize = "cover_small"
	ScreenshotMedium ImageSize = "screenshot_med"
	CoverBig         ImageSize = "cover_big"
	LogoMedium       ImageSize = "logo_med"
	ScreenshotBig    ImageSize = "screenshot_big"
	ScreenshotHuge   ImageSize = "screenshot_huge"
	Thumb            ImageSize = "thumb"
	Micro            ImageSize = "micro"
	P720             ImageSize = "720p"
	P1080            ImageSize = "1080p"
)

type ImageResolverMethods interface {
	GetImageURL(size ImageSize, imageHash string) string
}

type ImageResolver struct{}

func NewImageResolver() ImageResolverMethods {
	return &ImageResolver{}
}

func (i *ImageResolver) GetImageURL(size ImageSize, imageHash string) string {
	url := baseString
	url = strings.Replace(url, "{size}", string(size), 1)
	url = strings.Replace(url, "{hash}", imageHash, 1)
	return url
}
