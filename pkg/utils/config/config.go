package config

import (
	"io"

	"github.com/go-yaml/yaml"
)

var QipConfig Config

// LoadConfig 設定を読み込む
func LoadConfig(reader io.Reader) (Config, error) {
	if QipConfig.FQDN == "" {
		err := loadConfigFile(reader)
		if err != nil {
			return Config{}, err
		}
	}
	return QipConfig, nil
}

// 設定ファイルをディスクから読み込む
func loadConfigFile(reader io.Reader) error {
	var C Config
	err := yaml.NewDecoder(reader).Decode(&C)
	if err != nil {
		return err
	}
	QipConfig = C
	return nil
}
