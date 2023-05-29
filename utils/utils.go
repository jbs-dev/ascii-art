package utils

var ValidBanners = []string{
	"standard",
	"shadow",
	"thinkertoy",
	"colossal",
	"graffiti",
	"metric",
	"matrix",
	"rev",
	"card",
}

func IsValidBanner(banner string) bool {
	for _, validBanner := range ValidBanners {
		if banner == validBanner {
			return true
		}
	}
	return false
}
