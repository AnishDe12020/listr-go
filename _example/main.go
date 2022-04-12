package main

import (
	"errors"
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
					_, err := cmd.Output()
					// fmt.Printf("%s\n", out)

					err = errors.New("error")

					return err
				},
			},
		},
	})

	l.Run()

}
