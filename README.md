[![⚗️Check](https://github.com/kijimaD/gclone/actions/workflows/check.yml/badge.svg)](https://github.com/kijimaD/gclone/actions/workflows/check.yml)

# gclone

<img src="https://user-images.githubusercontent.com/11595790/192002784-3a72243d-2343-42d2-a8e5-581977faa382.jpg" width="40%" align=right>

gclone is literate `git clone` cli tool.

```
$ gclone
────────────────────────
Save dir: /root/Project1
Repo count: 2
────────────────────────
⢿✔ git@github.com:fatih/color.git
⡿✔ git@github.com:joho/godotenv.git

────────────────────────
Save dir: /root/Project2
Repo count: 2
────────────────────────
⡿✔ git@github.com:kijimaD/my_go.git
⣻✔ git@github.com:kijimaD/gin_hello.git

done!
Success: 4
Fail: 0
Process: 8760ms
```

# Install

```sh
$ go install github.com/kijimaD/gclone@main
```

# How to use

make `config.yml`
(↓example)
```yaml
groups:
  - dest: '~/Project/test0'
    repos:
      - git@github.com:kijimaD/my_go.git
      - git@github.com:kijimaD/gin_hello.git
  - dest: '~/Project/test1'
    repos:
      - git@github.com:fatih/color.git
      - git@github.com:joho/godotenv.git
```

and run!(on directory config.yml existing)
```shell
$ gclone
```

# Options

-f: config file path
```shell
$ gclone -f dir/config.yml
```

# Docker

This command is for testing, not save result your disk. If you want to save disk, mount save directory.
```shell
docker run --rm
           -it
           -v "${PWD}":/workdir \
           -v "${HOME}/.ssh":/root/.ssh \
           ghcr.io/kijimad/gclone:latest
```
