---
layout: page
title: Security
permalink: /security/
order: 3
---

# Your Bitbucket Password

You will be asked to provide the password you use to log in to Bitbucket _every time you start_ `Bittray`.

**It will not be stored.**

# Security

The WCM [is not a secure password store](https://github.com/michaelsanford/bittray/issues/14) by default.

In order to securely store the password in the WCM (or registry, or anywhere) in a retrievable form, it would need to be encrypted using
a passphrase provided by the user.

The user would need to enter that decryption passphrase every time `Bittray` launched to decrypt the stored Bitbucket
credentials.

So you'd have to enter a password every time you launch anyway, saving no effort.

Storing it is arguably actively harmful.

# Conclusion

[So I just don't  store it](https://github.com/michaelsanford/bittray/blob/master/config/config.go).
