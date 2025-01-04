package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, i int, wg *sync.WaitGroup) {
	defer wg.Done() // タスク完了を通知
	ch <- i * 2     // チャネルに値を送信
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch { // チャネルから値を受信
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
	fmt.Println("###############")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// producerのゴルーチンを起動
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i, &wg)
	}

	// producerが終了するのを待ち、チャネルを閉じる
	go func() {
		wg.Wait() // producer の完了を待つ
		close(ch) // チャネルを閉じる
	}()

	// consumerのゴルーチンを起動
	wg.Add(10) // 10個のタスクがconsumerで処理される
	go consumer(ch, &wg)

	// すべてのタスクが終了するのを待機
	wg.Wait()
	fmt.Println("Done")
}
