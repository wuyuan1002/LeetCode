
#include <condition_variable>
#include <iostream>
#include <mutex>
#include <thread>
#include <vector>

// n个线程交替打印1-n, 到最大值max后退出
//
// 实现原理:
// 每个线程有一个与之关联的条件变量，它会等待其条件变量被通知并在该由当前线程打印时打印当前数字。
// 然后，它会通知下一个线程，可以打印下一个数字了。
//
// 注意：这种方法需要确保每个线程在条件变量上的等待顺序与它们的通知顺序匹配。
// 因此，主要确保了在开始第一个线程之前，所有的线程都已经开始运行并在它们各自的条件变量上等待。
// 然后再通知首个线程开始工作，启动串行执行。

// -------

// 临界资源 -- i、max
int i = 1, max = 100;
int totalThreads = 5;

// 互斥锁 -- 用于保护临界资源, 确保临界资源同时只能有一个线程可访问,
// 在这里就是递增的n, 在生产者消费者就是生产消费所使用的队列（临界资源可能同时被多个线程访问，所以需要加锁保护，比如生产者通知消费者开始消费，而消费者条件队列上等待着很多消费者）
std::mutex mtx;

// 条件变量 -- 用于让交替打印的n个线程休眠在各自的条件变量上, 有几个线程就有几个条件变量
// 生产者消费者的变种, 这里相当于有一个生产者和一个消费者, 并且有其余类型如分解者, (生产->消费->分解->生产->...)
std::vector<std::condition_variable> cvs(5);

void printNumbers(int id) {
    while (true) {
        // 加锁确保对临界资源的访问是线程安全的
        std::unique_lock<std::mutex> lock(mtx);

        // 所有数字已经打印完毕
        if (i > max) {
            // 通知下一个线程可以开始工作了
            cvs[(id + 1) % totalThreads].notify_one();
            break;
        }

        if ((i - 1) % totalThreads != id) {
            // 如果当前的数字不应由此线程打印, 则等待 -- wait后自动释放lock, 被唤醒后自动获取lock
            cvs[id].wait(lock);
        } else {
            // 打印数字
            std::cout << "Thread " << id << ": " << i++ << std::endl;
            // 通知下一个线程可以开始工作了
            cvs[(id + 1) % totalThreads].notify_one();
        }
    }
}

int main() {
    std::vector<std::thread> threads;
    for (int i = 0; i < totalThreads; ++i) {
        threads.push_back(std::thread(printNumbers, i));
    }

    // 第一个线程开始工作
    cvs[0].notify_one();

    // 等待所有线程结束
    for (auto& t : threads) {
        t.join();
    }

    return 0;
}
