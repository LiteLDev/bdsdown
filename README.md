# bdsdown
Downloader for Minecraft BDS

## Usage
```
usage: bdsdown [-h|--help] [-p|--preview] [-y|--yes] [-e|--exclude "<value>"
               [-e|--exclude "<value>" ...]] [-v|--version "<value>"]

               Download and install BDS.

Arguments:

  -h  --help     Print help information
  -p  --preview  Use preview version
  -y  --yes      Skip the agreement
  -e  --exclude  Exclude existing files from the installation. Default:
                 [server.properties allowlist.json permissions.json]
  -v  --version  The version of BDS to install. If not specified, the latest
                 release(preview if -p specified) version will be used.
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Acknowledgements

### Third-party libraries used

- [akamensky/argparse](https://github.com/akamensky/argparse) - MIT License
- [schollz/progressbar](https://github.com/schollz/progressbar) - MIT License
