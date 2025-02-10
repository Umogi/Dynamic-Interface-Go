# Dynamic CLI

## Overview
Dynamic CLI is a Go-based command-line interface designed to handle user input dynamically. It provides flexible functions for user interaction, including typed input collection and menu-based choices.

## Features
- **Dynamic Input Handling:** Accepts user input based on specified data types (string, int, bool).
- **Interactive Menu System:** Allows users to choose from different options dynamically.
- **Custom Callbacks:** Processes collected input with user-defined functions.

## Installation
Ensure you have [Go](https://golang.org/dl/) installed. Then, clone the repository and navigate into it:

```sh
git clone https://github.com/Umogi/Dynamic-interface-Go.git
cd <project_directory>
```

## Usage
Run the program using:

```sh
go run main.go
```

## Code Structure
### Typing Function
```go
func Typing(name string, fields map[string]string, callback func(map[string]interface{}))
```
- Displays a prompt for user input.
- Collects data based on the specified types.
- Passes the collected data to a callback function.

### ReadDynamic Function
```go
func ReadDynamic(prompt, dataType string) interface{}
```
- Reads user input and converts it to the required type.

### Choice Function
```go
func Choice(name string, options map[string]func())
```
- Displays a list of choices.
- Executes the selected function.
  
## Example Configuration
You can modify the `Choice` and `Typing` functions to suit your application's needs:

```go
Choice("Main Menu", map[string]func(){
    "Add User": func() {
        Typing("User Details", map[string]string{
            "Name": "string",
            "Age": "int",
            "Admin": "bool",
        }, Search)
    },
})
```

## Contributing
Feel free to submit pull requests or open issues to suggest improvements!

## License
This project is licensed under the MIT [LICENSE](LICENSE).

