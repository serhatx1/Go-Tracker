package TaskManagment

import (
	"errors"
	"fmt"
)

func (task Task) ChangeProgress() error {
	var id int
	task.ListTheTasks()
	fmt.Println("Enter the id of the task that you wanna change:")
	fmt.Scan(&id)
	for i := range allTasks {
		if allTasks[i].Id == id {
			allTasks[i].IsDone = !allTasks[i].IsDone
			return nil
		}
	}
	return errors.New("Task not found in the list.")
}
