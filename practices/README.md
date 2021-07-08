# Effective GO

Learn from official examples(from library)

## Formatting

by using gofmt, we will get this

```go
package main

// 对齐
type T struct {
 name  string // name of the object
 value int    // its value
}
```

## Commentary

- C style: /\*\*/
- C++ style: //

Commentary in go mainly is using at describing package. If the package is simple, the package comment can be brief.  
See `fmt` package uses great comment as doc.

```go
// 1. Describe package

/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp

// Package path implements utility routines for
// manipulating slash-separated filename paths.

// 2. Describe function
type Regexp struct{}

// Compile parses a regular expression and returns, if successful, a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
 return nil, nil
}
```

## Names

- Make your package name short, concise and evocative. Avoid import a package with `.`.
- No need to make your getter func with `Get` prefix, Use Upper case is OK.
- one-method interfaces should name with the method name plus an -er suffix or similar modification, like `Reader`
  , `Buffer`.

## Semicolons

Instead, the lexer uses a simple rule to insert semicolons automatically as it scans, so the input text is mostly free
of them.

One consequence of the semicolon insertion rules is that you cannot put the opening brace of a control structure (if,
for, switch, or select) on the next line. If you do, a semicolon will be inserted before the brace, which could cause
unwanted effects.

## Control structures

There is no do or while loop, only a slightly generalized for; switch is more flexible; if and switch accept an optional
initialization statement like that of for; break and continue statements take an optional label to identify what to
break or continue; and there are new control structures including a type switch and a multiway communications
multiplexer, select. The syntax is also slightly different: there are no parentheses and the bodies must always be
brace-delimited.

```go
// type switch
package main

import "fmt"

func main() {
 var t interface{}
 t = ""
 switch t := t.(type) {
 default:
  fmt.Printf("unexpected type %T\n", t)
 case bool:
  fmt.Printf("boolean %t\n", t)
 case int:
  fmt.Printf("int %t\n", t)
 case *bool:
  fmt.Printf("ptr to boolean %t\n", t)
 case *int:
  fmt.Printf("ptr to int %t\n", t)
 }
}
```

## Functions

### Multiple Return Values

`func (file *File) Write(b []byte) (n int, err error)`

### Named Result Parameters

```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

### Defer

Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns. It's an unusual but effective way to deal with situations such as resources that must be released regardless of which path a function takes to return. The canonical examples are unlocking a mutex or closing a file.

```go
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```

## Data

### Allocation with `new`

It's a built-in function that allocates memory, but unlike its namesakes in some other languages it does not initialize the memory, it only zeros it. That is, new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type \*T. In Go terminology, it returns a pointer to a newly allocated zero value of type T.

### Constructors and composite literals

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

### Allocation with `make`

Back to allocation. The built-in function make(T, args) serves a purpose different from new(T). It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not \*T). The reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use.

## Initialization

### Constants

Constants in Go are just that—constant. They are created at compile time, even when defined as locals in functions, and can only be numbers, characters (runes), strings or booleans. Because of the compile-time restriction, the expressions that define them must be constant expressions, evaluatable by the compiler.

### Variables

Variables can be initialized just like constants but the initializer can be a general expression computed at run time.

### The init function

Finally, each source file can define its own niladic init function to set up whatever state is required. (Actually each file can have multiple init functions.) And finally means finally: init is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.

```go
func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line.
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}
```

## Methods

### Pointers vs. Values

```go
package o

type ByteSlice []byte

// With Values
func (s ByteSlice) Append(data []byte) []byte {
  return append(s, data...)
}

// With Pointers
func(p *ByteSlice) append(data []byte) {
 s := *p
 *p = append(s, data...)
}
```

## Interfaces and other types

### Interface

Interfaces in Go provide a way to specify the behavior of an object.

### Conversions

### Interface conversions and type assertions

```go
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}

var val interface{}
func main() {
 var valName = val.(map[string]interface{})["name"]
}
```

convert by using `value.(typeName)`, e.g. `str, ok := value.(string)`

## The blank Identifier
