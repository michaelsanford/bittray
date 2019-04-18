---
layout: page
title: Installation & Configuration
permalink: /installing/
order: 1
---

# Downloads

[![v1.1.1 Reference Build](https://img.shields.io/static/v1.svg?label=v1.1.1&message=Reference%20Build&color=green?style=flat&logo=appveyor)](https://ci.appveyor.com/project/michaelsanford/bittray/builds/23440146)
![GitHub Release Date](https://img.shields.io/github/release-date/michaelsanford/bittray.svg)
![GitHub All Releases](https://img.shields.io/github/downloads/michaelsanford/bittray/total.svg)

|Version|Milestone|Link|Checksum (SHA1)|
|---|---|---|---|
|1.1.1|["Glass Harp"](https://github.com/michaelsanford/bittray/milestone/3?closed=1)|[:floppy_disk: Zip](https://github.com/michaelsanford/bittray/releases/download/1.1.1/bittray-1.1.1.zip)|`bc03e593bb631573fb6abd47a323e4d14f939c6c`|

# Installation
1. Download the latest stable Release above. ([All Releases are available on Github](https://github.com/michaelsanford/bittray/releases).)
1. Verify the checksum with `certUtil -hashfile .\bittray-1.1.1.zip sha1`
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

## Configuring Polling Interval

By default, Bittray polls every 15 seconds.

If this is too frequent, set an environment variable called `BITTRAY:PINT`
to the number of seconds to poll at.

I'm working on a more elegant way to configure this.


## Older Releases

|Version|Link|Checksum (SHA1)|
|1.1.0|[Zip](https://github.com/michaelsanford/bittray/releases/download/1.1.0/bittray-1.1.0.zip)|`cd283afd10f613919bb4dc694bce3c1f9bd23483`|
|1.0.1|[Zip](https://github.com/michaelsanford/bittray/releases/download/v1.0.1/bittray-1.0.1.zip)|`03c103369cc3c2bbd4386048eea9b8e3d936dac5`|
|1.0.0|[Zip](https://github.com/michaelsanford/bittray/releases/download/v1.0.0/bittray-1.0.0.zip)|`7c5f45eb57bc98b81a3581a3e99486c20c92f8fa`|

