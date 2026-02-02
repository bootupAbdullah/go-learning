### <u>Chapter 1</u>

#### Third-Party Go Tools

\- Go code is shared via their source code through repos - go tools can be installed using go install (location of repo (HTTP address)) followed by an @ and the version of the tool you want. 

\- the contents of Go repos are cached in proxy servers.

\- depending on the repo and the values in your GOPROXY env variable, go install may download from proxy or directly from a repo

### ______

```bash
Go tool: Hey
go install github.com/rakyll/hey@latest
```

> GOPROXY more than likely being the a designation (variable) you can have in your .env file to choose/set a particular HTTP address for a proxy sever


#### Create & Run Executables

\- use the 'go build' command to create an executable

> go build hello.go

\- to run an executable:

> ./hello.go

\- use the -o flag to change name or location of binary file

> go build -o hello_world hello.go

### Notes

- Go tools can be distributed as pre-compiled binaries - can also be built from source 

- no "npm" like registry for Go

- go install takes the location of the source code repo as an argument followed by an @ and the version of the tool you want to install


``` bash 
Formatting options:

(!) gofmt - automatically reformats your code to match the standard format

# See what would change (doesn't modify):
gofmt main.go

# Actually format file:
gofmt -w main.go

# Format all .go files recursively:
gofmt -w .

# Check which files need formatting:
gofmt -l .

# Show differences before applying:
gofmt -d main.go

(!) golint - this tries to ensure your code follows style guidelines

#This runs golint over your entire project.
golint ./...


(!) go vet - this go tool detects mistakes such as passing the wrong number of parameters to formatting methods or assigning values to variables that are never used. 

#run with: 
go vet ./...

```




