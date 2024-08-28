// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi

const (
	TokenScopes = "Token.Scopes"
)

// ApiBalance Token balance in points.
type ApiBalance struct {
	Balance int `json:"balance"`
}

// ApiError API error
type ApiError struct {
	Error string `json:"error"`
}

// DlFileID defines model for dl.FileID.
type DlFileID struct {
	// FileId Telegram's `file_id` used for sending cached files
	FileId string `json:"file_id"`
}

// DlResult defines model for dl.Result.
type DlResult struct {
	// FileId Telegram's `file_id` used for sending cached files
	FileId string `json:"file_id"`

	// Status One of: `queued`, `downloading`, `ready`, `failed`
	Status string `json:"status"`
}

// DlVideo Video with available formats
type DlVideo struct {
	// Formats Available formats
	Formats map[string]DlVideoFormat `json:"formats"`
}

// DlVideoFormat Video format
type DlVideoFormat struct {
	// Cached Set to `true` if the format is cached for quick downloading<br>
	// See: <ins>[Get cached video](#tag/download/GET/dl/video/{id}/cached/{f})</ins>
	Cached bool `json:"cached"`

	// FileSize File size in bytes, available only if video is not cached
	FileSize int `json:"file_size"`

	// Format Video resolution
	Format string `json:"format"`
}

// YtChannel YouTube channel
type YtChannel struct {
	BannerUrl       string                 `json:"banner_url"`
	CommentCount    int                    `json:"comment_count"`
	Description     string                 `json:"description"`
	Id              string                 `json:"id"`
	SubscriberCount int                    `json:"subscriber_count"`
	Thumbnails      map[string]YtThumbnail `json:"thumbnails"`
	Title           string                 `json:"title"`
	Url             string                 `json:"url"`
	VideoCount      int                    `json:"video_count"`
	ViewCount       int                    `json:"view_count"`
}

// YtThumbnail YouTube thumbnail
type YtThumbnail struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

// YtVideo YouTube video
type YtVideo struct {
	CategoryId   string                 `json:"category_id"`
	ChannelId    string                 `json:"channel_id"`
	ChannelTitle string                 `json:"channel_title"`
	ChannelUrl   string                 `json:"channel_url"`
	CommentCount int                    `json:"comment_count"`
	Description  string                 `json:"description"`
	Id           string                 `json:"id"`
	LikeCount    int                    `json:"like_count"`
	PublishedAt  string                 `json:"published_at"`
	Thumbnails   map[string]YtThumbnail `json:"thumbnails"`
	Title        string                 `json:"title"`
	Url          string                 `json:"url"`
	ViewCount    int                    `json:"view_count"`
}

// GetYtSearchParams defines parameters for GetYtSearch.
type GetYtSearchParams struct {
	// Q Query
	Q string `form:"q" json:"q"`
}