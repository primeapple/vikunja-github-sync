package task

type Task interface {
	Id() string
}

type Tasks = []Task

