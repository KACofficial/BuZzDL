# BuZzDL
![BuZzDL Logo](images/BuZzDL-Banner.png)
BuZzDL is a feature rich downloader for [Bsocial.Buzz](https://bsocial.buzz) Videos.  
It supports downloading in almost every format using ffmpeg.  

## Install/Update
This has only been tested on linux.
- clone this repo `git clone https://github.com/KACofficial/BuZzDL.git`
- change directory `cd BuZzDL`
- run `./scripts/install.sh`
- then you should be able to use `buzzdl` in your terminal.

## Usage
There are 3 options:
- `-u`/`-url` to specify the url of the video(if unused it will prompt in the program)
- `-f`/`-format` to specify the video format(defaults to mp4)
- `-o`/`-output` to specify the output location(defaults to current working directory)  
### Examples
```bash
# This will download to dump-chatgpt...olEgE6nzqOgkAA6.mp4
❯ buzzdl -u https://bsocial.buzz/shorts/dump-chatgpt-use-fabric-instead_olEgE6nzqOgkAA6.html

# This will download to dump-chatgpt...olEgE6nzqOgkAA6.mkv
❯ buzzdl -u https://bsocial.buzz/shorts/dump-chatgpt-use-fabric-instead_olEgE6nzqOgkAA6.html -f mkv

# This will download to videos/dump-chatgpt...olEgE6nzqOgkAA6.mkv
❯ buzzdl -u https://bsocial.buzz/shorts/dump-chatgpt-use-fabric-instead_olEgE6nzqOgkAA6.html -f mkv -o videos
```

## Showcase Video.
(Click the image)
[![BuZzDL Logo](images/BuZzDL-Banner.png)](https://bsocial.buzz/watch/buzzdl-showcase_b2F2EvbKoMCcieQ.html)

## Uninstall
- run `./scripts/uninstall.sh`
- check if `buzzdl` is removed from your terminal.
