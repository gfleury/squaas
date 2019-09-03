package executors

type Executor interface {
	Init(interface{}) (interface{}, error)
	Run() (interface{}, error)
	SetData(interface{}) error
}
