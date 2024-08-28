package sdk

import "time"

func (a API) Balance() (int, error) {
	balance, err := get[Balance](a, "balance")
	if err != nil {
		return 0, err
	}
	return balance.Balance, nil
}

func (a API) YoutubeVideo(id string) (YoutubeVideo, error) {
	return get[YoutubeVideo](a, "yt/video/"+id)
}

func (a API) YoutubeChannel(id string) (YoutubeChannel, error) {
	return get[YoutubeChannel](a, "yt/channel/"+id)
}

func (a API) Video(id string) (Video, error) {
	return get[Video](a, "dl/video/"+id)
}

func (a API) VideoCached(id, format string) (DownloadFileID, error) {
	return get[DownloadFileID](a, "dl/video/"+id+"/cached/"+format)
}

func (a API) VideoDownload(id, format string) (DownloadResult, error) {
	return get[DownloadResult](a, "dl/video/"+id+"/download/"+format)
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
