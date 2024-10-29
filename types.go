package sdk

import (
	"go.massbots.xyz/sdk/internal/openapi"
)

type (
	Balance        = openapi.Balance
	RequestStat    = openapi.RequestStat
	Video          = openapi.Video
	Channel        = openapi.Channel
	VideoFormat    = openapi.VideoFormat
	VideoFormats   map[string]VideoFormat
	DownloadResult = openapi.DownloadResult
)

func (v VideoFormats) Cached() []VideoFormat {
	var cached []VideoFormat
	for _, format := range v {
		if format.Cached {
			cached = append(cached, format)
		}
	}
	return cached
}

func (v VideoFormats) NotCached() []VideoFormat {
	var cached []VideoFormat
	for _, format := range v {
		if !format.Cached {
			cached = append(cached, format)
		}
	}
	return cached
}
