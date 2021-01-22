package routinetest

import (
	"fmt"
	"sync"
)

type Account struct {
	Btc int
	lock sync.Mutex
}

func (a *Account) PlusFund(f int) {
	a.Btc += f
}

func (a *Account) PlusFundWithLock(f int) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.Btc += f
}

func ConcurrentPlusFund(lock bool) int {
	var wg sync.WaitGroup
	nWorker := 10
	wg.Add(nWorker)

	a := Account{ Btc: 0.0 }

	for i := 0; i < nWorker; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				if lock {
					a.PlusFundWithLock(1)
				} else {
					a.PlusFund(1)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	
	fmt.Println(a.Btc)
	return a.Btc
}