# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.4.0] - 2023-12-25

### Added

- Local archive support.

## [0.3.1] - 2023-04-05

### Added

- Add classical command parameter support

## [0.3.0] - 2023-03-30

### Added

- `-e` or `--exclude` option to exclude existing files from the installation and default value is `[server.properties allowlist.json permissions.json]` (#1)
- `-v` or `--version` flag to specify the version of BDS to install. If not specified, the latest release(preview if -p specified) version will be used.
- `--help` flag to print help information

### Changed

- `-preview` flag to `-p` or `--preview`
- Use [akamensky/argparse](https://github.com/akamensky/argparse) to parse arguments

## [0.2.0] - 2023-02-17

### Added

- `-y` flag to skip confirmation

## [0.1.0] - 2023-02-09

### Added

- Basic functionality

[0.4.0]: https://github.com/LiteLDev/bdsdown/compare/v0.3.1...v0.4.0
[0.3.1]: https://github.com/LiteLDev/bdsdown/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/LiteLDev/bdsdown/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/LiteLDev/bdsdown/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/LiteLDev/bdsdown/releases/tag/v0.1.0
