---
layout: page
title: Installation & Configuration
permalink: /installing/
order: 1
---

# Downloads

[![v1.1.3 Reference Build](https://img.shields.io/static/v1.svg?label=v1.1.3&message=Reference%20Build&color=green?style=flat&logo=appveyor)](https://ci.appveyor.com/project/michaelsanford/bittray/builds/23991990)
![GitHub Release Date](https://img.shields.io/github/release-date/michaelsanford/bittray.svg)
![GitHub All Releases](https://img.shields.io/github/downloads/michaelsanford/bittray/total.svg)

|Version|Milestone|Link|Checksum (SHA1)|
|---|---|---|---|
|[1.1.3](https://github.com/michaelsanford/bittray/tree/1.1.3)|[Bone Mill](https://github.com/michaelsanford/bittray/milestone/4?closed=1)|[:floppy_disk: Zip](https://github.com/michaelsanford/bittray/releases/download/1.1.3/bittray-1.1.3.zip)|`c9dd0660abff03fd5d5df7f96c3f5685ea23af51`|

# Changelog

- 1.1.3 was simply compiled with Go 1.13.6 (up from Go 1.12 in release 1.1.2).
- `-poll` flag was removed.

# Installation
1. Download.
1. Verify the checksum with `certUtil -hashfile .\bittray-1.1.3.zip sha1`
1. Unpack
1. Run `bittray.exe`

#### Binaries are unsigned

The artifacts produced by the current build process are not signed. You will probably be presented with a UAC security warning
when you run the application. [It is not feasible to sign them (#19).](https://github.com/michaelsanford/bittray/issues/19)

In Windows 10, click "More" and <kbd>Run Anyway</kbd> (if you trust me, I guess :man_shrugging: ).

# Configuration

On first launch, `Bittray` will ask for two things:

1. The _username_ you use to log in to Bitbucket, and
1. The _URL_ of your Bitbucket server

    :information_source: The  URL must take _exactly_ the format `http://example.com:7990`, with a port (if not `80`) and no trailing slash.

    These two settings are stored in the [Windows Credential Manager](https://support.microsoft.com/en-ca/help/4026814/windows-accessing-credential-manager)
    as a _Windows Credential_ under the key `bittray:conf`.

 1. You will then be asked for your Bitbucket _password_.
 
     It will be not stored and is only valid until you quit Bittray. See [Security]({{ site.baseurl }}/security) for more details.

#### Start on Login

If you'd like `bittray` to start automatically on login:

1. Open a Run dialog with <kbd>Windows Key</kbd> + <kbd>R</kbd>
1. Run the command `shell:startup` to open your autostart folder
1. Create a shortcut to `bittray.exe` by clicking the exe, <kbd>CTRL</kbd> + <kbd>SHIFT</kbd> + dragging it into the folder you just opened.

## Next step

[Running It and understanding the UI]({{ site.baseurl }}/using).
