# learngo
This repo created to store learning path of GoLang


<details>
<summary><b>Go configuration for the OS</b></summary>

Make sure to insert this in your `.bashrc` or `.zshrc` file.
## GO Config
export GOPATH=$HOME/golib
export PATH=$PATH:$GOPATH/bin
export GOBIN=$GOPATH/bin
export GOPATH=$GOPATH:$HOME/code

Create the directory if it does not exist
```bash
mkdir -p ~/golib
mkdir -p ~/code/{bin,src,pkg}
```
</details>


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



```
</details>
