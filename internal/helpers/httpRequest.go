package helpers

import (
	"encoding/json"
	"io"
)

func NewRequestFromReader[T any](r io.ReadCloser) *T {
	var req T
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&req)
	if err != nil {
		// ignore err
		return nil
	}

	return &req
}
