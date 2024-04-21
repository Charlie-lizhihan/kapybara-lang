# Kapybara Language🐾

（ｴ ´ ＼ ／ ｀ ｴ´）ww

Hey there! Welcome to Kapybara!

Kapybara is an interpreted language that's simple yet powerful, designed for those who are taking their first steps into language development. Inspired by "Writing An Interpreter In Go", Kapybara boasts modern language features wrapped in a syntax that’s easy to pick up.

## Features

Kapybara supports various features which include, but are not limited to:

- Basic arithmetic and boolean expressions
- Control structures like `if-else`
- `let` bindings for variable assignment
- First-class functions with closures
- Built-in functions for working with strings, arrays, and hashes
- Error handling with descriptive messages

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





