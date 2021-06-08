package rctfgolf

import (
	"fmt"
	"os"
	"time"

	"github.com/redpwn/rctf-golf/internal/api"
)

func unixMillisToTime(millis int64) time.Time {
	return time.Unix(millis/1000, millis%1000*int64(time.Millisecond))
}

func CalculateWithClient(c *api.APIClient, chall string, mapper func(time.Duration) interface{}) (interface{}, error) {
	if debug, ok := os.LookupEnv("RCTF_GOLF_DEBUG"); ok {
		elapsed, err := time.ParseDuration(debug)
		if err != nil {
			panic(fmt.Errorf("Illegal debug elapsed time value: %w", err))
		}
		if int64(elapsed) < 0 {
			panic("Cannot set debug with negative elapsed time!")
		}
		return mapper(elapsed), nil
	}

	clientConfig, err := c.GetClientConfig()
	if err != nil {
		return nil, err
	}
	start := unixMillisToTime(clientConfig.StartTime)
	solves, err := c.GetChallengeSolves(chall, api.GetChallengeSolvesParams{
		Limit:  1,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}
	var current time.Time
	if len(solves) > 0 {
		current = unixMillisToTime(solves[0].CreatedAt)
	} else {
		current = time.Now()
	}
	elapsed := current.Sub(start)
	return mapper(elapsed), nil
}

func Calculate(baseURL string, chall string, mapper func(time.Duration) interface{}) (interface{}, error) {
	return CalculateWithClient(api.NewClient(baseURL), chall, mapper)
}
