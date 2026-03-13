# Turn on the darkmode in CLI


## Original Source Code
- https://github.com/mafredri/macos-darkmode/blob/master/cmd/darkmode/main.go
- https://github.com/pndurette/zsh-lux/blob/main/zsh-lux.plugin.zsh#L125C1-L137C2


## Usages
```bash
# Get current mode (off or on)
darkmode

# Turn on darkmode
darkmode on

# Turn off darkmode
darkmode off

# Toggle the mode
darkmode toggle
```

```bash
source ./macos_is_dark.sh

# Get current mode (true or false)
is_dark
```

## Build and Installation
```bash
# You will need to install Golang to compile it yourself
go build -o /usr/local/bin
```


## Download pre-compiled binary

```bash
wget https://github.com/yejun614/macos-darkmode-go/releases/download/v0.1.0/darkmode -O /usr/local/bin
```
