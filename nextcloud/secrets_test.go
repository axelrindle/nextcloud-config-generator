package nextcloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigSecretsParse(t *testing.T) {
	payload := `{"instanceid": "foo", "passwordsalt": "bar", "secret": "baz"}`

	config := &ConfigSecrets{}
	config.Parse([]byte(payload))

	assert.Equal(t, config.InstanceID, "foo")
	assert.Equal(t, config.PasswordSalt, "bar")
	assert.Equal(t, config.Secret, "baz")
}
