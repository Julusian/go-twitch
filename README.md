# go-twitch

# Forked from [go-twitch](https://github.com/mrshankly/go-twitch)

[![travis-ci status](https://api.travis-ci.org/julusian/go-twitch.png)](https://travis-ci.org/julusian/go-twitch)
[![Coverage Status](https://coveralls.io/repos/github/Julusian/go-twitch/badge.svg?branch=master)](https://coveralls.io/github/Julusian/go-twitch?branch=master)

Go library for accessing v5 of the [Twitch-API](https://dev.twitch.tv/docs/).

**This is still a work in progress.**

Currently only GET requests are implemented.

## Usage

To install `go-twitch` run the command:

```bash
$ go get github.com/julusian/go-twitch
```

Full docs at (GoDocs)[https://godoc.org/github.com/julusian/go-twitch]

Here's an example program that gets the top 10 twitch games:

```go
package main

import (
	"fmt"
	"github.com/julusian/go-twitch/twitch"
	"github.com/julusian/go-twitch/twitchapi"
	"log"
	"net/http"
)

func main() {
	client := twitchapi.NewClient(&http.Client{}, os.Getenv("CLIENT_ID"))
	opt := &twitch.ListOptions{
		Limit:  10,
		Offset: 0,
	}

	games, err := client.Games.Top(opt)
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range games.Top {
		fmt.Printf("%d - %s (%d)\n", i+1, s.Game.Name, s.Viewers)
	}
}
```

### Authentication

**TODO**

## License

All files under this repository fall under the MIT License (see the file LICENSE).
