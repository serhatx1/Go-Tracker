package TaskManagment

import "fmt"

var id int = 1

func (task Task) AddTask() error {
	var title string
	fmt.Println("Please enter the task:\n")
	fmt.Scan(&title)
	if title == "" {
		return fmt.Errorf("Task title is a must,can't be empty")
	}
	newTask := Task{Id: id, Title: title, IsDone: false}
	allTasks = append(allTasks, newTask)
	return nil
}
