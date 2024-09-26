package utils

import (
	"github.com/google/uuid"
)

func UUID() string {
	id, _ := uuid.NewRandom()
	return id.String()
}

func MustUUID(uid string, val []byte) string {
	namespace := uuid.MustParse(uid)
	return uuid.NewSHA1(namespace, val).String()
}
