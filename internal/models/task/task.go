package task

import "time"

const (
	NOT_URGENT = iota
	DUE_SOON
	OVERDUE
)

type Task struct {
	ID          uint      `db:"id,omitempty" json:"id" rql:"filter,sort"`
	StatusID    uint      `db:"status_id" json:"status_id" rql:"filter,sort"`
	UserID      uint      `db:"user_id" json:"user_id" rql:"filter,sort"`
	Name        string    `db:"name" json:"name" valid:"required" rql:"filter,sort"`
	Description string    `db:"description" json:"description" valid:"required"`
	DueDate     time.Time `db:"due_date" json:"due_date" valid:"required" rql:"sort,layout=2006-01-02 15:04"`
	CreatedAt   time.Time `db:"created_at,omitempty" json:"created_at,omitempty" rql:"sort,layout=2006-01-02 15:04"`
	UpdatedAt   time.Time `db:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Pagination struct {
	Data []Task `json:"data"`
	Meta struct {
		PerPage      uint   `json:"per_page"`
		TotalResults uint64 `json:"total_results"`
		LastPage     uint   `json:"last_page"`
		CurrentPage  uint   `json:"current_page"`
	} `json:"meta"`
}
