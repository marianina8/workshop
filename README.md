# Building Modern CLI Applications in Go

This workshop is designed to teach developers the complete process of designing, building, and testing a modern Command Line Interface (CLI) application. Students will learn how to build a lightweight CLI application from scratch and then explore the benefits of using the Cobra framework to increase development proficiency. They will also gain insights into building applications for different operating systems, using build tags for feature sets, and releasing and distributing applications across Unix, Linux, and Windows operating systems via popular package managers such as Homebrew and GoFish.

## Workshop Presentation:
If you want to follow along with my presentation on your own computer, please visit the slides online [here](https://docs.google.com/presentation/d/e/2PACX-1vQ8w3lODjTfvlulM-7X0cY4eyDcoqSMECyatzT3wh8VlUDrvxVvNbDOcB4YuAPDSODAAOn0Q15sbg1d/pub).

## To prepare for this workshop, please install the following applications:
- [Golang](https://go.dev/doc/install)
- [VSCode](https://code.visualstudio.com/download) or any other IDE.
- [Cobra CLI](https://github.com/spf13/cobra#usage)
- [GoReleaser](https://goreleaser.com/install/)

Also, please make sure to have a [GitHub](https://github.com/) account for Session 4.

## Examples:

### Session 1: Intro to Golang and CLI Basics
- [Example 1 - Build a CLI using the Go standard library “hello world!”](session1/examples/example1)
- [Example 2 - Create a CLI that takes in a file input, reads in the file, then outputs the content](session1/examples/example2)
- [Example 3 - Run a couple commands and to show command (verb), argument (noun), and flag (adjective)](session1/examples/example3)
- [Example 4 - Modify the CLI to validate that the input is of JSON format.](session1/examples/example4)
- [Example 5 - Modify the CLI to output in YAML format](session1/examples/example5)
- [Example 6 - Refactor code to use modularization](session1/examples/example6)
- [Example 7 - Convert flat structure to group by context](session1/examples/example7)
- [Example 8 -  Pass an invalid file, watch an error returned.  Rewrite error to use clear and concise language.](session1/examples/example8)
- [Example 9 -  Add flags to customize output, silence flag, or output to specify the name of an output file.](session1/examples/example9)

### Session 2: Designing, developing, and testing a CLI tool
- [Example 10 - Create a simple CLI with cobra-cli](session2/examples/example10)
- [Example 11 - Create a convert command with three flags: filename, silence, and output](session2/examples/example11)
- [Example 12 - Nested commands](session2/examples/example12)
- [Example 13 - Global, local, and required flags](session2/examples/example13)
- [Example 14 - Error handling](https://drive.google.com/file/d/1c-hmiDDxfH8fb2GXQXRyvOIZZcnXbSOt/view?usp=sharing)
- [Example 15 - Autogeneration of man pages](session2/examples/example15)
- [Example 16 - Intelligent suggestions](https://drive.google.com/file/d/1zx7IdwxD18z1hT5EdJDafTeVS503TmxV/view?usp=sharing)
- [Example 17 - Write test for Cobra CLI command](session2/examples/example17)
- [Example 18 - Show how changing code breaks test](https://drive.google.com/file/d/1-McxHBkDHAL7JjU7BwbuV5dkgJuzqzDd/view?usp=sharing)

### Session 3: Integrating APIs and Third Party Libraries
- [Example 19 - Create new store command](session3/examples/example19/)
- [Example 20 - Create request to https://jsonbin.io/](session3/examples/example19/)
- [Example 21 - Handle response](session3/examples/example19/)
- [Example 22 - Add survey](session3/examples/example22/)
- [Example 24 - Create command and create termindal dashboard layout](session3/examples/example24/)
- [Example 25 - Add components to the terminal dashboard layout](session3/examples/example25/)
- [Example 26 - Add styling to the terminal dashboard layout](session3/examples/example26/)

#### Session 3: Framework links
- [Survey](https://github.com/AlecAivazis/survey/v2)
- [Color](https://github.com/fatih/color)
- [Bubbletea (TUI Framework)](https://github.com/charmbracelet/bubbletea)

### Session 4: Documentation, Deployment, and Distribution
- [Example 27 - Building with tags](session4/examples/example27/)
- [Example 28 - Compiling with GOOS/GOARCH]([session4/examples/example27/](https://drive.google.com/file/d/1KqAEeXA-d31VEZHoR0xUOoGVxtB83JRO/view?usp=sharing))
- [Example 29 - Test on a different platform with Docker](session4/examples/example29/)
- [Examples 30-36 - Distributing with Homebrew](https://github.com/marianina8/cli)

## Videos
[Click here](https://drive.google.com/drive/folders/1MPDlHbxo6MHF4lRr_IZWXn3G5ALAHpJ-?usp=sharing) to view all videos from the presentation. 

## Handouts
- [History of a CLI](/handouts/history.pdf)
- [Good vs Bad CLI Design](/handouts/good_vs_bad_cli.png)
- [Common Program Layouts](/handouts/layouts.pdf)
- [Operating System Differences](/handouts/OSdifferences.jpg)