# Gopher-Panic Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.0] - 2024-01-24

### Added

- GNU Standard error format
- Possibility to change between legacy and GNU Standard error format
- Possibility to remove file data from error format
- Documentation Examples
- Code documentation

### Changed

- Set GNU stardard as default for error interface without file data
- Add lib examples into Readme

## [0.2.0]

### Added

- Code struct to identify more easily the error kind
- FormatJSON can now be indented

### Changed

- Enhance Error and Trace message
- New Error function use now a params for Traces
- Convert Error into Error pointer for a better usability

## [0.1.1]

### Fixed

- Package name

## [0.1.0]

### Added

- Error system
- ErrorBuilder system
- Trace system
- Documentation
- CI

[unreleased]: https://github.com/ulphidius/gopherpanic/compare/v0.3.0...master
[0.3.0]: https://github.com/ulphidius/gopherpanic/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/ulphidius/gopherpanic/compare/v0.1.1...v0.2.0
[0.1.1]: https://github.com/ulphidius/gopherpanic/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/ulphidius/gopherpanic/compare/v0.1.0
