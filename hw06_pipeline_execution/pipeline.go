package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		cancelableIn := getCancelableInChannel(in, done)
		in = stage(cancelableIn)
	}
	return in
}

func getCancelableInChannel(in In, done In) Bi {
	cancelableIn := make(Bi)
	go func(in In) {
		defer close(cancelableIn)
		for {
			select {
			case val, ok := <-in:
				if !ok {
					return
				}
				cancelableIn <- val
			case <-done:
				return
			}
		}
	}(in)
	return cancelableIn
}
