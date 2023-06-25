# Example 6 - Refactor code to use modularization

This simple CLI requests input for a file that's contents match a specific format structure: 

```
{
    "title": "Workshop Title",
    "instructors": [
        {
            "name": "Instructor 1",
            "email": "instructor1@example.com"
        },
        {
            "name": "Instructor 2",
            "email": "instructor2@example.com"
        }
    ]
}
```

If the input file validates against this structure then the data is converted to YAML format and then printed out, otherwise the application will exit 1.


## Usage

To use this program, build the `cli.go` file:

`$ go build -o cli .` // Unix, Linux
`$ go build -o cli.exe .` // Windows

This will create an executable file named `cli` in the current directory. You can then run the program with the `./cli` command and provide the `file` flag to print its contents:

`./cli -file <filename>` // Unix, Linux
`.\cli.exe -file <filename>` // Windows

Replace `<filename>` with the name of the file you want to read.

If the file cannot be opened or read, an error message will be printed and the program will exit 1.  If the file can be read, the file's contents will be printed out.