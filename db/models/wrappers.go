package models

type TaskWrapper struct {
	Task *Task
	Tags []*Tag
}
