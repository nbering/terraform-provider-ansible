# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.4] - 2018-04-19
### Changed
- Update vendored dependencies to for Terraform to v0.11.7

## [0.0.3] - 2018-02-25
### Added
- `ansible_group` resource type (requires [nbering/terraform-inventory@1.0.1](https://github.com/nbering/terraform-inventory/releases/tag/v1.0.1) or greater)
## [0.0.2] - 2018-02-25
### Added
- Release build scripts
- CI configuration for Concourse test and release builds

### Fixed
- Releases going forward will have correct binary name for Terraform installation (fixes [#2](https://github.com/nbering/terraform-provider-ansible/issues/2))

[Unreleased]: https://github.com/nbering/terraform-provider-ansible/compare/v0.0.4...HEAD
[0.0.4]: https://github.com/nbering/terraform-provider-ansible/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/nbering/terraform-provider-ansible/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/nbering/terraform-provider-ansible/compare/v0.0.1...v0.0.2
