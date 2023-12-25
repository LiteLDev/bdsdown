package utils

type Config struct {
	UsePreview    bool
	SkipAgree     bool
	ClearCache    bool
	UseCache      bool
	CacheDir      string
	TargetVersion string
	ExcludedFiles []string
}

var globalConfig Config

func SetConfig(config Config) {
	globalConfig = config
}

func GetConfig() Config {
	return globalConfig
}

func SetTargetVersion(version string) {
	globalConfig.TargetVersion = version
}
