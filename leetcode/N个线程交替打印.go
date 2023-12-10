package main

import (
	"fmt"
	"sync"
)

// N个线程交替打印 1~max, 到最大值max后退出

// printAlternately .
// 每个线程有一个与之关联的条件变量，它会等待其条件变量被通知并在该由当前线程打印时打印当前数字。
// 然后，它会通知下一个线程，可以打印下一个数字了。
//
// 这种方法需要确保每个线程在条件变量上的等待顺序与它们的通知顺序匹配。
// 因此，主要确保了在开始第一个线程之前，所有的线程都已经开始运行并在它们各自的条件变量上等待。
// 然后再通知首个线程开始工作，启动串行执行。
//
// threadCount：线程数量
// maxNum：输出的最大值
func printAlternately(threadCount int, maxNum int) {
	num := 0               // 打印的数字
	wg := sync.WaitGroup{} // 用于主协程等待所有协程结束

	// 为每个协程构造自己的条件变量，让每个协程在各自的条件变量上等待
	readyChans := make([]chan bool, threadCount)
	for i := range readyChans {
		// 这里每个协程的条件变量必须是有缓冲区的channel，保证发送信号的协程不会因为channel上没有阻塞的协程而导致发送信号的协程也被阻塞。
		// 若channel没有缓冲区，会导致通知协程结束时，最后一个结束的协程向下一个协程的channel发送信号，
		// 但是下一个channel上已经没有协程阻塞在上面了，会导致最后一个协程在发送信号时阻塞，进而无法执行wg.Done()，
		// 此时主协程和最后一个打印协程都阻塞了，出现协程泄漏deadlock
		readyChans[i] = make(chan bool, 1)
	}

	// 构造threadCount个协程，每个协程开始执行打印函数，默认都会先阻塞等待允许该协程打印的信号
	for i := 0; i < threadCount; i++ {
		// wg+1 -- 主协程等待所有协程结束后退出
		wg.Add(1)

		// 当前协程开始执行打印函数
		go func(gid int) {
			for {
				// 等待允许当前协程打印的信号
				<-readyChans[gid]

				// 获取当前协程要打印的数字
				num += 1

				// 若要打印的数字已经大于最大值
				if num > maxNum {
					readyChans[(gid+1)%threadCount] <- true // 通知下一个协程退出
					break                                   // 退出当前协程
				}

				// 打印数字
				fmt.Printf("goroutine %d：%d\n", gid, num)

				// 通知下一个协程开始打印
				readyChans[(gid+1)%threadCount] <- true
			}

			// 标记当前协程执行结束
			wg.Done()
		}(i)
	}

	// 通知第一个协程开始打印
	readyChans[0] <- true

	// 等待所有goroutine完成
	wg.Wait()
}

func main() {
	printAlternately(3, 5)
}
