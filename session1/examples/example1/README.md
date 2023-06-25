# Example 1 - Build a CLI using the Go standard library “hello world!

This is a simple command-line interface (CLI) written in Go that outputs "Hello World" when a user enters a specific command. The program uses the `flag` package from the Go standard library to define a command-line flag named `message`, which allows users to customize the output message.

## Usage

To use this program, you can save the code to a file with a `.go` extension (e.g., `cli.go`), and then compile it using the `go build` command:

`$ go build cli.go`

This will create an executable file named `cli` in the current directory. You can then run the program with the `./cli` command and provide the `message` flag to customize the output:

```
$ ./cli -message "Hello, world!"
Hello, world!
```

If you provide any additional arguments, the program will print an error message and exit 1:

```
$ ./cli -message "Hello, world!" extra argument
Error: unexpected argument(s)```




