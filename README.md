# Clipboard to Terminal (pbpaste for Linux)

This Go program replicates the functionality of the `pbpaste` command found on macOS, allowing you to easily retrieve the contents of your clipboard and display it in the terminal.

Since is done with Go you can compile for other platforms like x86 Windows 11 (included in the releases)

## Features

- Simple and straightforward command-line interface.
- Cross-platform compatibility - works on any Linux distribution.
- Dependable clipboard access using the `github.com/atotto/clipboard` package.

## Usage

To use the program, simply run:

```bash
./pbpaste
```

This will output the current content of your clipboard directly in the terminal.

Alternative
While you can always achieve similar functionality using the command:

```bash
xclip -o
```

Using this Go program adds a fun and educational twist to the simple task of accessing clipboard content. It's a great opportunity to explore and enhance your Go programming skills!

## Installation
Ensure you have Go installed on your system.

Clone this repository:

```bash
git clone https://github.com/Taikun/pbpaste.git
cd pbpaste
```
Build the program:
```bash
go build pbpaste.go
``` 

Run the program:
```bash
./pbpaste
```

Enjoy the clipboard magic in your terminal with this Go tool!