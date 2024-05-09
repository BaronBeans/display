## Display

Simple command line app written in Golang which returns the IPv4 address of the WSL virtual network adapter.

This allows you to add the following to ~/.bashrc (as long as you have run `make install` in this directory)

```bashrc
export DISPLAY=$(display):0
```
