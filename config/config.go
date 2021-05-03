package config

import (
	"flag"

	"github.com/ian-kent/envconf"
)

func DefaultConfig() *Config {
	return &Config{
		AuthFile: "",
	}
}

type Config struct {
	AuthFile string
	WebPath  string

	ProfilingEnabled bool
}

var cfg = DefaultConfig()

func Configure() *Config {

	//sanitize webpath
	//add a leading slash
	if cfg.WebPath != "" && !(cfg.WebPath[0] == '/') {
		cfg.WebPath = "/" + cfg.WebPath
	}

	return cfg
}

func RegisterFlags() {
	flag.StringVar(&cfg.AuthFile, "auth-file", envconf.FromEnvP("MH_AUTH_FILE", "").(string), "A username:bcryptpw mapping file")
	flag.StringVar(&cfg.WebPath, "ui-web-path", envconf.FromEnvP("MH_UI_WEB_PATH", "").(string), "WebPath under which the UI is served (without leading or trailing slashes), e.g. 'mailhog'. Value defaults to ''")
	flag.BoolVar(&cfg.ProfilingEnabled, "profiling-enabled", envconf.FromEnvP("PROFILING_ENABLED", false).(bool), "When set, profiling via pprof is enabled")
}
