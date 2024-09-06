# massbots.dl 
[![Docs](https://img.shields.io/badge/api-reference-lightgray.svg)](https://api.massbots.xyz/docs) [![GoDoc](https://pkg.go.dev/badge/go.massbots.xyz/sdk)](https://pkg.go.dev/go.massbots.xyz/sdk) [![PyPI version](https://badge.fury.io/py/massbots.svg)](https://badge.fury.io/py/massbots) [![Better Stack Badge](https://uptime.betterstack.com/status-badges/v1/monitor/1iew8.svg)](https://uptime.massbots.xyz)

## [SDK: Go](go/README.md)

> `$ go get go.massbots.xyz/sdk`

```go
package main

import (
	"os"
	"fmt"
)

import (
	massbots "go.massbots.xyz/sdk"
)

func main() {
	mb := massbots.New(
		os.Getenv("BOT_ID"),
		os.Getenv("TOKEN"), 
	)

	video, _ := mb.Video("dQw4w9WgXcQ")
	fmt.Println("Video:", video.Title)

	channel, _ := mb.Channel(video.ChannelID)
	fmt.Println("Channel:", channel.Title)
}
```

## [SDK: Python](python/README.md)

> `$ pip3 install massbots`

```python
import os
import massbots


mb = massbots.Api(os.environ["TOKEN"])

video = mb.video("dQw4w9WgXcQ")
print(f"Video: {video.title}")

channel = mb.channel(video.channel_id)
print(f"Channel: {channel.title}")
```
