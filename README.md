# Go-Golang
&emsp;&emsp;Go语言作为21世纪的语言，其天生就支持并发，能够充分利用多核处理器，并且Go语言实现了开发效率和执行效率的完美结合，其具有C/C++的性能同时兼备python代码的效率。

俗话说，技多不压身，今天开始每天打卡，记录学习内容，慢慢地熟悉Go语言，争取转型Go

此学习笔记来源于[Go语言学习视频](https://www.bilibili.com/video/BV16E411H7og?p=7&spm_id_from=pageDriver)，库中目前包含的有以下知识

### 第一部分 **基本语法**
#### 今日学习
&emsp;&emsp;变量、常量、itoa、基本数据类型、流程控制语句，运算符    

&emsp;&emsp;( 近期一直在学c++，突然学其他语法还有点不太习惯，还是需要多联想和多记一下他们之间的区别 )

<p align="right">2021.5.10 night</p>

#### 今日学习
&emsp;&emsp;数组、切片、切片的相关操作、指针、map、函数、defer关键字、全局变量、函数作为形参和返回值类型、匿名函数、闭包。

&emsp;&emsp;( 切片相当于一个框，框住的是底层数组的一部分，是数组的引用类型，切片是动态拓展的，而数组是静态空间的 )

<p align="right">2021.5.11 night</p>

#### 今日学习
&emsp;&emsp;panic崩溃、fmt包、函数的递归、type关键字、struct结构体、构造函数、方法(特殊的函数)、匿名结构体、结构体嵌套、结构体“继承”、结构体和json。

&emsp;&emsp;( 理解结构体的使用，结构体和方法的联用，最后结合以上知识点做了一个作业1--函数版的学生管理系统homework1_stu_manage)

<p align="right">2021.5.12 night</p>

#### 今日学习
&emsp;&emsp;接口、指针接收者实现接口、同一个结构体实现多个接口、空接口、类型断言、包、包的导入、文件操作、使用bufio获取用户输入、在文件中插入指定内容(用到临时文件)、time包的使用。

&emsp;&emsp;( 掌握接口的使用，文件操作等，结合以上知识做了一个作业2--向文件中插入内容homework2_file_insert)

<p align="right">2021.5.13 night</p>

#### 今日学习
&emsp;&emsp;log库、runtime获取时间信息、reflect反射。

&emsp;&emsp;( 今天主要是结合之前学习的内容做了一个日志库，详情见作业3--我的日志库homework3_mylogger)

<p align="right">2021.5.14 night</p>

#### 周末学习
&emsp;&emsp;周末两天主要是修改论文，然后出去走了走，就没看go的新知识。 下周继续学习go

<p align="right">2021.5.15 & 2021.5.16 night</p>

#### 今日学习
&emsp;&emsp;ini配置文件解析，类似于手动实现unmarshal功能，整个流程实现起来较为复杂，主要是有很多函数都不太熟，比如TypeOf和ValueOf函数，它们对应的有哪些方法也不是很清楚，所以动手实现起来比较慢，不过好在最后还是完成了个简单版本的ini配置文件解析，也还行，明天开始看并发和channel。(作业4--ini配置文件解析homework4_ini)

<p align="right">2021.5.17 night</p>

### 第二部分 **并发编程**
#### 今日学习
&emsp;&emsp;strconv包(字符串数字转整型数字等)、goroutine创建线程、waitGroup等待goroutine执行结束、GOMAXPROCS决定跑满CPU核心的个数、channel创建通道、有无缓冲区的通道和单向通道、通道close时的注意事项。

<p align="right">2021.5.18 night</p>

#### 今日学习
&emsp;&emsp;今天把论文返修回去了，其余时间看了一下goroutine池，主要是创建多个goroutine它们之间通过通道进行数据交互，最后完成了一个小作业homework5_goroutine_channel(利用goroutine和channel实现一个计算int64随机数各位数和的程序)

<p align="right">2021.5.19 night</p>

#### 今日学习
&emsp;&emsp;互斥锁、读写互斥锁、map的多线程安全版本sync.map、atomic包、tcp的客户端服务器端通信、tcp粘包时的处理方式、udp的客户端服务器端通信、http的服务器端编写。(感觉go语言写socker通信好方便，比在linux中容易实现多了，很快就能上手。 最后做了一个小作业homework6_tcp实现了客户端和服务器端的tcp通信)

<p align="right">2021.5.20 night</p>

#### 今日学习
&emsp;&emsp;http的客户端、服务器端程序、条件测试、性能基准测试。 http的客户端服务端可实现局域网内主机的彼此通信。 条件测试分为测试组和子测试，即测试一组案例或测试单个案例。性能基准测试为查看程序执行的次数，时间，申请内存次数和占用CPU内核数。

<p align="right">2021.5.21 night</p>

#### 周末学习
&emsp;&emsp;pprof性能调优(查看CPU和内存的profile性能信息)、os.Args获取命令行参数、用flag包获取更加详细的命令行参数。

<p align="right">2021.5.22 & 2021.5.23 night</p>
