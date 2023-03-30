# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.0] - 2023-03-30
### Added
- `-e` or `--exclude` option to exclude existing files from the installation and default value is `[server.properties allowlist.json permissions.json]`
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


[unreleased]: https://github.com/LiteLDev/Lip/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/LiteLDev/Lip/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/LiteLDev/Lip/releases/tag/v0.1.0