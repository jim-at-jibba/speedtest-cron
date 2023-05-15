<h1 align="center">Welcome to speedtest-cron 👋</h1>
<p>
  <a href="https://github.com/jim-at-jibba/speedtest-cron/blob/main/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/jimgbest" target="_blank">
    <img alt="Twitter: jimgbest" src="https://img.shields.io/twitter/follow/jimgbest.svg?style=social" />
  </a>
</p>

> Needed to check my internet regularly for flake. Did this.

## Requirements

- [Speedtest CLI](https://www.speedtest.net/apps/cli)
- [MacOS](https://github.com/julienXX/terminal-notifier)

## Usage

```sh
# I have this setup to run on a cronjob. Do with it as you please.
$ git checkout
$ go build .
$ ./speedtest-cron


# I have this for my cron entry
SHELL=/bin/sh
PATH=/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin:/opt/homebrew/bin:/opt/homebrew/sbin

# m h dom mon dow user   command
0 */2 * * * cd ~/code/other/speedtest-cron && ./speedtest-cron >> ~/code/other/speedtest-cron/speedtest.txt 2>&1

```

## Author

👤 **James Best**

- Website: jamesbest.uk
- Twitter: [@jimgbest](https://twitter.com/jimgbest)
- Github: [@jim-at-jibba](https://github.com/jim-at-jibba)

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
