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

When writing a file, you need to convert the data to a byte slice before writing it to the file.

Type conversion works just like in Java/C++

Setup a test file by running "go mod init <module_name>" -> This is required to run a test function from within VSCode and/or run "go test" from the terminal

Go is not like using a testing framework like other languages. Instead, it has a built-in testing package that is used to write tests.
* A test file must have _test.go
* To run all tests in a package, run the command "go test"
* Tests should start with capitals and "Test"
* Regular functions should be camel case

EvenOrOdd Project completed for this section.


## Section 4: Organizing Data with Structs

All about structs. They are just a collection of properties that are related together. Exactly like in C. You can combine structs with receiver functions to essentially create classes.

We can embed structs within structs. They can also have receiver functions.

Structs are used as pointers. This is exactly how it is managed in C.

A * in front of a type means wer are working with a pointer to that type. A * in front of the variable name means we are working with the value that the pointer is pointing to.

Arrays are primitive data structures, can't be resized, and they are rarely used since they cannot be resized.

Slices can grow and shrink and are used in 99% of the time for lists of elements. Note that the slice is a reference to the underlying array, so if you modify the slice, you are modifying the underlying array. AKA a reference type.

Value types include int, float, stirng, bool, and structs.
Reference types include slices, maps, channels, pointers, and functions.

The dot operator will automatically dereference the pointer for you (no arrow function like in C). You just need to know what is pass by reference and what is pass by value to read the code accordignly.


## Section 5: Maps

Maps are like dicts in Python and identical to maps in C++. Note that the keys must all be of the same type and the values must all be of the same type.

Some differences between maps and structs.
* Maps
    * All keys must be the same type
    * All values must be the same type
    * Keys are indexed - we can iterate over them
    * Use to represent a collection of related properties
    * Don't need to know all the keys at compile time
    * Reference type!
* Stucts
    * Values can be of different types
    * Keys don't support indexing
    * Value type!
    * You need to know all the different fields at compile time
    * Use to represent a "thing" with a lot of different properties


 ## Section 6: Interfaces

A key topic in the Go language.

We know that...
* Every value has a type.
* Every function has to specify the type of its arguments.

If we use receivers without interfaces, then we have to reuse the same code but define the functions with the various types we want to use.

An interface is treated as a type. In the exmaple, we defined a bot interface that requires a getGreeting() string function. Then we can use that interface as a type for a function argument to allow for passing types that adhere to the interface as arguments.

We have concrete types like map, struct, int, string, etc. We also have abstract types like interfaces (interface type).

Other interface notes:
* Interfaces are not generic types (Go does not hae generics)
* Interfaces are 'implicit' (We dont manually have ot say that our custom type satisfies some interface)
* Interfaces are a contract to help us manage types (Garbage in -> Garbage out - If our custom type's implementation of a function is broken then interfaces won't help us)
* Interfaces are tough to read 

We can place a interface as a value meaning that the value can be any type that satisfies the interface.

We can also have an interface contain a group of other interfaces.

We can write an interface function that takes one or more types and then returns a common type that works for all the combinations but utilizes the same logic.


## Section 7: Channels and Go Routines

All about concurrent programming!

When a go program is executed, one go routine is started. If we place the keyword "go" in front of a function, the function will be executed within it's own go routine!

The go scheduler runs one routine until it finishes or makes a blocking call (like an HTTP request). 

Note that with one CPU, we are never truly running multiple routines at the same time. We are just switching between them based on the CPU clock speed and the scheduler.

The scheduler will run one thread on each "logical" core.

* Concurrency - We can have multiple threads executing code. If one thread blocks, another one is picked up and worked on.
* Parallelism - Multiple threads executed at the exact same time. Requires multiple CPU's.
