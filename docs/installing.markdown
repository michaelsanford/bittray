---
layout: page
title: Installation & Configuration
permalink: /installing/
order: 1
---

# Installation

1. Download the [latest Release executable](https://github.com/michaelsanford/bittray/releases).
1. Run it.

# Configuration

On first launch, `Bittray` will ask for two things:

1. The URL of your Bitbucket Enterprise server, and
1. The username you use to log in to Bitbucket

These two settings are stored in the [Windows Credential Manager](https://support.microsoft.com/en-ca/help/4026814/windows-accessing-credential-manager)
 as a _Windows Credential_ under the key `bittray:conf`.

### Note
The  URL must take _exactly_ the format `http://example.com:7990`, with a port (if not `80`) and no trailing slash.

## Next step

:arrow_forward: [Running It and understanding the UI]({{ site.baseurl }}/using).