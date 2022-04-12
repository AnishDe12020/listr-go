package main

import (
	"fmt"
	"github.com/AnishDe12020/listr-go"
)

func main() {

	l := listr.New(listr.Options{
		Tasks: []listr.Task{
			{
				Title: "Task 1",
				Run: func() error {
					fmt.Println("Task 1")
					return nil
				},
			},
		},
	})

	fmt.Printf(l.Tasks[0].Title)
}
