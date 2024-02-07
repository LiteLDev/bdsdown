package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/liteldev/bdsdown/logger"
	"github.com/liteldev/bdsdown/utils"
)

func getDefaultCacheDir() string {
	currentUser, err := user.Current()
	if err != nil {
		return "./.cache/bdsdown"
	}

	return path.Join(currentUser.HomeDir, ".cache", "bdsdown")
}

func main() {
	fmt.Println(" BDS Downloader | Distributed under the MIT License. ")
	fmt.Println("=====================================================")

	parser := argparse.NewParser("bdsdown", "Download and install BDS.")
	usePreviewPtr := parser.Flag("p", "preview", &argparse.Options{Required: false, Help: "Use preview version"})
	skipAgreePtr := parser.Flag("y", "yes", &argparse.Options{Required: false, Help: "Skip the agreement"})
	clearCachePtr := parser.Flag("", "clear-cache", &argparse.Options{Required: false, Help: "Clear the cache directory and exit"})
	noCachePtr := parser.Flag("", "no-cache", &argparse.Options{Required: false, Help: "Clear the default cache directory and exit"})
	cacheDirPtr := parser.String("", "cache-dir", &argparse.Options{Required: false, Help: "The directory to store downloaded files", Default: getDefaultCacheDir()})
	excludedFilesPtr := parser.StringList("e", "exclude", &argparse.Options{Required: false, Help: "Exclude existing files from the installation", Default: []string{"server.properties", "allowlist.json", "permissions.json"}})
	targetVersionPtr := parser.String("v", "version", &argparse.Options{Required: false, Help: "The version of BDS to install. If not specified, the latest release(preview if -p specified) version will be used."})
	parser.Parse(os.Args)

	usePreview := *usePreviewPtr
	skipAgree := *skipAgreePtr
	clearCache := *clearCachePtr
	excludedFiles := make([]string, 0)
	targetVersion := *targetVersionPtr
	cacheDir := *cacheDirPtr
	useCache := !*noCachePtr

	if targetVersion == "" {
		for i := 0; i < len(os.Args); i++ {
			if strings.Contains(os.Args[i], "\\") || strings.Contains(os.Args[i], "/") || strings.Contains(os.Args[i], "-") {
				continue
			}
			targetVersion = os.Args[i]
		}
	}

	for _, file := range *excludedFilesPtr {
		if _, err := os.Stat(file); err == nil {
			excludedFiles = append(excludedFiles, file)
		}
	}

	utils.SetConfig(utils.Config{
		UsePreview:    usePreview,
		SkipAgree:     skipAgree,
		ClearCache:    clearCache,
		UseCache:      useCache,
		CacheDir:      cacheDir,
		ExcludedFiles: excludedFiles,
		TargetVersion: targetVersion,
	})

	if clearCache {
		err := os.RemoveAll(cacheDir)
		if err != nil {
			logger.LogError(err)
			return
		}
		logger.LogSuccess("Cache cleared.")
		return
	}

	logger.Log("Before using this software, please read: ")
	logger.Log("- Minecraft End User License Agreement   https://minecraft.net/terms")
	logger.Log("- Microsoft Privacy Policy               https://go.microsoft.com/fwlink/?LinkId=521839")
	fmt.Print("Please enter y if you agree with the above terms: ")
	var agree string
	if skipAgree {
		agree = "y"
		logger.Log(agree)
	} else {
		fmt.Scanln(&agree)
	}
	if agree != "y" {
		logger.LogWarning("You must agree to the terms to use this software.")
		return
	}
	logger.Log("=====================================================")

	if len(excludedFiles) > 0 {
		logger.Log("The following files will be excluded from installation: ", excludedFiles)
	}

	if usePreview {
		logger.LogWarning("Using preview version.")
	}

	if targetVersion != "" {
		err := utils.Install()
		if err != nil {
			logger.LogError(err)
			return
		}
		logger.LogSuccess("Install complete.")
	} else {
		var version string
		var err error
		if usePreview {
			version, err = utils.GetLatestPreviewVersion()
		} else {
			logger.Log("No version specified, using latest release version.")
			version, err = utils.GetLatestReleaseVersion()
		}
		if err != nil {
			logger.LogError(err)
			return
		}
		logger.Log("Latest version:", logger.ColorBlue+version+logger.ColorReset)
		utils.SetTargetVersion(version)

		err = utils.Install()
		if err != nil {
			logger.LogError(err)
			// TODO: Add rollback
			// TODO: Add tips for common errors(clearing cache, etc.)
			return
		}
		logger.LogSuccess("Install complete.")
	}
}
