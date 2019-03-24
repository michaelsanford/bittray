# Bittray

![GitHub release](https://img.shields.io/github/release/michaelsanford/bittray.svg)
![GitHub Release Date](https://img.shields.io/github/release-date/michaelsanford/bittray.svg)
![GitHub All Releases](https://img.shields.io/github/downloads/michaelsanford/bittray/total.svg)

[![Travis CI](https://travis-ci.org/michaelsanford/bittray.svg?branch=master)](https://travis-ci.org/michaelsanford/bittray)
[![AppVeyor Status](https://ci.appveyor.com/api/projects/status/github/michaelsanford/bittray)](https://ci.appveyor.com/project/michaelsanford/bittray)
[![Go Report Card](https://goreportcard.com/badge/github.com/michaelsanford/bittray?branch=master)](https://goreportcard.com/report/github.com/michaelsanford/bittray)

A system tray application for monitoring your BitBucket pull requests, written in Go.

# API Support

This version supports [Bitbucket REST API v1.0](https://docs.atlassian.com/bitbucket-server/rest/4.10.1/bitbucket-rest.html).

It's tested to work up to Atlassian Bitbucket v5.16.3.

I do plan to update it to API 2, once I have a need. Non-breaking PRs to that end are welcome (but please file an issue).

# Details and Use

See https://michaelsanford.github.io/bittray/ for details.

# Building with Windows Resource Extensions

You'll need Electron's [`rcedit-x64`](https://github.com/electron/rcedit/releases) in your PATH before continuing.

By default, Windows 10's PowerShell Execution Policy is set to `Restricted`. To build, you'll need to temporarily set that to `Unrestricted` for the packager to run.

You _do not_ need to run it with Administrative privileges.

```powershell
powershell.exe -ExecutionPolicy Unrestricted .\build.ps1 -version 0.0.0
```

Make sure `config.CurrentVersionTag` (in `constants.go`) matches what you provide in `-version`.

Read up on [PowerShell Execution Policy](https://docs.microsoft.com/en-ca/powershell/module/microsoft.powershell.core/about/about_execution_policies).

# License

[MIT](https://github.com/michaelsanford/bittray/blob/master/LICENSE)

# Disclaimer

This project is in no way affiliated with or officially endorsed by Atlassian Inc.

Icons are courtesy [Recep Kütük](https://www.iconfinder.com/iconsets/bitsies).
