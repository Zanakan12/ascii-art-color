
# ASCII Art Project

## Description
This project, named ASCII Art, is a Go application that transforms printable characters into ASCII art. The application takes a string input and outputs an artistic representation of that string using a specified font stored in text files.

## Project Structure
- **.git/**: Directory for Git versioning.
- **banner/**: Contains font files in text format used to generate ASCII art.
- **ext/**: An external module containing the logic to combine characters with fonts and generate ASCII art.
- **go.mod**: Dependency management file for Go.
- **main.go**: The main script that drives the application.

```sh
ascii-art/
├── go.mod
├── main.go
├── banner/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
└── ext/
    ├── basefonction
    │       
    |
    | 
    ├── colors
    ├── justify
    ├── output
    ├── resverse
    └── template

```

## How to Use
1. **Input**: The user needs to provide a string of characters as an argument when running the program. A path to a specific font file can also be given as a second argument.

2. **Execution**:
   - Example command:
     ```
     go run main.go "Hello World"
     ```
     Or with a specific font file:
     ```
     go run main.go "Hello World" "shadow or shadow.txt"
     ```

3. **Output**: The application outputs multiple lines of text, representing the ASCII art of the provided string, according to the chosen font.

## Dependencies
- The project requires Go (Golang) to run.
- No external libraries are needed outside of those specified in `go.mod`.

## Development and Contribution
Developers interested in contributing to the project can clone the Git repository and follow standard Go contribution guidelines, including testing and code reviews.

## License
This project is licensed under the Creative Commons Attribution-NonCommercial-NoDerivatives 4.0 International (CC BY-NC-ND 4.0) license. See the LICENSE file for more details.
