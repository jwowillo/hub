# `hub`

`hub` is a configurable website directory.

## Install

`go get github.com/jwowillo/hub/cmd/hub`

## Run `hub`

`hub`

`hub ?--port 80 ?--cache-duration 24`

Make sure a file named 'config.yaml' is in the working directory that matches
the format below. This file can be udpated while `hub` is running. The optional
flag `--port` determines what port to listen on. The optional flag
`--cache-duration` determines how frequently caches are cleared.

## Run `copy_config`

`copy_config 127.0.0.1 user '~'`

Make sure the config file is in the working directory. The script copies the
file to the passed remote host, user, and directory.

## Run `copy_assets`

`copy_assets 127.0.0.1 user '~'`

The script copies the directories containing assets from the `hub` repo to the
passed remote host, user, and directory.

## Run `deploy`

`deploy 127.0.0.1 user '~'`

Make sure the config file is in the remote directory. The script deploys `hub`
to the passed remote host, user, and directory.

## Example Config

```
- URL: https://github.com
  name: github
- URL: https://reddit.com/r/programming
  name: programming reddit
  ...
```
