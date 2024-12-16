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


