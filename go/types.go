package sdk

import (
	"go.massbots.xyz/sdk/internal/openapi"
)

type (
	Balance        = openapi.ApiBalance
	YoutubeVideo   = openapi.YtVideo
	YoutubeChannel = openapi.YtChannel
	Video          struct{ openapi.DlVideo }
	VideoFormat    = openapi.DlVideoFormat
	DownloadResult = openapi.DlResult
	DownloadFileID = openapi.DlFileID
)

func (v Video) Cached() []VideoFormat {
	var cached []VideoFormat
	for _, format := range v.Formats {
		if format.Cached {
			cached = append(cached, format)
		}
	}
	return cached
}

func (v Video) NotCached() []VideoFormat {
	var cached []VideoFormat
	for _, format := range v.Formats {
		if !format.Cached {
			cached = append(cached, format)
		}
	}
	return cached
}
