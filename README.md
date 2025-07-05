# learn GoLang

[//]: # ([![Go Report Card]&#40;https://goreportcard.com/badge/github.com/learn-go-lang/learn-go-lang&#41;]&#40;https://goreportcard.com/report/github.com/learn-go-lang/learn-go-lang&#41;)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

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
go build -o firstapp firstapp/cmd/main.go
```

</details>

<br>
<details>
<summary><b>printing</b></summary>

To print out a message in Go, There are several ways to do it, 
but the most common way is to use the `fmt` package.

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

This will print "Hello, World!" to the console.

### Type printing
To print the type of variable,

`%v` will print the value in a default format, 
`%T` will print the type of the variable.

```go
func main() {
	x:= 42
    fmt.Printf("Type of x: %T\n", x)
}
```

</details>

<br>
<details>
<summary><b>Variable declaration</b></summary>

## Variable conventions
In Go, variable names should be descriptive and follow the following conventions:
- Use camelCase for variable names (e.g., `myVariable`).
- Avoid using underscores in variable names (e.g., use `myVariable` instead of `my_variable`).
- Use meaningful names that reflect the purpose of the variable.
- Avoid using single-letter variable names, except for loop variables (e.g., `i`, `j`, `k`).
- Use short names for variables that are used in a small scope, such as loop variables.
- Avoid using reserved keywords as variable names (e.g., `if`, `for`, `func`, etc.).
- Use uppercase letters for package-level variables that need to be exported (accessible outside the package).
- Use lowercase letters for package-level variables that are not exported (accessible only within the package).
- Use `const` for constant values that do not change throughout the program.

In Go, you can declare variables using the `var` keyword or the short variable declaration syntax `:=`.
<br>
Variables can be declared inside a function or at the package level.
<br>
The differance in a function scope is only accessible within that function, 
while package-level variables are accessible throughout the package.

When using the package level variable, you can't use the `:=`, It's available only within the function.

Here is some examples of variable declaration in Go:

```go
package main

import "fmt"

// Single variable declaration
var y string = "Hello, Go!"

// Package-level variable declaration
var (
	a int = 10
    b float64 = 3.14
    c bool = true
)
// Another variables that can be declared
var (
	actorName string = "Jhon Smith"
	companion string= "Jane Doe"
	docrotNumber int = 3
	season int = 11
) 
func main() {
	// Using short variable declaration
	x := 42
	fmt.Println("Value of x:", x)
	fmt.Printf("Value of y: %v, %T", y, y)
	fmt.Printf("Value of a: %v, %T", a, a)
	fmt.Printf("Value of b: %v, %T", b, b)
	fmt.Printf("Value of c: %v, %T", c, c)
	fmt.Printf("Value of actorName: %v, %T", actorName, actorName)
	fmt.Printf("Value of companion: %v, %T", companion, companion)
	fmt.Printf("Value of docrotNumber: %v, %T", docrotNumber, docrotNumber)
	fmt.Printf("Value of season: %v, %T", season, season)
}
```





</details>
