package nextcloud

import "encoding/json"

type ConfigSecrets struct {
	InstanceID   string `json:"instanceid"`
	PasswordSalt string `json:"passwordsalt"`
	Secret       string `json:"secret"`
}

func (cfg *ConfigSecrets) Parse(encoded []byte) error {
	return json.Unmarshal(encoded, cfg)
}
