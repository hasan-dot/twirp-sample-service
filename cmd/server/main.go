package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

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

type LoggingStatter struct {
	io.Writer
}

func (ls LoggingStatter) Inc(metric string, val int64, rate float32) error {
	_, err := fmt.Fprintf(ls, "incr %s: %d @ %f\n", metric, val, rate)
	return err
}

func (ls LoggingStatter) TimingDuration(metric string, val time.Duration, rate float32) error {
	_, err := fmt.Fprintf(ls, "time %s: %s @ %f\n", metric, val, rate)
	return err
}
