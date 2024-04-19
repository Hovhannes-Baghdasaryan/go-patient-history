package helper

import "github.com/google/uuid"

func IsValidUUID(u string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(u)
	if err != nil {
		return uuid.UUID{}, err
	}

	return parsed, nil
}
