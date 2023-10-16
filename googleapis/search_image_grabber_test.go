package googleapis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	apiKey         = "testKey"
	searchEngineID = "testId"
)

func TestMakeUrlReplacesQuerySpacesWithPlus(t *testing.T) {
	asserts := assert.New(t)

	sut := New(apiKey, searchEngineID)

	res := sut.sanitizeQueryUrl("cute kittens are jumping", 0)

	asserts.NotContains(res, "cute kittens are jumping")
	asserts.Contains(res, "cute")
	asserts.Contains(res, "kittens")
	asserts.Contains(res, "are")
	asserts.Contains(res, "jumping")
}
