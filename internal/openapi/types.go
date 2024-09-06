// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi

const (
	TokenScopes = "Token.Scopes"
)

// Balance Token balance in points.
type Balance struct {
	Balance int `json:"balance"`
}

// Channel YouTube channel
type Channel struct {
	BannerUrl       string               `json:"banner_url"`
	CommentCount    int                  `json:"comment_count"`
	Description     string               `json:"description"`
	Id              string               `json:"id"`
	SubscriberCount int                  `json:"subscriber_count"`
	Thumbnails      map[string]Thumbnail `json:"thumbnails"`
	Title           string               `json:"title"`
	Url             string               `json:"url"`
	VideoCount      int                  `json:"video_count"`
	ViewCount       int                  `json:"view_count"`
}

// DownloadResult defines model for DownloadResult.
type DownloadResult struct {
	// FileId Telegram's `file_id` used for sending cached files
	FileId string `json:"file_id"`

	// Status One of: `queued`, `downloading`, `ready`, `failed`
	Status string `json:"status"`
}

// Error API error
type Error struct {
	Error string `json:"error"`
}

// Thumbnail YouTube thumbnail
type Thumbnail struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

// Video YouTube video
type Video struct {
	CategoryId   string               `json:"category_id"`
	ChannelId    string               `json:"channel_id"`
	ChannelTitle string               `json:"channel_title"`
	ChannelUrl   string               `json:"channel_url"`
	CommentCount int                  `json:"comment_count"`
	Description  string               `json:"description"`
	Id           string               `json:"id"`
	LikeCount    int                  `json:"like_count"`
	PublishedAt  string               `json:"published_at"`
	Thumbnails   map[string]Thumbnail `json:"thumbnails"`
	Title        string               `json:"title"`
	Url          string               `json:"url"`
	ViewCount    int                  `json:"view_count"`
}

// VideoFormat Video format
type VideoFormat struct {
	// Cached Set to `true` if the format is cached for quick downloading<br>
	Cached bool `json:"cached"`

	// FileSize File size in bytes, available only if video is not cached
	FileSize int `json:"file_size"`

	// Format Video resolution
	Format string `json:"format"`
}

// BotId Your Telegram Bot ID
type BotId = string

// GetSearchParams defines parameters for GetSearch.
type GetSearchParams struct {
	// Q Query
	Q string `form:"q" json:"q"`
}

// GetVideoIdDownloadFParams defines parameters for GetVideoIdDownloadF.
type GetVideoIdDownloadFParams struct {
	XBotId BotId `json:"X-Bot-Id"`
}

// GetVideoIdFormatsParams defines parameters for GetVideoIdFormats.
type GetVideoIdFormatsParams struct {
	XBotId BotId `json:"X-Bot-Id"`
}
