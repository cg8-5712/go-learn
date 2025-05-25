# Command-Line Calculator Program Specification

**Objective**  
Implement a simple command-line calculator in Go that supports basic arithmetic operations with robust error handling and type conversions.

---

## Core Requirements

### 1. Input Handling (`fmt.Scan`)
- Read user input from the command line using `fmt.Scan`/`fmt.Scanln`
- Accept input in the format: `<number1> <operator> <number2>`
- Example valid input: `5 + 3` or `12.5 / 2`

### 2. Supported Operations
| Operator | Description  |
|----------|--------------|
| `+`      | Addition     |
| `-`      | Subtraction  |
| `*`      | Multiplication |
| `/`      | Division     |

### 3. Error Handling
Implement comprehensive error checking for:
- Invalid numeric inputs (non-number values)
- Division by zero attempts
- Unsupported operators
- Incorrect number of arguments

### 4. Type Conversion
- Convert input strings to `float64` for calculations
- Format output to display:
    - Integer results as whole numbers (e.g., `5` instead of `5.0`)
    - Floating-point results with 2 decimal places (e.g., `3.14`)

---

## Technical Implementation

### Program Flow
```text
Start
  │
  ├─ Read user input
  │   └─ Error: Invalid format → Show message
  │
  ├─ Convert operands to float64
  │   └─ Error: Non-numeric → Show message
  │
  ├─ Validate operator
  │   └─ Error: Invalid op → Show message
  │
  ├─ Perform calculation
  │   └─ Error: Division by 0 → Show message
  │
  └─ Display formatted result
```