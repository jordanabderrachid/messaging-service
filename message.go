package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

func validateMessage(b []byte) (message, error) {
	var m message

	if err := json.Unmarshal(b, &m); err != nil {
		return m, err
	}

	if m.ID == 0 || m.Data == "" {
		return m, fmt.Errorf("Message has no ID or Data")
	}

	return m, nil
}
