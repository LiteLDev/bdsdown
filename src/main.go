package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jasonzyt/bdsdownloader/utils"
)

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
)

func main() {
	fmt.Println(" BDS Downloader | Distributed under the MIT License. ")
	fmt.Println("=====================================================")
	fmt.Println("Before using this software, please read: ")
	fmt.Println("- Minecraft End User License Agreement   https://minecraft.net/terms")
	fmt.Println("- Microsoft Privacy Policy               https://go.microsoft.com/fwlink/?LinkId=521839")

	fmt.Print("Please enter y if you agree with the above terms: ")
	var agree string
	fmt.Scanln(&agree)
	if agree != "y" {
		fmt.Println("You must agree with the above terms to use this software.")
		return
	}
	fmt.Println("=====================================================")

	usePreview := false
	flagSet := flag.NewFlagSet("bdsdownloader", flag.ExitOnError)
	flagSet.BoolVar(&usePreview, "preview", false, "Use preview version")
	flagSet.Usage = func() {
		fmt.Println("Usage: bdsdownloader [options] [version]")
		fmt.Println("Options:")
		flagSet.PrintDefaults()
	}
	flagSet.Parse(os.Args[1:])
	if usePreview {
		fmt.Println(ColorYellow + "Using preview version.")
	}
	if flagSet.NArg() > 1 {
		fmt.Println(ColorRed + "ERROR: Too many arguments.")
		flagSet.Usage()
		return
	}
	if flagSet.NArg() == 1 {
		ver := flagSet.Arg(0)
		fmt.Println("Downloading BDS v", ver)
		err := utils.DownloadVersion(ver, usePreview)
		if err != nil {
			fmt.Println(ColorRed+"ERROR: ", err)
			return
		}
		fmt.Println(ColorGreen + "Download complete.")
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
			fmt.Println(ColorRed+"ERROR: ", err)
			return
		}
		fmt.Println("Downloading BDS v", ver)
		err = utils.DownloadVersion(ver, usePreview)
		if err != nil {
			fmt.Println(ColorRed+"ERROR: ", err)
			return
		}
		fmt.Println(ColorGreen + "Download complete.")
		return
	}
}
