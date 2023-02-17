package main

import (
	"flag"
	"fmt"
	"os"

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
	fmt.Println("Before using this software, please read: ")
	fmt.Println("- Minecraft End User License Agreement   https://minecraft.net/terms")
	fmt.Println("- Microsoft Privacy Policy               https://go.microsoft.com/fwlink/?LinkId=521839")

	usePreview := false
	skipAgree := false
	flagSet := flag.NewFlagSet("bdsdown", flag.ExitOnError)
	flagSet.BoolVar(&usePreview, "preview", false, "Use preview version")
	flagSet.BoolVar(&skipAgree, "y", false, "Skip the agreement")
	flagSet.Usage = func() {
		fmt.Println("Usage: bdsdown [options] [version]")
		fmt.Println("Options:")
		flagSet.PrintDefaults()
	}
	flagSet.Parse(os.Args[1:])
	if flagSet.NArg() > 1 {
		fmt.Println(ColorRed + "ERROR: Too many arguments." + ColorReset)
		flagSet.Usage()
		return
	}

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

	if usePreview {
		fmt.Println(ColorYellow + "Using preview version." + ColorReset)
	}
	if flagSet.NArg() == 1 {
		ver := flagSet.Arg(0)
		err := utils.Install(ver, usePreview)
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
		err = utils.Install(ver, usePreview)
		if err != nil {
			fmt.Println(ColorRed+"ERROR:", err, ColorReset)
		}
		fmt.Println(ColorGreen + "Install complete." + ColorReset)
		return
	}
}
