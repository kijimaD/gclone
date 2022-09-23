[![⚗️Check](https://github.com/kijimaD/gclone/actions/workflows/check.yml/badge.svg)](https://github.com/kijimaD/gclone/actions/workflows/check.yml)

# gclone

gclone is literate `git clone` cli tool.

<img src="https://user-images.githubusercontent.com/11595790/192002784-3a72243d-2343-42d2-a8e5-581977faa382.jpg" width="40%" align=right>

# Install

```sh
$ go install github.com/kijimaD/gclone@main
```

# How to use

make `config.yml`
(↓example)
```yaml
jobs:
  - dest: '~/Project/test0' # specify exist directory
    repos:
      - git@github.com:kijimaD/my_go.git
      - git@github.com:kijimaD/gin_hello.git
  - dest: '~/Project/test1' # specify exist directory
    repos:
      - git@github.com:fatih/color.git
      - git@github.com:joho/godotenv.git
```

and run!(on directory config.yml existing)
```shell
$ gclone install
```

# Options

-f: config file path
```shell
$ gclone -f dir/config.yml
```
