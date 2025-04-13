package runtime

import "context"

type Executable = func(ctx context.Context, args ...interface{}) (interface{}, error)
