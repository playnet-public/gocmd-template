# gocmd-template
Our template project for building command-line tools with golang

## Dependencies

This project has a pretty complex Makefile and therefore requires `make`.

Go Version: 1.8

Install all further requirements by running `make deps`

## Usage

This project is meant to offer a basic template for developing PlayNet command-line tools.
The Makefile is configurable to some extent by providing variables at the top.
Any further changes should be thought of carefully as they might brake CI/CD compatibility.

One project might contain multiple tools whose main packages reside under `bin`. Other packages go directly into the root directory.
Single projects can be handled by calling `make toolname maketarget` like for example:
```
make template dev
```
All tools at once can be handled by calling `make full maketarget` like for example:
```
make full build
```
Build output is being sent to `./build/`.

If you only package one tool this might seam slightly redundant but this is meant to provide consistence over all projects.
To simplify this, you can simply call `make maketarget` when only one tool is located beneath `bin`. If there are more than one, this won't do anything (including not return 1) so be careful.

## Dependency Management

This project uses https://github.com/golang/dep to manage dependencies.
For updating you local vendor directory, simply run `make deps` or `dep ensure`

For updating the manifest and adding new dependencies run `dep ensure -update` or refer to the
[GoDep FAQ](https://github.com/golang/dep/blob/master/docs/FAQ.md).

## Output and Logging

Command-line tools mostly have too two main output targets. One is direct user communication, the other is logging internal mechanics and details.

For user output this should always happen via stdout and therefor via the `fmt` package.
Logging on the other hand is something quite different and might be handled in multiple ways.

While error logging at PlayNet is done via [Sentry](https://sentry.io) we got two logging packages currently in use.
One is [glog](https://github.com/golang/glog) used for fine grained debugging output.

Glog offers us a fine grained level system for debugging meaning we can go quite deep in detail by simply setting -v. This is not meant for production logging as it only contains text and might be too verbose. The default log level for glog should be 0 or 1.
To set the log level run the tool like this:
```
./build/template -logtostderr -v=3
```

The other logging package is [zap](https://github.com/uber-go/zap) which offers us flexible, production-grade and strctured logging.
This template contains an example for configuring and using zap. For logging important (non-debugging) information inside PlayNet Applications always use zap.

For more information on how to use it, refer to the [docs](https://godoc.org/go.uber.org/zap).

## Versioning

PlayNet projects are managed using [semantic versions](https://semver.org).
For creating a new version, tag your current commit and push it. It should then be handled by CI/CD and be built.

Same applies for Docker Images and Debian Packages.
If using different repos for building those images/packages, please refer to subtag-versions (toolversion-buildscriptversion).

## Docker

The project contains two Dockerfiles for building and running the provided tools.
It is only possible to build and package all tools (if multiple) for now.

To build the docker image run `make docker` and to upload the image `make upload`.
The Dockerfile containing the built binaries is from scratch and very small. If debugging is required better do this locally.

## Standards

This Project represents internal code standards related to Go development at PlayNet. We are permanentely improving these standards, but this state is already being used in production.

PlayNet tools and apps should always match highest Golang standards. Testing is not optionally!
Please refer to our other Guideline Documents in the [Wiki](https://wiki.play-net.org).

Suggestions are always welcome and questions will be happily answered.
If you would like to