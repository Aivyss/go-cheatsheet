package pool

import (
	"context"
	"go-cheatsheet/rec"
	"sync"
)

type AsyncPool interface {
	SubmitJob(job func(synchronizer Synchronizer))
}

type asyncPool struct {
	ctx       context.Context
	batchSize int
	ch        chan func(synchronizer Synchronizer)
	Synchronizer
}

type AsyncPoolOption struct {
	BatchSize   *int                           // optional (default: 5)
	RecoverFunc func(rec rec.StackError) error // optional
}

func (a *asyncPool) SubmitJob(job func(synchronizer Synchronizer)) {
	a.ch <- job
}

func NewAsyncPool(ctx context.Context, setOptions ...func(option *AsyncPoolOption)) AsyncPool {
	batchSize := 5
	recoverFunc := rec.DefaultSyncRecoverFunc
	obj := &asyncPool{
		ctx:          ctx,
		Synchronizer: NewSynchronizer(),
		ch:           make(chan func(synchronizer Synchronizer)),
	}

	if len(setOptions) > 0 {
		option := new(AsyncPoolOption)
		for _, setOption := range setOptions {
			setOption(option)
		}

		if option.BatchSize != nil {
			batchSize = *option.BatchSize
		}
		if option.RecoverFunc != nil {
			recoverFunc = option.RecoverFunc
		}
	}
	obj.batchSize = batchSize

	for i := 0; i < obj.batchSize; i++ {
		go func() {
			for {
				select {
				case <-obj.ctx.Done():
					return
				case runnable, ok := <-obj.ch:
					if !ok {
						return
					}

					_ = rec.SyncWithRecover(func() error {
						runnable(obj.Synchronizer)

						return nil
					}, recoverFunc)
				}
			}
		}()
	}

	return obj
}

type Synchronizer interface {
	Sync(runnable func())
}

type synchronizer struct {
	*sync.Mutex
}

func (s *synchronizer) Sync(runnable func()) {
	s.Lock()
	runnable()
	s.Unlock()
}

func NewSynchronizer() Synchronizer {
	var mutex sync.Mutex

	return &synchronizer{
		Mutex: &mutex,
	}
}
