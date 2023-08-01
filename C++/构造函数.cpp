
#include <iostream>

class A {
public:
    int x;

    A() {
        std::cout << "无参构造" << std::endl;
    }
    A(int x) : x(x) {
        std::cout << "有参构造" << std::endl;
    }
    A(A& a) : x(a.x) {
        std::cout << "拷贝构造" << std::endl;
    }
    A& operator=(A& a) {
        x = a.x;
        std::cout << "拷贝赋值构造" << std::endl;
        return *this;
    }
    A(A&& a) : x(a.x) {
        std::cout << "移动构造" << std::endl;
    }
    A& operator=(A&& a) {
        x = a.x;
        std::cout << "移动赋值构造" << std::endl;
        return *this;
    }
};

int main() {
    A a1; // 无参构造
    A a2(2); // 有参构造
    a1 = a2; // 拷贝赋值构造
    a1 = std::move(a2); // 移动赋值构造
    A a3(a1); // 拷贝构造
    A a4(std::move(a3)); // 移动构造

    std::cout << "--------" << std::endl;
    // 编译器优化 -- 若左侧对象还不存在, 则会直接调用对应的构造函数原地构造b1b2b3, 而不是先创建临时对象b1b2b3再调用他们的operator=进行赋值构造
    A b1 = A(3);  // 有参构造 -- 直接构造b1, 等效于 A b1(3); 的写法
    A b2 = b1;  // 拷贝构造 -- 直接把b1作为b2的构造函数入参
    A b3 = std::move(b2);  // 移动构造 -- 直接把b2作为b3的构造函数入参
    return 0;
}