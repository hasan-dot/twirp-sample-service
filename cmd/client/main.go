package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hasan-dot/twirp-service/cmd/server/proto"
	"github.com/twitchtv/twirp"
)

func main() {
	client := proto.NewHaberdasherJSONClient("http://localhost:8080", &http.Client{})

	var (
		hat *proto.Hat
		err error
	)
	for i := 0; i < 5; i++ {
		hat, err = client.MakeHat(context.Background(), &proto.Size{Inches: 12})
		if err != nil {
			if twerr, ok := err.(twirp.Error); ok {
				if twerr.Meta("retryable") != "" {
					// Log the error and go again.
					log.Printf("got error %q, retrying", twerr)
					continue
				}
			}
			// This was some fatal error!
			log.Fatal(err)
		}
	}
	fmt.Printf("%+v", hat)
}
