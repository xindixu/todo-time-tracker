package goutils

import (
	"fmt"

	"todo-time-tracker/db/models"
	"todo-time-tracker/proto/go/model"
)

func TaskStatusPBToDB(status model.Task_Status) (models.TaskStatus, error) {
	switch status {
	case model.Task_TODO:
		return models.TaskStatusTodo, nil
	case model.Task_IN_PROGRESS:
		return models.TaskStatusInProgress, nil
	case model.Task_DONE:
		return models.TaskStatusDone, nil
	case model.Task_BLOCKED:
		return models.TaskStatusBlocked, nil
	default:
		return models.TaskStatusInvalid, fmt.Errorf("invalid task status: %s", status)
	}
}

func TaskStatusDBToPB(status models.TaskStatus) (model.Task_Status, error) {
	switch status {
	case models.TaskStatusTodo:
		return model.Task_TODO, nil
	case models.TaskStatusInProgress:
		return model.Task_IN_PROGRESS, nil
	case models.TaskStatusDone:
		return model.Task_DONE, nil
	case models.TaskStatusBlocked:
		return model.Task_BLOCKED, nil
	default:
		return model.Task_NONE, fmt.Errorf("invalid task status: %s", status)
	}
}
