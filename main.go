package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Typing(name string, fields map[string]string, callback func(map[string]interface{})) {
	fmt.Printf("\n==== %s ====\n", name)
	data := make(map[string]interface{})

	for field, dataType := range fields {
		data[field] = ReadDynamic(field, dataType)
	}

	callback(data)
}

func ReadDynamic(prompt, dataType string) interface{} {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch dataType {
		case "string":
			return input
		case "int":
			if num, err := strconv.Atoi(input); err == nil {
				return num
			}
			fmt.Println("Invalid input, please enter a valid number.")
		case "bool":
			lowerInput := strings.ToLower(input)
			if lowerInput == "yes" || lowerInput == "true" {
				return true
			} else if lowerInput == "no" || lowerInput == "false" {
				return false
			}
			fmt.Println("Invalid input, please enter 'yes' or 'no'.")
		default:
			fmt.Println("Unknown data type. Returning as string.")
			return input
		}
	}
}

func Choice(name string, options map[string]func()) {
	keys := []string{}

	for key := range options {
		keys = append(keys, key)
	}

	for {
		fmt.Printf("\n==== %s ====\n", name)
		for i, key := range keys {
			fmt.Printf("%d. %s\n", i+1, key)
		}
		fmt.Print("Choice: ")

		choice := readInput()
		index := stringToInt(choice)

		if index > 0 && index <= len(keys) {
			options[keys[index-1]]()
		} else {
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func stringToInt(s string) int {
	var num int
	fmt.Sscanf(s, "%d", &num)
	return num
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func Search(data map[string]interface{}){

}

func main() {
	fmt.Println("Welcome to the Dynamic CLI!")

	Choice("<Name of the>", map[string]func(){
		//Adding functions here like this 
		//"<Name>" : <Function Name>,
	})

	fmt.Println("Welcome to the Dynamic CLI!")

	Typing("Add User", map[string]string{
		//Adding data name and type
		//"<Name of the data>" : "<Type of data>",
	}, /*and here you add the anem of the function you wanna send to <Search>*/ Search)
}