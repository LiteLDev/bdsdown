package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/jasonzyt/bdsdownloader/utils"
)

const (
	ColorRed    = "\033[91m"
	ColorGreen  = "\033[92m"
	ColorYellow = "\033[93m"
	ColorBlue   = "\033[94m"
	ColorReset  = "\033[0m"
)

func main() {
	fmt.Println(" BDS Downloader | Distributed under the MIT License. ")
	fmt.Println("=====================================================")

	parser := argparse.NewParser("bdsdown", "Download and install BDS.")
	usePreviewPtr := parser.Flag("p", "preview", &argparse.Options{Required: false, Help: "Use preview version"})
	skipAgreePtr := parser.Flag("y", "yes", &argparse.Options{Required: false, Help: "Skip the agreement"})
	excludedFilesPtr := parser.StringList("e", "exclude", &argparse.Options{Required: false, Help: "Exclude existing files from the installation", Default: []string{"server.properties", "allowlist.json", "permissions.json"}})
	targetVersionPtr := parser.String("v", "version", &argparse.Options{Required: false, Help: "The version of BDS to install. If not specified, the latest release(preview if -p specified) version will be used."})
	parser.Parse(os.Args)

	usePreview := *usePreviewPtr
	skipAgree := *skipAgreePtr
	excludedFiles := make([]string, 0)
	targetVersion := *targetVersionPtr
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

	fmt.Println("Before using this software, please read: ")
	fmt.Println("- Minecraft End User License Agreement   https://minecraft.net/terms")
	fmt.Println("- Microsoft Privacy Policy               https://go.microsoft.com/fwlink/?LinkId=521839")
	fmt.Print("Please enter y if you agree with the above terms: ")
	var agree string
	if skipAgree {
		agree = "y"
		fmt.Println(agree)
	} else {
		fmt.Scanln(&agree)
	}
	if agree != "y" {
		fmt.Println(ColorYellow + "You must agree with the above terms to use this software." + ColorReset)
		return
	}
	fmt.Println("=====================================================")

	if len(excludedFiles) > 0 {
		fmt.Println("The following files will be excluded from installation: ", excludedFiles)
	}

	if usePreview {
		fmt.Println(ColorYellow + "Using preview version." + ColorReset)
	}
	if targetVersion != "" {
		err := utils.Install(targetVersion, usePreview, excludedFiles)
		if err != nil {
			fmt.Println(ColorRed+"ERROR:", err, ColorReset)
			return
		}
		fmt.Println(ColorGreen + "Install complete." + ColorReset)
		return
	} else {
		var ver string
		var err error
		if usePreview {
			ver, err = utils.GetLatestPreviewVersion()
		} else {
			fmt.Println("No version specified, using latest release version.")
			ver, err = utils.GetLatestReleaseVersion()
		}
		if err != nil {
			fmt.Println(ColorRed+"ERROR:", err)
			return
		}
		fmt.Println("Latest version: " + ColorBlue + ver + ColorReset)
		err = utils.Install(ver, usePreview, excludedFiles)
		if err != nil {
			fmt.Println(ColorRed+"ERROR:", err, ColorReset)
		}
		fmt.Println(ColorGreen + "Install complete." + ColorReset)
		return
	}
}
