package process

type Process interface {
	Start()
	Stop()
	Chan() chan interface{}
}

type BaseProcess struct {
	Channel chan interface{}
}

func (p *BaseProcess) Chan() chan interface{} {
	return p.Channel
}
