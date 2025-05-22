package helper

import "github.com/google/uuid"

func PtrString(s string) *string {
	return &s
}

func PtrFloat64(f float64) *float64 {
	return &f
}

func PtrUUID(id uuid.UUID) *uuid.UUID {
	return &id
}
