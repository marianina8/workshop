# Example 4 - JSON File Validator

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

If the input file validates against this structure then the data is extracted and printed out, otherwise the application will exit 1.

## Usage

To use this program, build the `cli.go` file:

`$ go build cli.go`

This will create an executable file named `cli` in the current directory. You can then run the program with the `./cli` command and provide the `file` flag to print its contents:

```
$ ./cli -file <filename>
```

Replace `<filename>` with the name of the file you want to read.  An example file is available, `example.json`, with data within the expected structure and results in a successful execution.

If the file cannot be opened or read, an error message will be printed and the program will exit 1.  If the file can be read, the data will be printed out.