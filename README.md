# learngo
This repo created to store learning path of GoLang

<details>
<summary><b>Install Go</b></summary>

[Download Go](https://go.dev/dl/)

Choose the latest version and download the binary for your OS.
You can also choose the installer for your OS, such as `go1.20.3.darwin-amd64.pkg` for macOS.
</details>
<br>

<details>
<summary><b>Tools and cheat sheet </b></summary>

goenv aims to be as simple as possible and follow the already established successful version management model of pyenv and rbenv.
[Go-env](https://github.com/go-nv/goenv)

</details>

<br>

<details>
<summary><b>Go configuration for the OS</b></summary>

Make sure to insert this in your `.bashrc` or `.zshrc` file.
## GO Config
export GOPATH=$HOME/golib
export PATH=$PATH:$GOPATH/bin
export GOBIN=$GOPATH/bin
export GOPATH=$GOPATH:$HOME/code
export GO111MODULE=on

Create the directory if it does not exist
```bash
mkdir -p ~/golib
mkdir -p ~/code/{bin,src,pkg}
```
</details>
<br>

<details>

<summary><b>Work with GitHub</b></summary>

Make sure to insert this in your `.bashrc` or `.zshrc` file.
## src Config
In order to work with GitHub, you need to set the `GOPATH` to your `src` directory.

Create the directory if it does not exist
```bash
mkdir -p ~/code/src/github.com/<GitHubUserName>/firstapp
touch ~/code/src/github.com/<GitHubUserName>/firstapp/main.go
```

</details>

<br>

<details>
<summary><b>Compile, build, run</b></summary>

To build and run your Go application, you can use the following commands:

```bash
go build -o firstapp firstapp/main.go
```

</details>

