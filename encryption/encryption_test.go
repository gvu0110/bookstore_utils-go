package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMD5(t *testing.T) {
	hash := GetMD5("123")
	assert.EqualValues(t, "202cb962ac59075b964b07152d234b70", hash)
}
