# hello-golang

## tutorial
1. [Go语言中文网](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter03/03.1.html)
## questions
1. 什么是引用类型, 与之并列的还有什么类型?
2. new与make
3. 对于结构体, 结构体指针和结构体本身都可以用s.val来调用属性, 这TM太奇怪了


## recommendations from The Way to Go
1. 对于容量很大的map, 最好在make的时候指定容量



## understanding
1. go之所以保留指针是因为想提高函数参数的传值效率
2. go的是强类型语言, 与python相比, 是好事, python的多态感觉有点让人不安
3. go面向对象的特点偏弱, 数组, 切片和map没有自己的方法, 只能用內建函数去操作
4. 函数(包括map的检索)一般都有两个返回值, 但是当你用一个值去接受时, 只会得到前面的值,后一个值代表操作是否成功
5. 数组在声明的时候就被初始化了, 而map在声明后还需要初始化


