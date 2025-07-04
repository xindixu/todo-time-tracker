package goutils

import (
	"fmt"

	"todo-time-tracker/db/models"
	"todo-time-tracker/proto/go/model"
)

// -- Task --

// TaskStatusPBToDB converts a protobuf task status to a database task status
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

// TaskStatusDBToPB converts a database task status to a protobuf task status
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
		return model.Task_STATUS_NONE, fmt.Errorf("invalid task status: %s", status)
	}
}

// TaskLinkPBToDB converts a protobuf task link to a database task link
func TaskLinkPBToDB(link model.Task_Link) (models.TaskLink, error) {
	switch link {
	case model.Task_PARENT_OF:
		return models.TaskLinkParentOf, nil
	case model.Task_BLOCKS:
		return models.TaskLinkBlocks, nil
	case model.Task_RELATES_TO:
		return models.TaskLinkRelatesTo, nil
	case model.Task_DUPLICATE_OF:
		return models.TaskLinkDuplicateOf, nil
	default:
		return models.TaskLinkInvalid, fmt.Errorf("invalid task link: %s", link)
	}
}

// TaskLinkDBToPB converts a database task link to a protobuf task link
func TaskLinkDBToPB(link models.TaskLink) (model.Task_Link, error) {
	switch link {
	case models.TaskLinkParentOf:
		return model.Task_PARENT_OF, nil
	case models.TaskLinkBlocks:
		return model.Task_BLOCKS, nil
	case models.TaskLinkRelatesTo:
		return model.Task_RELATES_TO, nil
	case models.TaskLinkDuplicateOf:
		return model.Task_DUPLICATE_OF, nil
	default:
		return model.Task_LINK_NONE, fmt.Errorf("invalid task link: %s", link)
	}
}
