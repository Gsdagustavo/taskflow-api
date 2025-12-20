package entities

import "time"

type Board struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	TaskLists  []TaskList `json:"task_lists"`
	StatusCode int        `json:"status_code"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt time.Time  `json:"modified_at"`
}

type TaskList struct {
	ID         int       `json:"id"`
	IDBoard    int       `json:"id_board"`
	Name       string    `json:"name"`
	Tasks      []Task    `json:"tasks"`
	StatusCode int       `json:"status_code"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type Task struct {
	ID          int       `json:"id"`
	IDTaskList  int       `json:"id_task_list"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	StatusCode  int       `json:"status_code"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}
