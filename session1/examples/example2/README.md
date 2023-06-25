# Example 2 - Create a CLI that takes in a file input, reads in the file, then outputs the content

This simple CLI requests input for a file name, which it usest to output the result.

## Usage

To use this program, build the `cli.go` file:

`$ go build cli.go` // Unix, Linux
`$ go build -o cli.exe cli.go` // Windows

This will create an executable file named `cli` in the current directory. You can then run the program with the `./cli` command and provide the `file` flag to print its contents:

`./cli -file <filename>` // Unix, Linux
`.\cli.exe -file <filename>` // Windows

Replace `<filename>` with the name of the file you want to read.

If the file cannot be opened or read, an error message will be printed and the program will exit 1.  If the file can be read, the file's contents will be printed out.