# Go-Golang
&emsp;&emsp;Go语言作为21世纪的语言，很契合当今时代的发展，支持并发(充分地利用多核处理器)，能够做网络编程(TCP/HTTP的客户端服务端交互)，能连接数据库完成数据交互等很多功能。Go语言它实现了开发效率和执行效率的完美结合，其具有C/C++的性能同时兼备python代码的效率。

今天开始每天打卡，记录学习内容，慢慢地熟悉Go语言，争取转型Go

此学习笔记来源于[Go语言学习视频](https://www.bilibili.com/video/BV16E411H7og?p=7&spm_id_from=pageDriver)，库中目前包含的有以下知识

## 第一部分 基本语法
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

### 本周小结：
&emsp;&emsp;这一周主要学习了Go语言的基本语法，语法基础必须要牢牢掌握。

## 第二部分 网络编程

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
&emsp;&emsp;pprof性能调优(查看CPU和内存的profile性能信息)、os.Args获取命令行参数、用flag包获取更加详细的命令行参数、使用Go语言连接数据库并实现数据库的增删改查操作，和数据库的预处理操作。

<p align="right">2021.5.22 & 2021.5.23 night</p>

### 本周小结：
&emsp;&emsp;Go语言能做网络编程已经很让我惊讶了，没想到Go语言还能连接数据库完成相关的操作，真是amazing，现在好期待go语言后面又会给我带来什么惊喜。 本周重点内容是：创建goroutine线程、channel通道、互斥锁读写锁、socket编程(TCPUDP客户端服务器端通信)、HTTP客户端服务器端通信、连接mysql数据库完成增删改查操作。

#### 今日学习
&emsp;&emsp;mysql事务的实现、sql的提高版sqlx，利用sqlx实现查询和插入数据、sql注入、消息队列nsq初识(go操作nsq实现消费者和生产者模型)、context实现线程的退出、agent服务端开发项目的初步了解(kafka和zookeeper的原理，后续会有tailf和etcd等知识点)

<p align="right">2021.5.24 night</p>

#### 今日学习
&emsp;&emsp;利用tailf读日志和利用sarama向kafka里写日志，这几天需要完成一个logAgent服务端开发项目。

<p align="right">2021.5.25 night</p>

#### 今日学习
&emsp;&emsp;zookeeper版本的logagent小作业已经完成，详情见homework7_logagent_zookeeper。其中用到了一个第三方的ini解析库，tailf对ini解析后的日志地址进行读取，将读取到的日志信息利用通道传输，然后sarama将该日志信息发送给kafka。此外还学习了etcd的搭建和操作，包括put、get、del和watch操作，现在是在利用etcd对zookeeper版本的日志收集项目进行优化。

<p align="right">2021.5.26 night</p>

#### 今日学习
&emsp;&emsp;老师的etcd版本用到的东西很多，对现在的我来说有点复杂，一下有点接受不过来，于是自己就先做了个低配版的etcd_log。主要思路是先初始化kafka和etcd，连接上etcd之后不再从配置文件里读，而是从etcd里面获取日志信息，然后创建子goroutine对日志信息进行watch监视并将日志发送到kafka，遗憾的是未能实现日志的动态新增功能，后续会对老师的版本进行理解和掌握。详情见homework8_logagent_etcd。

<p align="right">2021.5.27 night</p>

#### 今日学习
&emsp;&emsp;elasticsearch的介绍和环境搭建、go操作ES实现插入数据、利用POSTMAN向ES中插入数据和检索和用Kibana查询ES中的数据键值对。

<p align="right">2021.5.28 night</p>

#### 周末学习
&emsp;&emsp;logtransfer的实现，主要思路是kafka消费者从kafka中读出数据，然后将数据发往ES中，发送成功后可以通过POSTMAN进行索引查询或通过Kibana进行可视化的数据查看。到这里的话日志收集项目差不多就完成了，有tailf版本的，有etcd版本的，然后再通过ES搜索引擎进行优化的，整个日志收集项目用到的东西还是蛮多的，难度其实不算大，主要是思路一定要清晰，做起来就会容易很多，后续会再复习代码进行深度理解和掌握。

<p align="right">2021.5.29 night</p>

## 第三部分 Web框架
#### 周末学习
&emsp;&emsp;用prometheus和grafana来监控系统信息和查看性能指标等。go语言web框架的初识gin，用gin获取API参数和URL参数，获取表单参数，上传单个文件和上传多个文件。

<p align="right">2021.5.30 night</p>

#### 今日学习
&emsp;&emsp;gin数据解析和绑定；其中包括json数据解析和绑定、表单数据解析和绑定和URI数据解析和绑定， gin数据渲染，HTML模式渲染，重定向，异步和同步，中间件的了解和练习，cookie状态信息的了解和练习。晚上在做一个gin版本的数据库增删改查操作，其中用到了gin框架，连接mysql及其相关操作，HTML等知识，预计明天完成。

<p align="right">2021.5.31 night</p>

#### 今日学习
&emsp;&emsp;gin版本的mysql作业已完成，详情见homework9_database_gin。另外学习了一下etcd的raft算法的底层原理，看了一下raft的动画，主要是自动选举和日志复制，虽然跟着老师走了一遍，但还是有点迷糊，后续还是需要再重温加深一下印象，接着学习了选项设计模式、gRPC的了解和gRPC版的客户端和服务端的通信，感觉和直接用go实现的tcp通信差不多，明天学习goMicro和docker，这一周应该能结束go语言的基础知识。

<p align="right">2021.6.1 night</p>

#### 今日学习
&emsp;&emsp;早上看一下gomicro，对它的概念有一个初步的印象，但还没有真正实践它。下午学习了docker，没在虚拟机上跑，自己买了个阿里云的轻量服务器，在服务器上跑docker，同时今天也首次注册了dockerhub，拥有了属于自己的一个小库，跟着老师写了几个简单的docker image，也push到了dockerhub中，感觉docker还是很有趣的，功能很强大而且和github也相通。然后心里又在想，既然都有云服务器了，何不买个域名自制个网页让所有人都看到，于是晚上买了个域名，马上安排！ 域名今天还未处理好，因为还要备案什么的，不过会很快的。后面计划着要完成好多好多事情，一步一个脚印，keep studying~

<p align="right">2021.6.2 night 22:40 </p>

#### 今日学习
&emsp;&emsp;妈的github又炸了，晚上硬是没提交上，断了我的打卡，干。  今天主要学习了docker和集群的概念，感觉没有应用场景的话有种无从下手的感觉，到今天为止，就结束了go语言240集的视频学习，后续的话就是总结知识点查漏补缺，回顾一下之前做的小项目，搭建自己的博客网站，同时再刷刷算法题就差不多了，期待一个好的结果。

<p align="right">2021.6.3 night 00:01 </p>

## 第四部分 总结与完善
#### 周末学习
&emsp;&emsp;周末两天回顾了一下之前学的知识点，同时也学习了一下博客的搭建，数据库有4个，而且表的数据比较多，在写后台的结构体时，很容易搞混淆，不知道该用哪个结构体，不过现在有点思路了，就是要搞清楚每个页面要展示哪些数据，也就对应着哪一个结构体，以这个作为线索来顺着写就容易很多了。

<p align="right">2021.6.5 & 2021.6.6 night 00:01 </p>
