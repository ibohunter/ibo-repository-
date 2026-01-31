package simple_sql

import "time"

type TaskModel struct {
	ID          int
	Title       string
	Discussion  string
	Completed   bool
	CreatedAt   *time.Time
	CompletedAt *time.Time
}
