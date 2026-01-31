package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context,conn *pgx.Conn,) ([]TaskModel, error) {
	sqlQuery := `
	SELECT id, title, discussion, completed, created_at, completed_at
	FROM tasks
	ORDER BY id ASC;

	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Discussion,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)

		prinTaskRow(task)
	}
	return tasks, nil
}

func prinTaskRow(task TaskModel) {
	fmt.Println("------------------------------")
	fmt.Println("ID:", task.ID)
	fmt.Println("Title:", task.Title)
	fmt.Println("Discussion:", task.Discussion)
	fmt.Println("Completed:", task.Completed)
	fmt.Println("Completed At:", task.CompletedAt)
	fmt.Println("Created At:", task.CreatedAt)
}
