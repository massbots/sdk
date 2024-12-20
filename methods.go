package sdk

import "time"

func (a API) Balance() (int, error) {
	balance, err := get[Balance](a, "me/balance")
	if err != nil {
		return 0, err
	}
	return balance.Balance, nil
}

func (a API) Stats() ([]RequestStat, error) {
	return get[[]RequestStat](a, "me/stats")
}

func (a API) Video(id string) (Video, error) {
	return get[Video](a, "video/"+id)
}

func (a API) Channel(id string) (Channel, error) {
	return get[Channel](a, "channel/"+id)
}

func (a API) Search(q string) ([]Video, error) {
	return a.SearchVideos(q)
}
func (a API) ChannelFeed(id string) (ChannelFeed, error) {
	return get[ChannelFeed](a, "channel/"+id+"/feed")
}

func (a API) SearchVideos(q string) ([]Video, error) {
	return get[[]Video](a, "search?q="+q+"&kind=video")
}

func (a API) SearchChannels(q string) ([]Channel, error) {
	return get[[]Channel](a, "search?q="+q+"&kind=channel")
}

func (a API) VideoFormats(id string) (VideoFormats, error) {
	return get[VideoFormats](a, "video/"+id+"/formats")
}

func (a API) VideoDownload(id, format string) (DownloadResult, error) {
	return get[DownloadResult](a, "video/"+id+"/download/"+format)
}

func (a API) PollDownload(id, format string, d time.Duration) (DownloadResult, error) {
	if d < 1*time.Second {
		d = 1 * time.Second
	}
	for {
		result, err := a.VideoDownload(id, format)
		if err != nil {
			return result, err
		}
		if result.Status == "ready" || result.Status == "failed" {
			return result, nil

		}
		time.Sleep(d)
	}
}
