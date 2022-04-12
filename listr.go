package listr

type Listr struct {
	Tasks []Task
}

type Task struct {
	Title   string
	Enabled bool
	Skip    SkipFunction
	Run     RunFunction
}

type RunFunction func() (string error)
type SkipFunction func() bool

type Options struct {
	Tasks []Task
}

func New(options Options) *Listr {
	return &Listr{
		Tasks: options.Tasks,
	}
}
