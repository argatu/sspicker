package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Relay struct {
	IP string `json:"ipv4"`
}

type Location struct {
	Name  string  `json:"desc"`
	Relay []Relay `json:"relays"`
}

type Pops map[string]Location

type DatagramConfig struct {
	Pops Pops `json:"pops"`
}

func Decode() (Pops, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/SteamDatabase/SteamTracking/master/Random/NetworkDatagramConfig.json")
	if err != nil {
		return nil, fmt.Errorf("error fetching network datagram config: %w", err)
	}
	defer resp.Body.Close()

	var config DatagramConfig
	err = json.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error decoding network datagram config: %w", err)
	}

	return config.Pops, nil
}
