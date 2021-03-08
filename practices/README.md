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

- C style: /**/
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
- one-method interfaces should name with the method name plus an -er suffix or similar modification, like `Reader`, `Buffer`.

## Semicolons

Instead, the lexer uses a simple rule to insert semicolons automatically as it scans, so the input text is mostly free of them.

One consequence of the semicolon insertion rules is that you cannot put the opening brace of a control structure (if, for, switch, or select) on the next line. If you do, a semicolon will be inserted before the brace, which could cause unwanted effects. 
