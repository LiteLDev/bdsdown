package main

import (
	"github.com/alexflint/go-arg"
	set "github.com/deckarep/golang-set/v2"
	"net/url"
)

var platformMapping = map[string]string{
	"windows":         "serverBedrockWindows",
	"linux":           "serverBedrockLinux",
	"windows-preview": "serverBedrockPreviewWindows",
	"linux-preview":   "serverBedrockPreviewLinux",
}
var acceptedScheme = set.NewSet("http", "https", "file", "version")

type Config struct {
	Confirm       bool    `arg:"-y,--yes" help:"input yes by default if need user confirm"`
	TargetPackage url.URL `arg:"-s,--source" help:"specific package source" default:"version://windows/latest"`
	Procs         int     `arg:"-p,--proc" help:"number of concurrent download" default:"4"`
	// InstallerExclude is only used for the argument parser, use InstallerExcludeSet instead.
	InstallerExclude    []string        `arg:"--exclude" help:"exclude files from the installation"`
	InstallerExcludeSet set.Set[string] `arg:"-"`
}

var config = Config{}

// ParseCommandLine parses command line arguments and returns a Config.
func ParseCommandLine() {

	p := arg.MustParse(&config)

	if config.TargetPackage.String() != "" && !acceptedScheme.Contains(config.TargetPackage.Scheme) {
		p.Fail("invalid scheme, supported schemes: http, https, file, version")
	}

	config.InstallerExcludeSet = set.NewSet("server.properties", "allowlist.json", "permissions.json")
	for _, file := range config.InstallerExclude {
		config.InstallerExcludeSet.Add(file)
	}

}
