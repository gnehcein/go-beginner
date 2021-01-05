package main

import (
	"strings"
	"sync"
	"time"
	"fmt"
)

type (
	subscriber chan interface{}
	topicFunc func(v interface{}) bool
)
type Publisher struct {
	//m  互斥锁。
	//buffer  管道长度，即可接受最多的订阅主题。
	//timeout 最多等待时间。
	//subscriber 订阅者的消息
	m			  sync.RWMutex
	buffer		  int
	timeout		  time.Duration
	subscribers   map[subscriber]topicFunc
}


// 构建⼀个发布者对象, 可以设置发布超时时间和缓存队列的⻓度
func NewPublisher(publishtime time.Duration,buffer int)  *Publisher{
	return &Publisher{
		buffer:			buffer,
		timeout: 		publishtime,
		subscribers:  	make(map[subscriber]topicFunc),
	}
}

//添加一个新的订阅者，订阅通过过滤器（即topicFunc）筛选的主题（v）
func (p *Publisher) SubscribeTopic(topic topicFunc)chan interface{}{
	ch:=make(chan interface{},p.buffer)
	p.m.Lock()	//锁起来防止变化
	p.subscribers[ch]=topic	 //map插入
	p.m.Unlock()
	return ch
}

// 添加⼀个新的订阅者，订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

//退出订阅
func (p *Publisher) Evict(sub chan interface{}){
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers,sub)
	close(sub)
}

//发布一个主题
func (p *Publisher) Publish(v interface{}){
	p.m.RLock()
	defer p.m.RUnlock()
	var wg sync.WaitGroup
	for sub,topic:=range p.subscribers{
		wg.Add(1)
		go p.sendTopic(sub,topic,v,&wg)
	}
	wg.Wait()
}

//发布主题，可以容忍一定的超时
func (p *Publisher)sendTopic(sub subscriber,topic topicFunc,v interface{},wg *sync.WaitGroup){
	defer wg.Done()
	if topic!=nil&&!topic(v){
		return
	}
	select {
	case sub<-v:
	case <-time.After(p.timeout):
		fmt.Println("没有发布成功一次")
	}
}

// 关闭发布者对象，同时关闭所有的订阅者管道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub:=range p.subscribers{
		delete(p.subscribers,sub)
	}
}



func main(){
	p:=NewPublisher(200*time.Millisecond,10)  //制造publisher
	defer p.Close()
	all:=p.Subscribe()
	nie:=p.SubscribeTopic(func(v interface{})bool{ //匿名函数
		s,ok:=v.(string)
		if ok&&strings.Contains(s,"nie"){
			return true
		}
		return false
	})

	p.Publish("wljer")
	p.Publish("wfwf")
	p.Publish("niech")

	go func() {  	//使用协程
		for v:=range all{
			fmt.Println(v)
		}
	}()

	go func(){
		for v:=range nie{
			fmt.Println(v)
		}
	}()

	time.Sleep(5*time.Second)
}






