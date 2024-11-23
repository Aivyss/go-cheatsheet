package rec

import (
	"errors"
	"fmt"
	"runtime"
)

var NoFunctionalObjectErr = errors.New("no functional object")

type AsyncRecoverFunc func(rec StackError)

func AsyncWithRecover(runnable func() error, recoverFunc AsyncRecoverFunc) {
	if runnable == nil {
		panic(NoFunctionalObjectErr)
	}

	go func() {
		var err error

		defer func() {
			if rec := recover(); rec != nil {
				buf := make([]byte, 1024)
				var isErr bool

				if err, isErr = rec.(error); !isErr {
					err = fmt.Errorf("%v", rec)
				}

				if recoverFunc == nil {
					recoverFunc = DefaultAsyncRecoverFunc
				}
				recoverFunc(&stackErrorImpl{
					error: err,
					stack: buf[:runtime.Stack(buf, true)],
				})
			}
		}()

		err = runnable()
	}()
}

func DefaultAsyncRecoverFunc(rec StackError) {}
