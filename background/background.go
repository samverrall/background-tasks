package background

import "log"

type JobFunc func()

func Go(fn JobFunc) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()

		fn()
	}()
}
