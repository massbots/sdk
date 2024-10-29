package sdk_test

import (
	"os"
	"testing"

	massbots "go.massbots.xyz/sdk"
)

func TestExample(t *testing.T) {
	mb := massbots.New(os.Getenv("BOT_ID"), os.Getenv("TOKEN"))

	balance, err := mb.Balance()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Balance: %d", balance)

	stats, err := mb.Stats()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Stats: %v", stats)

	video, err := mb.Video("R6rhxJjVNCU")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Video: %s", video.Title)

	channel, err := mb.Channel(video.ChannelId)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Channel: %s", channel.Title)

	formats, err := mb.VideoFormats(video.Id)
	if err != nil {
		t.Fatal(err)
	}

	cached := formats.Cached()
	t.Logf("Cached formats: %v", cached)

	if len(cached) > 0 {
		download, err := mb.VideoDownload(video.Id, cached[0].Format)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Telegram file_id: %s", download.FileId)
	}
}
