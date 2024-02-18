# Basics

Go is an open source programming language designed for building simple, fast, and reliable software

# Install Golang 1.22.0 using Bash

1. Open a terminal.

2. Download the Golang archive for version 1.22.0 using `wget`:

   ```bash
   wget https://golang.org/dl/go1.22.0.linux-amd64.tar.gz

   tar -xvf go1.22.0.linux-amd64.tar.gz

   sudo mv go /usr/local
    ```
3. Set the environment variables for Golang by adding the following lines to your `~/.bashrc` or `~/.bash_profile `file:
    ```bash
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    ```
4. Verify the installation by checking the Golang version: `go version`