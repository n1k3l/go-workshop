package config

// Server config
type Server struct {
	JokeURL    string `yaml:"joke-url" env:"JOKE_URL"`
	CustomJoke string `env:"CUSTOM_JOKE"`
}
