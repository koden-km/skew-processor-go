package skew

import (
	"errors"
)

type TaskRegistry struct {
	tasks map[string]*TaskInterface
}

func (tr *TaskRegistry) Add(task TaskInterface) error {
	if tr.tasks[task.Name()] {
		return errors.New("Duplicate key error")
	}

	tr.tasks[task.Name] = task
	return nil
}

func (tr *TaskRegistry) Get(name string) *TaskInterface {
	return tr.tasks[task.Name()]
}

func (tr *TaskRegistry) Capabilities() []string {

	// TO DO

	// var capabilities []string = make([]string, len(tr.tasks))
	// for n, _ := range tr.tasks {
	// 	capabilities.Append(n)
	// }
	// return capabilities
}
