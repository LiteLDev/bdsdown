package main

import (
	"bytes"
	"net/url"
	"path"
	"strings"
	"text/template"
)

const (
	CACHE_DIR         = "./.cache/bdsdown"
	VERSION_QUERY_URL = "https://www.minecraft.net/en-us/download/server/bedrock"
)

var VERSION_TEMPLATE = map[string]*template.Template{
	"windows":         template.Must(template.New("windows").Parse("https://minecraft.azureedge.net/bin-win/bedrock-server-{{.}}.zip")),
	"linux":           template.Must(template.New("linux").Parse("https://minecraft.azureedge.net/bin-linux/bedrock-server-{{.}}.zip")),
	"windows-preview": template.Must(template.New("windows-preview").Parse("https://minecraft.azureedge.net/bin-win-preview/bedrock-server-{{.}}.zip")),
	"linux-preview":   template.Must(template.New("linux-preview").Parse("https://minecraft.azureedge.net/bin-linux-preview/bedrock-server-{{.}}.zip")),
}

func main() {
	ParseCommandLine()

	log.Info("This software only provides a convenient way to download and install Minecraft Bedrock Edition server")
	log.Info("You must agree with the Minecraft End User License Agreement and Privacy Policy to continue")
	log.Info("- Minecraft End User License Agreement   https://minecraft.net/terms")
	log.Info("- Microsoft Privacy Policy               https://go.microsoft.com/fwlink/?LinkId=521839")
	if !UserConfirm("Do you agree with the above terms? (y/n): ") {
		log.Info("You must agree with the above terms to use this software.")
		return
	}

	processHttpPackage := func(u *url.URL) {
		_, f := path.Split(u.Path)
		f = path.Join(CACHE_DIR, f)
		err := DownloadFile(u.String(), f)
		if err != nil {
			log.Fatal(err)
		}
		err = UnzipPackage(f, ".")
		if err != nil {
			log.Fatal(err)
		}

	}
	if config.TargetPackage.Scheme == "file" {
		err := UnzipPackage(path.Join(config.TargetPackage.Host, config.TargetPackage.Path), ".")
		if err != nil {
			log.Fatal(err)
		}
	} else if config.TargetPackage.Scheme == "http" || config.TargetPackage.Scheme == "https" {
		processHttpPackage(&config.TargetPackage)
	} else if config.TargetPackage.Scheme == "version" {
		platform := config.TargetPackage.Host
		version := strings.TrimLeft(config.TargetPackage.Path, "/")
		if version == "" || version == "latest" {
			log.Infof("fetching latest version for platform %s", platform)
			versions, err := FetchVersions(VERSION_QUERY_URL)
			if err != nil {
				log.Fatal(err)
			}
			u, ok := versions[platformMapping[platform]]
			if !ok {
				log.Fatalf("failed to find version for platform %s", platform)
			}
			processHttpPackage(u)
		} else {
			buf := bytes.NewBuffer([]byte{})
			VERSION_TEMPLATE[platform].Execute(buf, version)
			u, _ := url.Parse(buf.String())
			processHttpPackage(u)
		}
	} else {
		log.Fatalf("unsupported scheme for target package %s", config.TargetPackage.Scheme)
	}

}
