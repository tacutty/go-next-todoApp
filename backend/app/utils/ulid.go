package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// GeneULIDString is a function that generates a ULID string
// @return string, error
func GeneULIDString() (string, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())

	id, err := ulid.New(ms, entropy)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
