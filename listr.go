package listr

import (
	"github.com/AnishDe12020/spintron"
)

type Listr struct {
	Tasks []Task
}

type Task struct {
	Title   string
	Enabled bool
	Skip    SkipFunction
	Run     RunFunction
}

type RunFunction func() error
type SkipFunction func() bool

type Options struct {
	Tasks []Task
}

func New(options Options) *Listr {
	return &Listr{
		Tasks: options.Tasks,
	}
}

func (l *Listr) Run() (string error) {
	for _, task := range l.Tasks {

		spintron := spinner.New(spinner.Options{
			Text: task.Title,
		})

		spintron.Start()

		if err := task.Run(); err != nil {
			spintron.Fail(err.Error())
			return err
		}

		spintron.Succeed(task.Title)
	}

	return nil
}
