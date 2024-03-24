package helpers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
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

func GetUuidFromPath(r *http.Request, path string) (*uuid.UUID, error) {
	id := r.PathValue(path)
	guid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &guid, nil
}
