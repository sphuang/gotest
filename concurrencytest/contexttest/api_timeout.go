package contexttest

import (
	"fmt"
	"time"
	"context"
)

func launchAPI(ctx context.Context) {
	count := 0

	for {
		select {
		case <- ctx.Done():
			fmt.Println("Time out!")
			return
		default:
			time.Sleep(time.Second)

			count++
			if count > 5 {
				fmt.Println("Finished")
				return
			}
		}
	}
}

func LaunchAPI(expired time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), expired)
	defer cancel()

	launchAPI(ctx)
}