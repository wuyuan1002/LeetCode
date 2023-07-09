

#include <condition_variable>
#include <iostream>
#include <mutex>
#include <thread>
#include <vector>

// n个线程交替打印1-n, 到最大值max后退出

class PrintN {
public:
    void print(int n, int max) {
        // 临界资源 -- i、max
        int i = 0;

        // 互斥锁 -- 用于保护临界资源, 确保临界资源同时只能有一个线程可访问,
        // 在这里就是递增的n, 在生产者消费者就是生产消费所使用的队列
        std::mutex mutex;

        // 条件变量 -- 用于让交替打印的n个线程休眠在各自的条件变量上, 有几个线程就有几个条件变量
        // 生产者消费者的变种, 这里相当于只有一个生产者和一个消费者, 并且有其余类型如分解者, (生产->消费->分解->生产->...)
        std::vector<std::condition_variable> cv = std::vector<std::condition_variable>(n);

        // 创建n个线程
        for (int idx = 0; idx < n; idx++) {
            std::thread worker([&]() {
                cv[idx < n - 1 ? idx + 1 : 0].wait(mutex);
                for (;;) {
                    std::unique_lock<std::mutex> lock(mutex);

                    // 打印
                    std::cout << std::this_thread::get_id() << ": " << ++i << std::endl;

                    // 唤醒下一个线程打印
                    cv[idx < n ? idx + 1 : 0].notify_all();

                    // 若达到最大值则退出
                    if (i >= max) {
                        break;
                    }
                }
            });
            worker.join();
        }
    }
};

int main() {
    // PrintN().print(3, 20);
    std::cout << std::this_thread::get_id() << ": " << std::endl;
    return 0;
}