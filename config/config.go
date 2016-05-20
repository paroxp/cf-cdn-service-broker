package config

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Settings struct {
	Port               string `envconfig:"port";default:"3000"`
	Username           string `envconfig:"username";required:"true"`
	Password           string `envconfig:"password";required:"true"`
	DatabaseUrl        string `envconfig:"database_url";required:"true"`
	Email              string `envconfig:"email";required:"true"`
	AcmeUrl            string `envconfig:"acme_url";required:"true"`
	Bucket             string `envconfig:"bucket";required:"true"`
	AwsAccessKeyId     string `envconfig:"aws_access_key_id";required:"true"`
	AwsSecretAccessKey string `envconfig:"aws_secret_access_key";required:"true"`
	AwsDefaultRegion   string `envconfig:"aws_default_region";required:"true"`
}

func NewSettings() (Settings, error) {
	var settings Settings
	err := envconfig.Process("cdn", &settings)
	if err != nil {
		return Settings{}, err
	}
	return settings, nil
}

func Connect(settings Settings) (*gorm.DB, error) {
	return gorm.Open("postgres", settings.DatabaseUrl)
}
