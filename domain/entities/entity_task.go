package entities

import "time"

type TaskCompletionStatus int64

const (
	TaskNotFinished = 0
	TaskFinished    = 1
)

type Board struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedBy   User       `json:"created_by"`
	Users       []User     `json:"users"`
	TaskLists   []TaskList `json:"task_lists"`
	StatusCode  int        `json:"status_code"`
	CreatedAt   time.Time  `json:"created_at"`
	ModifiedAt  time.Time  `json:"modified_at"`
}

type TaskList struct {
	ID          int       `json:"id"`
	IDBoard     int       `json:"id_board"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   User      `json:"created_by"`
	Tasks       []Task    `json:"tasks"`
	StatusCode  int       `json:"status_code"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

type Task struct {
	ID          int                  `json:"id"`
	IDTaskList  int                  `json:"id_task_list"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	CreatedBy   User                 `json:"created_by"`
	Status      TaskCompletionStatus `json:"status"`
	StatusCode  int                  `json:"status_code"`
	CreatedAt   time.Time            `json:"created_at"`
	ModifiedAt  time.Time            `json:"modified_at"`
}
