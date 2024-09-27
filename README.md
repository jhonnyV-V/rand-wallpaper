# What is this?
just a small cli tool to change your wallpaper image to a random one from a set of your choice

this creates a folder in your config directory and expects to read the images from that directory

# Dependency
this project only depends on hsetroot

# How to use?
create a directory named rand-wallpaper under your config directory or just the tool and it will create it,
then copy or move the images you want to use to that folder, in most linux distros it will be ~/.config/rand-wallpaper/

# How to install
## With Go
```bash
go install github.com/jhonnyV-V/rand-wallpaper@latest
```

## Compile it yourself

```bash
git clone git@github.com:jhonnyV-V/rand-wallpaper && cd rand-wallpaper
```
```bash
go build -o rand-wallpaper
```
```bash
sudo mv ./rand-wallpaper /usr/local/bin
```

# Use With i3wm
you can create a binding 
```
bindsym $mod+Shift+w rand-wallpaper
```
or use it on start
```
exec --no-startup-id rand-wallpaper
```
