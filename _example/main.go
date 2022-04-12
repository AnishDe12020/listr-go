package main

import (
	"fmt"
	"os/exec"

	"github.com/AnishDe12020/listr-go"
)

func main() {

	l := listr.New(listr.Options{
		Tasks: []listr.Task{
			{
				Title: "Task 1",
				Run: func() error {
					cmd := exec.Command("echo", "-l")
					fmt.Println("Task 1")
					out, err := cmd.Output()
					fmt.Printf("%s\n", out)

					return err
				},
			},
		},
	})

	l.Run()

}
