# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- `variable_priority` added to all resources, to be interpreted by inventory script to merge higher priority vars over lower

## [0.0.5] - 2019-05-20
### Added
- back-ported `ansible_host_var` and `ansible_group_var` from v1.0.1

### Changed
- Moved to `go mod` dependency management
- Updated dependencies to Terraform v0.11.14

## [1.0.1] - 2019-05-20
### Added
- `ansible_host_var` and `ansible_group_var` resources for setting individual variables

## [1.0.0] - 2019-05-05
### Changed
- Update dependencies for Terraform v0.12-beta2 - breaking due to plugin version change

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

[Unreleased]: https://github.com/nbering/terraform-provider-ansible/compare/v1.0.1...HEAD
[0.0.5]: https://github.com/nbering/terraform-provider-ansible/compare/v0.0.4...v0.0.5
[1.0.1]: https://github.com/nbering/terraform-provider-ansible/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/nbering/terraform-provider-ansible/compare/v0.0.4...v1.0.0
[0.0.4]: https://github.com/nbering/terraform-provider-ansible/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/nbering/terraform-provider-ansible/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/nbering/terraform-provider-ansible/compare/v0.0.1...v0.0.2
