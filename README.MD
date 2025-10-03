# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

### Runtime requirements:
Golang 1.24+

### Installing Golang:
On Arch-based distributions
```sh
sudo pacman -S go
```

On Debian-based distributions
```sh
sudo apt update && sudo apt install go
```

On Redhat-based distributions (>= RHEL 9)
```sh
dnf install golang
```

### Getting Started
Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```
