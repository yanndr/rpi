package process

type Process interface {
	Start()
	Stop()
	Channel() chan interface{}
}

type baseProcess struct {
	channel chan interface{}
}

func (p *baseProcess) Channel() chan interface{} {
	return p.channel
}
