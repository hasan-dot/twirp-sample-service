// Copyright 2018 Twitch Interactive, Inc.  All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the License is
// located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/hasan-dot/twirp-service/cmd/server/proto"
	"github.com/twitchtv/twirp"
	"github.com/twitchtv/twirp/hooks/statsd"
)

type randomHaberdasher struct{}

func (h *randomHaberdasher) MakeHat(ctx context.Context, size *proto.Size) (*proto.Hat, error) {
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("Inches", "I can't make a hat that small!")
	}
	colors := []string{"white", "black", "brown", "red", "blue"}
	names := []string{"bowler", "baseball cap", "top hat", "derby"}
	return &proto.Hat{
		Size:  size.Inches,
		Color: colors[rand.Intn(len(colors))],
		Name:  names[rand.Intn(len(names))],
	}, nil
}

func main() {
	hook := statsd.NewStatsdServerHooks(LoggingStatter{os.Stderr})
	server := proto.NewHaberdasherServer(&randomHaberdasher{}, hook)
	log.Println("Server is running. You can curl the following endpoint:")
	log.Println("curl -X POST -H 'Content-Type: application/json' -d '{\"inches\": 7}' http://localhost:8080/twirp/service.Haberdasher/MakeHat | jq")
	log.Fatal(http.ListenAndServe(":8080", server))
}
