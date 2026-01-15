# GoPracticePrograms
Programs used while learning the Go programming language.

## Section 1: Getting Started

* Download Go - https://go.dev/dl/
* Install VSCode - https://code.visualstudio.com/
* Install Go extension for VSCode - https://marketplace.visualstudio.com/items?itemName=ms-vscode.go
* Install various tools for Go - https://go.dev/doc/tool_install

Gopls, undeclared name, type and module errors in VSCode:
* gopls was not able to find modules in your workspace -> This can happen if you loaded a directory that contains mulitple Go projects into your editor workspace.
The simplest solution to resolve these errors would be to instead open only the single Go project directory you are currently working on in the editor.


## Section 2: Simple Start

How to run the code?
* go run main.go
* go build -> compiles the code and creates an executable
* go run -> compiles and runs the code
* go fmt -> formats the code
* go install -> compiles and installs a package
* go get -> downloads a package
* go test -> runs tests
* go clean -> cleans up the workspace
* go env -> shows environment variables
* go version -> shows version

What is package main?
Package == Project == Workspace
* A collection of common grouped source files
* The very first line of each file must define the name of the package
* Two Types of Packages:
    * Executable - gneerates a file that we can run
    * Reusable - Code used as 'helpers' -> Good place to put reusable logic
* The name of the package determines the type of package that is being created -> main creates an executable, other names create reusable packages 
* package main - Defines a pakcage that can be compiled and then "executed" -> must have a func called 'main'
* package calculator (reusable package) - Defines a package that cn be used as a dependency (helper code)
* package uploader (reusable package) - Defines a package that can be used as a dependency (helper code)

What does import "fmt" do?
* Imports a package that is stored elsewhere (in this case, the standard library)
* You can see all the standard library packages here: https://pkg.go.dev/std

What is func?
* Defines a function as in any other main coding language

How is the main.go file organized?
* package declaration -> import other packages we need -> declare functions, tell Go to do things


## Section 3: Deeper into Go

A project that will be a package that will simulate a deck of playing cards.
* newDeck -> Create a list of playing cards. Essentially an array of strings
* print -> Log out the contents of a deck of cards
* shuffle -> Shuffles all the cards in a deck
* deal -> Create a 'hand' of cards
* saveToFile -> Save a list of cards to a file on the local machine
* newDeckFromFile -> Load a list of cards form the local machine

Go is a statically typed language.

Basic Go Types:
* bool
* string
* int
* float64
(this is not a complete list)

Note that you cannot initialize a variable outside of a function. Note that the variable can be initialized, just not defined outside of a function.

Go has arrays and slices. A slice is an array that can grow or shrink in size dynamically. Note that every element in a slice must be of the same type.

Go is not object oriented. In go, instead of classes, we can "extend" a base type ad add some extra fucntionality to it. For example, we can tell Go
to create an array fo strings and add a bunch fo functions specifically made to work with it. A function with a reciever is like a "method" - a function that belongs to an "instance".

Note: You need to ensure that all files are included in go run when running a project
Note: When a function has a receiver of a particular type, any variable with that type gets access to that particular function

Indexing in Go is just like in Python.