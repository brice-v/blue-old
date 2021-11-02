# TODO

- [ ] Display() should probably take in an int for indent level
    - That way subsequent calls will have the proper indentation
    - Maybe just jsonify it so we can make test programs?
- [ ] Parser should use lexer spans for better error messages
- [ ] Separate out logic in ast, and parser
    - Separate files for different nodes for ast
    - Separate files for private vs public vs utils
- [ ] Have Parser Print out to stdout with indentation
    - Include 'b' test programs to see if parser is working well
- [ ] Refactor and cleanup code in lexer, ast, and parser
    - Make functions make more sense, add helpers where necessary
- [ ] Add more ast and parser tests
- [ ] Generate some form of docs, whether to stdout or HTML
- [ ] Start analyzing how this will translate to go
- [ ] Will need to add back imports at some point
- [ ] For loop parsing should work pretty much like python or go
    - `blue - for i in 1 .. 10 {`
    - `go - for i := 0; i < 10; i++ {`
    - `python - for i in range(10):`
    - We want to be able to define for loops in a similar way, no parens should be needed
- [ ] If expressions should not need parens
- [ ] Global vars for some things like ENV, ARGV, STDOUT, STDIN, STDERR, etc.
- [ ] Proper immutability
- [ ] Remove lambdas, just use `fun() {}`
- [ ] support all functions using dot call syntax ie:
    ```
        fun hello(arg1, arg2) {
            println("hello #{arg1} and #{arg2}")
        }

        val x = "Brice"
        val y = "Name"

        x.hello(y)
    ```
- [ ] Dont really want `null` can we possibly use optionals? Some()/None?
- [ ] Printing formatting and easy of use debugging printing like `"#{=obj}"`
    - the `=` sign here should show the string version of the object, or something along those lines

## Future TODOs

- [ ] Figure out errors and how they will be handled
- [ ] Robust CLI for building, getting packages, running from CLI
- [ ] Reading input from cli
- [ ] http client/server - should be easy to get content from page
- [ ] regex
    - [ ] Builtin regex like JS/ruby?  `/.*word$/g`
- [ ] File IO
- [ ] Shell commands (maybe some cross platform alternatives in go with unix names ie. `rm`, `ls`, etc.)
- [ ] Multiple assignment `val x, var y = get_two_values(); val a, b = get_two_values();`
- [ ] Make sure default args work
- [ ] Test framework builtin
- [ ] Doc framework builtin
    - [ ] Should be easy enough to use, maybe like python where you can type `help()` to get info
- [ ] Simple GUI setup that works with build, using fyne?
- [ ] Config reading and writing - maybe use `.toml`?
- [ ] To/From JSON easily - maybe custom operator - probably just a function
- [ ] Automating browser?
- [ ] Types?
- [ ] Definitely want arbitrary precision numbers but easy to use like python
- [ ] Symbols? (`:symbol_name`)
- [ ] Enums? - Maybe this works with symbols/match somehow?
- [ ] Async code, channels, send and receive
- [ ] Package Manager
- [ ] Integrating with Go code
- [ ] Datetime library
- [ ] CLI library
- [ ] Mnesia like in memory db, something like redis but for this lang specifically
- [ ] Supervisors and OTP like concepts?
- [ ] ORM/SQL support - builtin support for sqlite would be nice
- [ ] Embed all to one binary
- [ ] Macros of some sort?
- [ ] Benchmarks - use builtin go?
- [ ] Performance test the implementation and improve
- [ ] Self host?? - realistically this isnt necessary but it would give a lot of insight into anything missing