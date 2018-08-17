package controller

import (
	"sync"
	"fmt"

	"app/panicroutinepro/service"
	"app/panicroutinepro/common"
)

func SignUp(username string) {
	defer common.Catch(nil, func(err interface{}) {
		fmt.Printf("[panicroutinepro] SignUp: %s \n", err.(error).Error())
	})
	performMustSignUp(username)
	fmt.Printf("[panicroutinepro] SignUp: %s \n", "ok")
}

func performMustSignUp(username string) {
	ch := make(chan error, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go service.MustSignUp(username, ch, wg)

	wg.Wait()
	close(ch)

	if err := <-ch; err != nil {
		panic(err)
	}
}