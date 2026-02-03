package main

import (
	"fmt"
	"os"
)

func main() {
	val := os.Getenv("phone_number")
	if val != "" {
		fmt.Println("val:", val)
	} else {
		fmt.Println("Переменная val не задана!")
	}

	/*ctx := context.Background()

	conn, err := simple_connection.CreatConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)

	}

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 8 {
			task.Title = " Покормить кота"
			task.Discussion = "Кот очень голоден"
			task.Completed = true
			now := time.Now()
			task.CreatedAt = &now

			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}

			break
		}
	}

	fmt.Println("succeed!")*/

}
