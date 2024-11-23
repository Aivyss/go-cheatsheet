package rec

import (
	"fmt"
	"runtime"
)

func SyncWithRecover(runnable func() error, recoverFunc func(rec StackError) error) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			buf := make([]byte, 1024)
			var isErr bool

			if err, isErr = rec.(error); !isErr {
				err = fmt.Errorf("%v", rec)
			}

			err = recoverFunc(&stackErrorImpl{
				error: err,
				stack: buf[:runtime.Stack(buf, true)],
			})
		}
	}()

	err = runnable()

	return err
}

func DefaultSyncRecoverFunc(rec StackError) error {
	return rec
}
