package utils

import (
	"testing"

	"github.com/google/uuid"
)

func TestGenAppKeyOrSecret(t *testing.T) {
	u1 := uuid.New()
	u2, _ := uuid.NewRandom()
	t.Log("u1", u1)
	t.Log("u2", u2)

}
