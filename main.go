package main

import (
	"bufio"
	"fmt"
	"os"
)

var current, direction string

func main() {

	compass := [4]string{"NORTH", "EAST", "SOUTH", "WEST"} // reserved on system

	fmt.Println("DIRECTIONS: =", compass)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Your Current Direction = ")
	scanner.Scan()
	current = scanner.Text()
	if validateCurrent(current) == 1 { // validation for user input . Must be NORTH/EAST/SOUTh/WEST
		return // if return 1 , out from program, else continue next step
	}

	fmt.Print("Which side you want to know = ")
	scanner.Scan()
	direction = scanner.Text()
	if validateDirection(direction) == 1 { // validation for user input . Must be LEFT/RIGHT
		return // if return 1 , out from program, else continue next step
	}

	var result = result(compass, current, direction)
	fmt.Printf("On Your %s is %s \n", direction, result)

}

func validateCurrent(x string) int {

	switch current {
	case "NORTH":
	case "EAST":
	case "SOUTH":
	case "WEST":
	default:
		fmt.Println("==================================================")
		fmt.Println("[YOU MUST INPUT NORTH/EAST/SOUTH/WEST (UPPERCASE)]")
		fmt.Println("==================================================")
		return 1
	}
	return 0

}

func validateDirection(x string) int {
	switch direction {
	case "LEFT":
	case "RIGHT":
	default:
		fmt.Println("==========================================")
		fmt.Println("[DIRECTION MUST BE LEFT/RIGHT (UPPERCASE)]")
		fmt.Println("==========================================")
		return 1
	}
	return 0
}

func result(compass [4]string, current string, direction string) (result string) {
	for i := 0; i < len(compass); i++ {
		switch current { // get current position
		case compass[i]:
			if direction == "LEFT" {
				if compass[i] == "NORTH" { // Start from NORTH and last is WEST, so if the current is NORTH, LEFT is WEST
					result = "WEST"
				} else {
					result = compass[i-1] // "Left" mean the array get minus 1 ( -1 )
				}
			} else {
				if compass[i] == "WEST" { // Start from NORTH and last is WEST, so if the current is WEST, RIGTH is NORTH
					result = "NORTH"
				} else {
					result = compass[i+1] // "Right" mean the array get plus 1 ( -1 )
				}
			}
		}
	}
	return result
}
