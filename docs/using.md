---
layout: page
title: Using It
permalink: /using/
order: 2
---

# What Counts as "Pending"

A Pull Request will appear in the Pending queue and provoke the alarm icon if

1. You are involved as a Reviewer, and
1. You have not clicked `Approved` or `Needs Work`

Clicking `Needs Work` says to the developer that the PR requires attention and will remove it from your queue because you've "done something about it".

When the developer pushes a new changeset, the `Needs Work` flag is automatically cleared and the request will fall back into Pending.

# Icon States

|Icon|Meaning|Tooltip Text|
|---|---|---|
|![Locked](/assets/lock.png)|The API is locked, you need to provide your Bitbucket password _or_ Bittray cannot connect (VPN down).|_Locked..._|
|![PRs waiting](/assets/alarm.png)|You have Pull Requests that need attention|_X PR(s) waiting..._|
|![Queue clear](/assets/checkmark.png)|All of your Pull Requests have been actioned|_PR Queue clear!_|

# Menus

Clicking the tray shows two menu options:

### Update Available

When a new release is pushed to Github, a menu item will appear to inform you. Click it to open the download page.

### Quit

Quits Bittray.

### Review X Pull Requests / Go to Bitbucket

Opens the Bitbucket pull request Dashboard you provided.

# Next Step

A note on [Security]({{ site.baseurl }}/security).