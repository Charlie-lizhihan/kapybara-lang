# Kapybara Language🐾

（ｴ ´ ＼ ／ ｀ ｴ´）ww

Hey there! Welcome to Kapybara!

Kapybara is an interpreted language that's simple yet powerful, designed for those who are taking their first steps into language development. Inspired by "Writing An Interpreter In Go", Kapybara boasts modern language features wrapped in a syntax that’s easy to pick up.

## Features

Kapybara supports various features which include, but are not limited to:

- When you input is not aligh with kapybara-lang syntax, one capybara will pop up and tell you it can't handle your code
```
>>[1,2,3,4]
[1, 2, 3, 4]
>>let double = fn(x) {x*2};
>>[1, double(1), 3*3, 4-3]
[1, 2, 9, 1]
>>[1,2
／* ｀ ｴ´）www
Woops! Kapybara can't handle your code!
parser errors:
        expected next token to be ], got EOF instead
```

- Variables and constants for storing values.
- Data types like integers, booleans, strings, arrays, and hash maps.
- Arithmetic expressions (addition, subtraction, multiplication, etc.).
- Comparison expressions and boolean expressions.
- Control structures such as conditionals (if, else) and loops.
- Functions (including first-class and higher-order functions).
- Closures which allow for function encapsulation with environments.

## Quick Start

To run the Kapybara compiler:

1. Enter the `kapybara-lang` directory:

   ```sh
   cd path/to/kapybara-lang
   ```

1. Build the project using Go:

   ```
   go build
   ```

2. Run the compiler:

   ```
   go run main.go
   ```

## Syntax and Examples

### Let Bindings

Use `let` to bind values to identifiers:

```
let identifier = expression;
```

Example:

```
let name = "Kapybara";
let age = 1;
```

### If-Else Clause

Kapybara uses `if-else` clauses for control flow:

```
if (condition) {
  // code if condition is true
} else {
  // code if condition is false
}
```

Example:

```
let x = 10;
if (x < 20) {
  x + 5;
} else {
  x - 5;
}
// Outputs 15
```

### Data Structures

Kapybara supports arrays and hashes as fundamental data structures.

**Arrays:**

```
let myArray = [1, 2, 3, 4];
myArray[2]; // Accesses the element at index 2 (third element), which is 3
```

**Hashes:**

```
let myHash = {"name": "Kapybara", "age": 1};
myHash["name"]; // Accesses the value associated with the key "name", which is "Kapybara"
```

### Functions

Kapybara treats functions as first-class citizens:

```
let add = fn(a, b) { a + b; };
add(5, 5); // Outputs 10
```

## Contributing
Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make to the Kapybara project are **greatly appreciated**.

## License

Distributed under the MIT License. See `LICENSE` for more information.





