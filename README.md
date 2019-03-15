# Bittray

[![Travis CI](https://travis-ci.org/michaelsanford/bittray.svg?branch=master)](https://travis-ci.org/michaelsanford/bittray)
[![AppVeyor Status](https://ci.appveyor.com/api/projects/status/github/michaelsanford/bittray)](https://ci.appveyor.com/project/michaelsanford/bittray)
[![Go Report Card](https://goreportcard.com/badge/github.com/michaelsanford/bittray?branch=master)](https://goreportcard.com/report/github.com/michaelsanford/bittray)

A system tray application for monitoring your BitBucket pull requests, written in Go.

# API Support

This version supports [Bitbucket REST API v1.0](https://docs.atlassian.com/bitbucket-server/rest/4.10.1/bitbucket-rest.html).

I do plan to update it to API 2, once I have a need. Non-breaking PRs to that end are welcome (but please file an issue).

# Details and Use

See https://michaelsanford.github.io/bittray/ for details.

# Building with Windows Resource Extensions

You'll need Electron's [rcedit](https://github.com/electron/rcedit/releases) before continuing.

By default, Windows 10's PowerShell Execution Policy is set to `Restricted`. To build, you'll need to temporarily set that to `Unrestricted` for the packager to run.

You _do not_ need to run it with Administrative privileges.

```powershell
powershell.exe -ExecutionPolicy Unrestricted .\build.ps1 -version 0.0.0
```

If you don't have `rcedit-x64.exe` in your PATH, you must specify a location to it with the `-rcedit $path` flag.

```powershell
powershell.exe -ExecutionPolicy Unrestricted .\build.ps1 -version 0.0.0 -rcedit "E:\somepath\rcedit-x64.exe"
```

Make sure `config.CurrentVersionTag` (in `constants.go`) matches what you provide in `-version`.

Read up on [PowerShell Execution Policy](https://docs.microsoft.com/en-ca/powershell/module/microsoft.powershell.core/about/about_execution_policies).

# License

MIT