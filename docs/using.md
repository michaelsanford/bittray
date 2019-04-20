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

When the developer pushes a new changeset, the `Needs Work` flag is automatically cleared and the request will fall back into `Pending`.

# Icon States

|Icon|Meaning|Tooltip Text|
|---|---|---|
|![Locked](/assets/lock.png)|The API is locked, you need to provide your Bitbucket password _or_ Bittray cannot connect (VPN down).|_Locked..._|
|![PRs waiting](/assets/alarm.png)|You have Pull Requests that need attention|_X PR(s) waiting..._|
|![Queue clear](/assets/checkmark.png)|All of your Pull Requests have been actioned|_PR Queue clear!_|
|![Rate limited](/assets/rate.png)|You exceeded the API rate limit; automatic request backoff is in effect.|_Rate Limited!_|

Are you often Rate limited? Increase the polling interval in seconds with the commandline flag `-poll=xx`. The default is 15 seconds.

# Menus

Clicking the tray shows some menu options:

|Menu Name|Function|
|---|---|
|Update Available|When a new release is pushed to Github, a menu item will appear to inform you. Click it to go to the releases page.|
|Review _X_ Pull Requests|Opens the Bitbucket pull request Dashboard you provided.|
|Go to Bitbucket|Opens Bitbucket Dashboard, but when you don't have any PRs|
|Quit|Quits Bittray|

# Next Step

A note on [Security]({{ site.baseurl }}/security).