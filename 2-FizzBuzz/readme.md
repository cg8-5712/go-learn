## FizzBuzz Program Description

**FizzBuzz** is a classic programming exercise that primarily tests skills in loops, conditional statements, and formatted output. The program's objective is to iterate through integers from 1 to 100 and output corresponding results based on the following rules:

### Output Rules

- If a number is divisible by 3, output `Fizz`
- If a number is divisible by 5, output `Buzz`
- If a number is divisible by both 3 and 5, output `FizzBuzz`
- Otherwise, output the number itself

### Implementation Requirements

- Implement in Go language
- Use a `for` loop to iterate from 1 to 100
- Use `if/else` or `switch` statements to check divisibility
- Format output using `fmt.Printf` or `fmt.Println`
- Code should have clear structure, concise comments, and good readability

#### Summary
- else if need to be written after } in the same line, this is because of the Go language syntax.
- Go's Automatic Semicolon Insertion Rules, In Go, semicolons are automatically inserted by the lexer following these specific rules from the Go Language Specification. Like the ASI in JavaScript, Go has its own rules for semicolon insertion, which are important to understand for writing correct code.
- For simple go file, if use `go run main.go` to execute the program, its package must be `main` and have a main function because go interpret these as the entrance.