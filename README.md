# qktemplates

This repository holds some [Quokka](https://github.com/Depado/quokka) templates

## Usage

First of all install [Quokka](https://github.com/Depado/quokka) using the
release page or install it from source. 

Then you can either use Quokka directly with the git repo, or clone this repo
and use it as a local (filesystem) provider.

```sh
# Use directly from the distant repo
$ qk git@github.com:Depado/qktemplates.git -p drone myproject
# Alternatively
$ git clone git@github.com:Depado/qktemplates.git
$ qk qktemplates/drone myproject
```

## Templates

### drone

This template will generate a `.drone.yml` file that should fit the needs of a
go project.

### license

Simple template that renders a single `LICENSE` file in the output directory and
lets you select which license you want. This template is also embedded in some
other templates

## simplego

This template will generate a Go project:
- Base Go code (with or without CLI base) that uses go modules for its 
  dependency management with build and version variable injection
- A Makefile
- An optional Dockerfile and .dockerignore (multi-step with alpine as the final 
  image), tagging with the commit's SHA1 and `latest`
- An optional `.goreleaser.yml` file and the associated Makefile target to use
  [goreleaser](https://goreleaser.com/) 
- An optional `.drone.yml` for [Drone 1.0](https://docs.drone.io/)
- An optional `LICENSE` file
- A bare `README.md` file
- A basic configuration for [golangci-lint](https://github.com/golangci/golangci-lint)