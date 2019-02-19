---
layout: page
title: Installation & Configuration
permalink: /installing/
order: 1
---

# Downloads

|Version|Code Name|Link|Checksum (SHA1)|
|---|---|---|---|
|1.0.0|"Flying Snow"|[Zip](https://github.com/michaelsanford/bittray/releases/download/v1.0.0/bittray-1.0.0.zip)|`7c5f45eb57bc98b81a3581a3e99486c20c92f8fa`|

# Installation
1. Download the latest stable Release above. ([All Releases are available on Github](https://github.com/michaelsanford/bittray/releases).)
1. Unpack
1. Run `bittray.exe`

# Configuration

On first launch, `Bittray` will ask for two things:

1. The _username_ you use to log in to Bitbucket, and
1. The _URL_ of your Bitbucket server

    :information_source: The  URL must take _exactly_ the format `http://example.com:7990`, with a port (if not `80`) and no trailing slash.

    These two settings are stored in the [Windows Credential Manager](https://support.microsoft.com/en-ca/help/4026814/windows-accessing-credential-manager)
    as a _Windows Credential_ under the key `bittray:conf`.

 1. You will then be asked for your Bitbucket _password_.
 
     It will be not stored and is only valid until you quit Bittray. See [Security]({{ site.baseurl }}/security) for more details.

## Binaries are unsigned

The artifacts produced by the current build process are not signed. You will be presented with a UCM security warning
when you run the application.

[I'm working on signing them (#19).](https://github.com/michaelsanford/bittray/issues/19)

Run it if you trust me, I guess 🤷‍♂️

## Next step

:ok: [Running It and understanding the UI]({{ site.baseurl }}/using).