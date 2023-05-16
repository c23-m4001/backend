package loader

import (
	"time"

	"github.com/graph-gophers/dataloader"
)

func newDataloader(batchFn dataloader.BatchFunc) dataloader.Loader {
	return *dataloader.NewBatchedLoader(
		batchFn,
		dataloader.WithWait(10*time.Microsecond),
		dataloader.WithBatchCapacity(50),
	)
}
