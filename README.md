# GPG Terraform Provider [![Build Status](https://travis-ci.com/invidian/terraform-provider-gpg.svg?branch=master)](https://travis-ci.com/invidian/terraform-provider-gpg) [![Maintainability](https://api.codeclimate.com/v1/badges/bc27e5bc32a9b40efaa0/maintainability)](https://codeclimate.com/github/invidian/terraform-provider-gpg/maintainability) [![codecov](https://codecov.io/gh/invidian/terraform-provider-gpg/branch/master/graph/badge.svg)](https://codecov.io/gh/invidian/terraform-provider-gpg) [![Go Report Card](https://goreportcard.com/badge/github.com/invidian/terraform-provider-gpg)](https://goreportcard.com/report/github.com/invidian/terraform-provider-gpg)

The GPG provider allows to generate GPG encrypted message in ASCII-armored format using Terraform.

This provider uses built-in [openpgp](https://godoc.org/golang.org/x/crypto/openpgp) Golang library to perform GPG encryption. Currently the only supported option is encrypting message with public keys.

Managing GPG keyring or signing files is currently not implemented.

## Table of contents
* [User documentation](#user-documentation)
* [Building and testing](#building-and-testing)
* [Authors](#authors)

## User documentation

For user documentation, see [Terraform Registry](https://registry.terraform.io/providers/invidian/gpg/latest/docs).

## Building

For testing builds, simply run `docker build .`, which will download all dependencies, run build, test and linter.

For local builds, run `make` which will build the binary, run unit tests and linter.

## Releasing

This project use `goreleaser` for releasing. To release new version, follow the following steps:

* Add a changelog for new release to CHANGELOG.md file.

* Tag new release on desired git, using example command:

  ```sh
  git tag -a v0.4.7 -s -m "Release v0.4.7"
  ```

* Push the tag to GitHub
  ```sh
  git push origin v0.4.7
  ```

* Run `goreleser` to create a GitHub Release:
  ```sh
  GITHUB_TOKEN=githubtoken GPG_FINGERPRINT=gpgfingerprint goreleaser release --release-notes <(go run github.com/rcmachado/changelog show 0.4.7)
  goreleaser
  ```

* Go to newly create [GitHub release](https://github.com/invidian/terraform-provider-gpg/releases/tag/v0.4.7), verify that the changelog
  and artefacts looks correct and publish it.

## Authors

* **Mateusz Gozdek** - *Initial work* - [invidian](https://github.com/invidian)
