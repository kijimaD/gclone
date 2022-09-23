Literate git clone cli tool.

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

and run!
```shell
$ gclone
```

# Options

-f: config file path
```shell
$ gclone -f dir/config.yml
```
