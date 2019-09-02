package executors

type Executor interface {
	Init() error
	Run() (interface{}, error)
	SetData(interface{}) error
}
