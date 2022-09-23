package config

type Config struct {
	IsDebug bool `env:"IS_DEBUG" env-default:"false"`
	Local   struct {
		UiFontColor  string `env:"UI_FONT_COLOR" env-default:"red"`
		MouseEnabled bool   `env:"MOUSE_ENABLED" env-default:"true"`
	}
}
