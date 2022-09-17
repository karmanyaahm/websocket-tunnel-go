package config

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	err := ParseFile("../example.toml")
	assert.NoError(t, err)
	b, err := json.Marshal(MyConfig)
	assert.NoError(t, err)
	fmt.Println(string(b))
}
