package pipeline

type Stage interface {
	Handle(any) any
}

type StageFunc func(in any) any

func (f StageFunc) Handle(in any) any {
	return f(in)
}

type Pipeline struct {
	out chan any
}

func New(a ...any) *Pipeline {
	out := make(chan any)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return &Pipeline{out: out}
}

func (p *Pipeline) Through(stages ...Stage) *Pipeline {
	in := p.out
	for _, s := range stages {
		s := s
		out := make(chan any)
		go func(in <-chan any) {
			for v := range in {
				out <- s.Handle(v)
			}
			close(out)
		}(in)
		in = out
	}
	p.out = in
	return p
}

func (p *Pipeline) Out() <-chan any {
	return p.out
}
