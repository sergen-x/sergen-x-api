package utils

import "github.com/jaevor/go-nanoid"

func GenerateUUID(length int) (string, error) {
	idGenerator, err := nanoid.CustomASCII("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", length)
	if err != nil {
		return "", err
	}
	return idGenerator(), nil
}
