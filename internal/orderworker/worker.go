package orderworker

import "context"


type Worker struct {
	queue chan OrderCommand
}

func NewWorker(buffer int) *Worker {
	return &Worker{
		queue: make(chan OrderCommand, buffer),
	}
}



func (w *Worker) Run(
	ctx context.Context,
	workers int,
	process func(ctx context.Context, cmd OrderCommand) error,
) {
	for i := 0; i < workers; i++ {
		go func(id int) {
			for {
				select {
				case <-ctx.Done():
					return
				case cmd := <-w.queue:
					err := process(ctx, cmd)
					cmd.Reply <- err
				}
			}
		}(i)
	}
}

func (w *Worker) Enqueue(cmd OrderCommand) {
	w.queue <- cmd
}