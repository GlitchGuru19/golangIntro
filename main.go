package main

import (
	"fmt"
)

// This is a struct which has exported fields
type Person struct {
	Name     string
	Role     string
	Age      int
	IsActive bool
}

var people []Person // Slice to store people's records

func displayMenu() {
	fmt.Println("\nWelcome to the People directory")
	fmt.Println()
	fmt.Println("please choose an option")
	fmt.Println("1. Add Person")
	fmt.Println("2. List People")
	fmt.Println("3. Activate Person")
	fmt.Println("4. Deactivate Person")
	fmt.Println("5. Exit")
}

// Function to add the books
// Still figuring out how to add this thing
func addPerson() {

	var (
		name string
		role string
		age  int
	)

	fmt.Println()
	fmt.Println("Add a new Person")
	fmt.Print("Enter the name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter Role (e.g., Member, Leader): ")
	fmt.Scanln(&role)
	fmt.Print("Enter the age: ")
	fmt.Scanln(&age)

	newPerson := Person{
		Name:     name,
		Role:     role,
		Age:      age,
		IsActive: false, // New person starts as inactive
	}
	people = append(people, newPerson)

	fmt.Println("Person added successfully.")
}

func activatePerson() {
	if len(people) == 0 {
		fmt.Println("No people to activate. Add someone first!")
		return
	}

	var number int
	fmt.Print("Enter the number of the person to activate: ")
	fmt.Scanln(&number)

	// Convert to zero-based index
	index := number - 1

	// Check if the number is valid
	if index < 0 || index >= len(people) {
		fmt.Println("Invalid number. Please try again.")
		return
	}

	// Check if already active
	if people[index].IsActive {
		fmt.Printf("%s is already active!\n", people[index].Name)
		return
	}

	// Mark as active
	people[index].IsActive = true
	fmt.Printf("%s is now active!\n", people[index].Name)
}

func deactivatePerson() {
	if len(people) == 0 {
		fmt.Println("No people in the directory. Add someone first!")
		return
	}

	var number int
	fmt.Print("Enter the number of the person to deactivate: ")
	fmt.Scanln(&number)

	index := number - 1

	if index < 0 || index >= len(people) {
		fmt.Println("Invalid number. Please try again.")
		return
	}

	if !people[index].IsActive {
		fmt.Printf("%s is already inactive!\n", people[index].Name)
		return
	}

	people[index].IsActive = false
	fmt.Printf("%s is now inactive.\n", people[index].Name)
}

func listPeople() {
	if len(people) == 0 {
		fmt.Println("No people in the directory yet.")
		return
	}

	fmt.Println("List of all people:")
	for i, person := range people {
		status := "Inactive"
		if person.IsActive {
			status = "Active"
		}
		fmt.Printf("%d. Name: %s, Role: %s, Age: %d, Status: %s\n",
			i+1, person.Name, person.Role, person.Age, status)
	}
}

func selectMenu(){
	for {
		var option int

		displayMenu()
		fmt.Print("Option => ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			addPerson()
		case 2:
			listPeople()
		case 3:
			activatePerson()
		case 4:
			deactivatePerson()
		case 5:
			fmt.Println("Thank you for using the program.")
			fmt.Println()
			return
		default:
			fmt.Println("Wrong option. Please ty again.")

		}
	}
}

func main() {
	selectMenu()
	
}
