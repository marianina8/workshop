# Example 10 - Create a Cobra command

This is a simple Cobra CLI application that provides a single command called "message". It allows you to print a hello world message with an optional name.

## Usage

```
$ ./cli message
Hello World
```
or pass in a name with the long form flagname, --name:
```
$ ./cli message --name Jane
Hello World, Jane
```
or pass in a name with the shorthand flag, -n:
```
$ ./cli message -n John
Hello World, John
```
