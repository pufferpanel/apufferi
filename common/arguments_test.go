package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceTokens(t *testing.T) {
	mapping := createSourceMap()

	resultTest := ReplaceTokens("TEST ${val1}", mapping)
	assert.Equal(t, "TEST RESULT1", resultTest)

	resultTest = ReplaceTokens("TEST val1", mapping)
	assert.Equal(t, "TEST val1", resultTest)

	resultTest = ReplaceTokens("TEST val1", mapping)
	assert.Equal(t, "TEST val1", resultTest)
}

func createSourceMap() map[string]interface{} {
	source := make(map[string]interface{})

	source["val1"] = "RESULT1"
	source["value2"] = "RESULT2"
	source["1234567"] = "RESULT3"
	source["val123"] = "RESULT4"
	source["int"] = 436

	return source
}
