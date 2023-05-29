package utils

import "fmt"

// BannerDetails contains information about a banner
type BannerDetail struct {
	LineCount int
	Offset    int
}

func GetBannerDetails(banner string) (BannerDetail, error) {
	switch banner {
	case "card":
		return BannerDetail{LineCount: 7, Offset: 223}, nil
	case "metric":
		return BannerDetail{LineCount: 11, Offset: 705}, nil
	case "graffiti":
		return BannerDetail{LineCount: 6, Offset: 198}, nil
	case "matrix":
		return BannerDetail{LineCount: 10, Offset: 320}, nil
	case "rev":
		return BannerDetail{LineCount: 11, Offset: 353}, nil
	default:
		return BannerDetail{}, fmt.Errorf("unknown banner: %s", banner)
	}
}
