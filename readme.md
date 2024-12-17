# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
| What? | When? |
| ---- | --- |
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (20 mins) |
| Lunch Break | 12:30 PM (1 hr) |
| Tea Break | 3:00 PM (20 mins) |
| Wind up | 5:00 PM |

## Software Requirements
- Go Tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)

### Verification
```shell
go version
```

## Methodology
- No powerpoint
- Code & Discussion

## Repository
- https://github.com/tkmagesh/Cisco-GoFoundation-Dec-2024

## Books
- Structure & Interpretation of Computer Programs
    - https://web.mit.edu/6.001/6.037/sicp.pdf


## Why Go?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers (No public/private/protected keywords)
    - No reference types (Everything is a value)
    - No pointer arithmatic 
    - No classes (only structs)
    - No inheritance (only composition)
    - No exceptions (only errors)
    - No try..catch..finally construct
    - No implicity type conversions
- Performance
    - Equivalent to C++
    - Compiled to Machine Code (No JIT compilation)
    - Tooling support for cross-compilation
- Managed Concurrency
    - Builtin Scheduler
    - A concurrent operation is represented as a "Goroutine"
    - Goroutines are cheap (~4 KB) while the OS Threads are costly (~2 MB)
    - Support for concurrency is built in the language
        - "go" keyword
        - channel "data type"
        - "<-" (channel "operator")
        - "range" construct
        - "select-case" construct
    - Standard Library API
        - "sync" package
        - "sync/atomic" package

## Compilation & Build
### Compile 
```shell
go build <filename.go>
```

### Compile & Execute
```shell
go run <filename.go>
```

### Cross compilation
#### To get the complete list of environment variables
```shell
go env
```

#### To get the list of supported platforms
```shell
go tool dist list
```

#### Cross compilation
```shell
 GOOS=windows GOARCH=amd64 go build 01-hello-world.go
```

## Data Types
- string
- bool
- integers
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points
    - float32
    - float64
- complex
    - complex64 ( real[float32] + imaginary[float32] )
    - complex128 ( real[float64] + imaginary[float64] )
- alias
    - byte (alias for unsigned int)
    - rune (alias for unicode code point)

### Zero values
| Data Type | Zero value |
------------ | ------------- |
|int family     | 0 |
|uint family    | 0 |
|complex family | (0+0i) |
|string         | "" (empty string) |
|bool           | false |
|byte           | 0 |
|interface      | nil |
|pointer        | nil |
|function       | nil |
|struct         | struct instance |

## Variables
### Function scope
- CAN use ":="
- CANNOT have "unused" variables
### Package scope
- CANNOT use ":="
- CAN have "unused" variables

## Functions
    - Can return more than 1 result
        - Supports "Named results"
    - Anonymous functions
        - Functions defined in another function
        - Cannot have a name
        - Must be immediately invoked
    - Variadic functions
        - Can accept varying number of arguments
    - Higher Order Functions
        - Assign a function as a value to a variable
        - Pass a function as an argument
        - Return a function as a return value
    - Deferred functions
        - Postpone the execution of a function until the current function execution is completed

## Collections
- Array
    - Fixed sized typed collecton
- Slice
    - Varying sized typed collection
    - A slice a pointer to an underlying array
- Map
    - typed collection of key/value pairs

## Error Handling
    - Errors are values returned from a function
    - By Convention, errors are objects implementing "error" interface
    - Creation
        - errors.New()
        - fmt.Errorf()
        - Custom type implementing "error" interface

## Panic & Recovery
### Panic
- State of the application where the application execution cannot proceed further
- Application is gradually brought to a halt after executing all the deferred functions
- using "panic()" to create a panic
    - By convention, an error is used as an argument for the panic() function
### Recovery
- "recover()" returns the error (argument to the panic() function) if there is a panic
- Apt to be used in the deferred functions

## Modules & Packages
### Module
- Any code that need to versioned and deployed together
- Typically, a folder with go.mod file
- go.mod
    - manifest file for the module
    - name
        - typically, includes the complete repo path
    - go runtime version
    - dependencies
#### Create a module
```shell
go mod init <module_name>
# ex:
go mod init github.com/tkmagesh/cisco-gofoundation-dec-2024/11-modularity-demo
```

#### Execute a module
```shell
go run .
```

#### Create a build
```shell
go build .
# OR
go build -o <binary_name> .
```

### Package
- Internal organization of code in a module
- Just folders
- Can be nested
- Scope can be defined at the package level
