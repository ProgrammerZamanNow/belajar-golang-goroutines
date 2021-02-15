package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			/**
			Jika terjadi error : panic: sync: WaitGroup is reused before previous Wait has returned
			Itu artinya, goroutine belum selesai menjalankan kode group.Add(1), naun goroutine unit test
			sudah melakukan group.Wait(), group tidak boleh di add ketika sudah di Wait(), hal ini biasanya
			terjadi jika resource hardware kurang cepat ketika menjalankan goroutine diawal
			Jika hal ini terjadi, silahkan pindahkan kode group.Add(1), ke baris 15 sebelum memanggil go func()
			*/
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}
