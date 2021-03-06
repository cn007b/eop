package controller

import (
	"fmt"
	"sync"

	"app/examples/error_goroutine/service"
)

// SignUp - entrypoint for error_goroutine implementation.
func SignUp(username string) {
	msg := "ok"
	if err := performSignUp(username); err != nil {
		msg = err.Error()
	}
	fmt.Printf("[error_goroutine] SignUp: %s \n", msg)
}

func performSignUp(username string) error {
	ch := make(chan error, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go service.SignUp(username, ch, wg)

	wg.Wait()
	close(ch)

	return <-ch
}
