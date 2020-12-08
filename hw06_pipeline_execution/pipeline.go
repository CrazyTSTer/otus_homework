package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, StageFunc := range stages {
		in = StageFunc(StreamFinisher(done, in))
	}
	return in
}

func StreamFinisher(done In, in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case data, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- data:
				case <-done:
					return
				}
			case <-done:
				return
			}
		}
	}()
	return out
}
