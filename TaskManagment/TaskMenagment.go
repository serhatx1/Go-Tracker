package TaskManagment

import (
	"fmt"
)

func printingOperations() {
	fmt.Println("1-List all the tasks\n2-Add\n3-Change Progress\n4-Exit ")
	fmt.Println("Enter an operation: ")
}
func (task Task) Start() {
	var operation int
	var exitCheck bool = true
	printingOperations()

	for exitCheck {

		fmt.Scan(&operation)
		switch operation {
		case 1:
			task.ListTheTasks()
			printingOperations()

		case 2:
			err := task.AddTask()
			if err != nil {
				return
			} else {
				id++
			}

			printingOperations()

		case 3:
			err := task.ChangeProgress()
			if err != nil {
				return
			}
			printingOperations()

		case 4:
			fmt.Println("You are leaving....")
			exitCheck = false
		default:
			printingOperations()
		}
	}

}
