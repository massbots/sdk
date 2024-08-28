package sdk_test

import (
	"os"
	"testing"

	massbots "go.massbots.xyz/sdk"
)

func TestExample(t *testing.T) {
	mb := massbots.New(os.Getenv("TOKEN"))

	balance, err := mb.Balance()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Balance: %d", balance)

	video, err := mb.YoutubeVideo("R6rhxJjVNCU")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Video: %s", video.Title)

	channel, err := mb.YoutubeChannel(video.ChannelId)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Channel: %s", channel.Title)

	formats, err := mb.Video(video.Id)
	if err != nil {
		t.Fatal(err)
	}

	cached := formats.Cached()

	t.Logf("Cached formats: %v", cached)

	if len(cached) > 0 {
		download, err := mb.VideoCached(video.Id, cached[0].Format)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Telegram file_id: %s", download.FileId)
	}

}
