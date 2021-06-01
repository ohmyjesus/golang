package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// 1. 模拟3节点raft选举
// 2. 改造代码成分布式选举代码，加入RPC调用
// 3. 演示一下完整代码 包含自动选主和日志复制

// 定义3节点常量
const raftCount = 3

// 声明一个leader对象
type Leader struct {
	// 任期
	Term int
	// 领导者编号
	LeaderId int
}

// 声明raft
type Raft struct {
	// 锁
	mu sync.Mutex
	// 节点编号
	me int
	// 当前任期
	currentTerm int
	// 为哪个节点投票
	votedFor int
	// 状态 0-follower 1-candidate 2-leader
	state int
	// 发送最后一条数据的时间
	lastMessageTime int64
	// 设置当前节点的leader
	currentLeader int
	// 节点发信息的通道
	message chan bool
	// 选举的通道
	electCh chan bool
	// 心跳信号通道
	heartBeat chan bool
	// 返回心跳信号的通道
	heartbeatRe chan bool
	// 超时时间 150-300ms
	timeout int64
}

// 0还没上任， -1还没有编号
var leader = Leader{
	Term:     0,
	LeaderId: -1,
}

func main() {
	// 过程：有3个节点，最初都是follower
	// 若有candidate状态，进行投票拉票
	// 会产生leader

	// 创建3个节点
	for i := 0; i < raftCount; i++ {
		// 创建3个raft节点
		Make(i)
	}

	select {}
}

// 创建节点
func Make(me int) *Raft {
	rf := &Raft{}
	rf.me = me
	rf.votedFor = -1 // -1 代表谁都不投，此时节点刚创建
	rf.state = 0     // 0 follower
	rf.timeout = 0
	rf.currentLeader = -1
	rf.setTerm(0)
	// 初始化通道
	rf.message = make(chan bool)
	rf.electCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartbeatRe = make(chan bool)
	// 设置随机种子
	rand.Seed(time.Now().UnixNano()) // 时间戳

	// 选举的协程
	go rf.Election()

	// 心跳检查的协程
	go rf.sendLeaderHeartBeat()

	return rf
}

// 设置任期
func (rf *Raft) setTerm(term int) {
	rf.currentTerm = term
}

// 选举
func (rf *Raft) Election() {
	// 设置一个标记，是否选出了leader
	var flag bool = false
	for {
		// 设置超时，150-300的随机数
		rf.timeout = randRange(150, 300)
		rf.lastMessageTime = millisecond()
		select {
		// 延迟等待一毫秒
		case <-time.After(time.Duration(rf.timeout) * time.Millisecond):
			fmt.Println("当前节点的状态为", rf.state)
		}
		for !flag {
			// 选主
			flag = rf.ElectionOneRound(&leader)
		}
	}
}

// 随机值方法
func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// 获取当前时间，发送最后一条数据的时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 选主逻辑
func (rf *Raft) ElectionOneRound(leader *Leader) bool {
	// 定义超时
	var timeout int64 = 100
	// 投票数量
	var vote int
	// 是否开始心跳信号的产生
	var triggerHeartBeat bool = false
	// 时间
	last := millisecond()
	// 用于返回值
	var success bool = false

	// 给当前节点变成candidate
	rf.mu.Lock()
	rf.BecomeCandidate()
	rf.mu.Unlock()
	fmt.Println("start election leader")
	for {
		// 遍历所有节点拉选票
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				// 拉选票
				go func() {
					if leader.LeaderId < 0 {
						// 设置投票
						rf.electCh <- true
					}
				}()
			}
		}
		// 设置投票数量
		vote = 1
		// 遍历节点加选票
		for i := 0; i < raftCount; i++ {
			// 计算投票的数量
			select {
			case ok := <-rf.electCh:
				if ok {
					// 如果ok为true，则数量加1
					vote++
					// 如果选票个数大于节点个数的一半，则当选为leader
					success = vote > raftCount/2
					if success && !triggerHeartBeat {
						// 变化成主节点 并触发心跳信号检测
						rf.mu.Lock()
						rf.BecomeLeader()
						rf.mu.Unlock()
						triggerHeartBeat = true
						// 由leader向其他节点发送心跳信号
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成为了leader")
						fmt.Println("leader开始发送心跳信号了")
					}
				}
			}
		}
		// 做最后检验工作
		// 若不超时，且票数大于一半，则选举成功
		if timeout+last < millisecond() || (vote > raftCount/2) || rf.currentLeader > -1 {
			break
		} else {
			// 等待操作
			select {
			case <-time.After(10 * time.Millisecond):

			}
		}
	}
	return success
}

// 修改状态candidate
func (rf *Raft) BecomeCandidate() {
	rf.state = 1
	rf.setTerm(rf.currentTerm + 1)
	rf.votedFor = rf.me
	rf.currentLeader = -1
}

// 修改状态leader
func (rf *Raft) BecomeLeader() {
	rf.state = 2
	rf.currentLeader = rf.me
}

// leader发送心跳信号
// 顺便完成数据同步，这里先不实现了
// 看小弟挂没挂
func (rf *Raft) sendLeaderHeartBeat() {
	// 死循环
	for {
		select {
		case <-rf.heartBeat:
			rf.sendAppendRntry()
		}
	}
}

// 用于返回给leader的确认信号
func (rf *Raft) sendAppendRntry() {
	// 如果是主节点，就不执行
	if rf.currentLeader == rf.me {
		// 记录确认信号的结点个数
		var successCount = 0
		// 设置确认信号
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					rf.heartbeatRe <- true
				}()
			}
		}
		// 计算返回确认信号的个数，校验
		for i := 0; i < raftCount; i++ {
			select {
			case ok := <-rf.heartbeatRe:
				if ok {
					successCount++
					if successCount > raftCount/2 {
						fmt.Println("投票选举成功， 心跳检测ok")
						log.Fatal("程序结束")
					}
				}
			}
		}
	}
}
