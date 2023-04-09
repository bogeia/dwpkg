package systemx

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetSystemInfo(t *testing.T) {
	item, err := GetSystemInfo()
	assert.NoError(t, err)
	assert.NotNil(t, item)

	data, _ := json.MarshalIndent(item, "", "   ")
	t.Log("\n", string(data))
}
