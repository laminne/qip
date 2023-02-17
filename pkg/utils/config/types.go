package config

type Config struct {
	FQDN string `yaml:"fqdn"`
	Port string `yaml:"port"`
	DB   DB     `yaml:"db"`
	Meta Meta   `yaml:"meta"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"name"`
}

type Meta struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Maintainer  Maintainer `yaml:"maintainer"`
}

type Maintainer struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}
