# Change Log
All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

## [1.3.2] - 2020-07-23
### Changed
- Updated Bonsai YAML to include the armv6 build release assets

## [1.3.1] - 2020-02-11
### Changed
- README updates, using the template from https://github.com/sensu-community/check-plugin-template

## [1.3.0] - 2020-02-11
### Changed
- Removed OSX i386 from Bonsai builds

## [1.2.0] - 2020-02-10
### Changed
- Using go modules

### Added
- TLS insecure skip verify support, e.g. `-insecure-skip-verify`

### Fixed
- Fixed error logging/output

## [1.1.6] - 2019-05-21
### Changed
- Fixed Bonsai Asset YAML filter

## [1.1.5] - 2019-01-18
### Changed
- Added Bonsai YAML config file

## [1.1.4] - 2018-12-10
### Changed
- Added InfluxDB line segment count validation (again)

## [1.1.3] - 2018-12-10
### Changed
- Validating InfluxDB tags, no newline and only one =
- Removed InfluxDB line segment count validation

## [1.1.2] - 2018-12-10
### Changed
- Removing newlines (e.g. "\n") from produced tags

## [1.1.1] - 2018-12-10
### Changed
- Dropped the extra trailing newline

## [1.1.0] - 2018-12-10
### Changed
- Validating the number of InfluxDB line segments (" ")
- Fixed linter violations

### Added
- Automated Asset builds with goreleaser and Travis CI

## [1.0.0] - 2018-02-07
### Added
- Exporter basic authentication support (@zsais)
- Exporter authorization header support (@discordianfish)

## [0.0.1] - 2017-10-02
### Added
- Initial release
