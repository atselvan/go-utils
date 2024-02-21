package structutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name string `mapstructure:"name_ms" json:"name_j,omitempty" yaml:"name_y"`
}

func TestGetFieldTagValue(t *testing.T) {
	user := new(User)
	tag := GetFieldTagValue(user, &user.Name)
	assert.Equal(t, "name_ms", tag)
}

func TestGetJsonFieldTagValue(t *testing.T) {
	user := new(User)
	tag := GetJsonFieldTagValue(user, &user.Name)
	assert.Equal(t, "name_j", tag)
}

func TestGetYamlFieldTagValue(t *testing.T) {
	user := new(User)
	tag := GetYamlFieldTagValue(user, &user.Name)
	assert.Equal(t, "name_y", tag)
}
