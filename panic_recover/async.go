package panic_recover

import (
	"fmt"
	"runtime"
)

func AsyncWithRecover(runnable func() error, recoverFunc func(rec StackError)) {
	go func() {
		var err error

		defer func() {
			if rec := recover(); rec != nil {
				buf := make([]byte, 1024)
				var isErr bool

				if err, isErr = rec.(error); !isErr {
					err = fmt.Errorf("%v", rec)
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
