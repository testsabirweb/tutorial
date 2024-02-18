
## Hello World program
```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

```bash
$ go run hello-world.go
hello world

#Sometimes we’ll want to build our programs into binaries. We can do this using go build.
$ go build hello-world.go
$ ls
hello-world    hello-world.go

#We can then execute the built binary directly.
$ ./hello-world
hello world
```
### What is package?
In Golang, a package is a way to organize and reuse code. A package consists of a collection of Go source files that reside in the same directory. Each source file within a package typically declares a package clause at the beginning, indicating the package to which it belongs.

### Module VS package
- A module in Go refers to the versioned collection of related packages and their dependencies.
- A module is represented by the go.mod file, which includes module-related information, dependencies, and version requirements.
- Modules were introduced in `Go 1.11` to address the need for versioning and dependency management in Go projects.
- Modules allow you to define the set of dependencies for your project, and they facilitate versioning and reproducibility of builds.
```
module example.com/myproject

go 1.22

require (
    github.com/somepackage v1.2.3
    // other dependencies...
)
```
### Some important Go commands
1. `go mod init`: This command is used to initialize a new module or convert an existing Go project into a module.
    It creates a go.mod file in the root directory of your project, which contains module-related information, dependencies, and version requirements.
    Example:

    ```bash
    go mod init <module-name>
    ```

2. `go tidy`: The go tidy command is used to clean up and remove any dependencies that are no longer required by the project.
It updates the `go.mod` and `go.sum` files to reflect the current set of dependencies used by the project.
Example:
```bash
go mod tidy
```

3. `go fmt`: The go fmt command is used to format Go source code files according to the standard Go formatting conventions.
It automatically updates the code to adhere to the style guidelines specified in the official Go formatting rules.
Example:

```bash
go fmt ./
```

### go.mod VS go.sum
- go.mod file: Configuration file for a Go module. Specifies the module name, Go version, and dependencies.
- go.sum file: Records checksums for specific versions of module dependencies. Ensures the integrity and reproducibility of builds.