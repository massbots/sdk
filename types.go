package sdk

import (
	"go.massbots.xyz/sdk/internal/openapi"
)

type (
	Balance        = openapi.Balance
	Video          = openapi.Video
	Channel        = openapi.Channel
	VideoFormat    = openapi.VideoFormat
	VideoFormats   struct{ openapi.VideoFormats }
	DownloadResult = openapi.DownloadResult
)

func (v VideoFormats) Cached() []VideoFormat {
	var cached []VideoFormat
	for _, format := range v.Formats {
		if format.Cached {
			cached = append(cached, format)
		}
	}
	return cached
}

func (v VideoFormats) NotCached() []VideoFormat {
	var cached []VideoFormat
	for _, format := range v.Formats {
		if !format.Cached {
			cached = append(cached, format)
		}
	}
	return cached
}
