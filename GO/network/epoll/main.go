//go:build linux
// +build linux

package main

// 在GO中处理epoll系统调用的示例
import (
	"flag"
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"runtime"
	"time"
)

const (
	MAX_EVENTS int = 128
	READ_QUEUE int = 0x18
	WRITE_QUEUE int = 0x19
)

type ReadyQueue struct {
	_type int
	event unix.EpollEvent
	readBuf []byte
	writeBuffer []byte
}

var threadFds = make([]chan unix.EpollEvent, nWorkerThread())

// 每个线程的读取队列
var threadReadQueue = make([][][]byte, MAX_EVENTS/ nWorkerThread())

// epoll fd
var epfd int

// addr
var (
	ip   = flag.String("ip", "0.0.0.0", "Ip地址")
	port = flag.Int("port", 9080, "端口地址")
)

func nWorkerThread() int {
	return runtime.NumCPU() - 1
}

// 初始化工作线程
func initSlaveThreads() {
	for i := 0; i < nWorkerThread(); i++ {
		threadFds[i] = make(chan unix.EpollEvent, MAX_EVENTS/ nWorkerThread())
	}
	for workId := 0; workId < nWorkerThread(); workId++ {
		go func(id int) {
			runtime.LockOSThread()
			// 私有读取队列
			readyReadOkQueue := make(map[int32]ReadyQueue,MAX_EVENTS/ nWorkerThread() / 2)
			readyWriteOkQueue := make(map[int32]ReadyQueue,MAX_EVENTS/ nWorkerThread() / 2)
			// 私有写入队列
			for {
				for key, value := range readyReadOkQueue {
					if !HandleRead(value) {
						continue
					} else {
						delete(readyReadOkQueue,key)
						readyWriteOkQueue[value.event.Fd] = value
					}
				}
				for key, value := range readyWriteOkQueue {
					if !HandleWrite(value) {
						continue
					} else {
						delete(readyWriteOkQueue,key)
						WriteCallback(value.event)
					}
				}
				select {
				case event := <-threadFds[id]:
					SlaveHandler(readyReadOkQueue,readyWriteOkQueue,event, HandleRead, HandleWrite,WriteCallback)
				}
			}
		}(workId)
	}
}

// read ok?
func HandleRead(queue ReadyQueue) bool {
	_, err := unix.Read(int(queue.event.Fd),queue.readBuf)
	if err != nil {
		return false
	}
	return true
}

func HandleWrite(queue ReadyQueue) bool {
	if len(queue.writeBuffer) > 0 {
		_, err := unix.Write(int(queue.event.Fd), queue.writeBuffer)
		if err != nil {
			return false
		}
		return true
	}
	queue.writeBuffer = append(queue.writeBuffer, "HTTP/1.1 200 OK\r\nServer: gnet\r\nContent-Type: text/plain\r\nDate: "...)
	queue.writeBuffer = time.Now().AppendFormat(queue.writeBuffer, "Mon, 02 Jan 2006 15:04:05 GMT")
	queue.writeBuffer = append(queue.writeBuffer, "\r\nContent-Length: 12\r\n\r\nHello World!"...)
	_, err := unix.Write(int(queue.event.Fd), queue.writeBuffer)
	if err != nil {
		return false
	}
	return true
}

func WriteCallback(event unix.EpollEvent) {
	// epll del
	err := unix.EpollCtl(epfd, unix.EPOLL_CTL_DEL, int(event.Fd), &event)
	if err != nil {
		fmt.Println(err)
	}
	err = unix.Close(int(event.Fd))
	if err != nil {
		fmt.Println(err)
	}
}

func SlaveHandler(rqs,wqs map[int32]ReadyQueue, event unix.EpollEvent, readHandle func(queue ReadyQueue) bool,
	writeHandle func(queue ReadyQueue) bool,writeCallback func(event unix.EpollEvent)) {
	rq := ReadyQueue{
		_type:       READ_QUEUE,
		event:       event,
		readBuf:     make([]byte,0,128),
		writeBuffer: make([]byte,0,128),
	}
	if !readHandle(rq) {
		rqs[event.Fd] = rq
		return
	}
	wq := rq
	wq._type = WRITE_QUEUE
	if !writeHandle(wq) {
		wqs[event.Fd] = wq
		writeCallback(event)
		return
	}
	wqs[event.Fd] = wq
}

func accept() {
	// 将套接字设置为非阻塞
	sockFd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM|unix.SOCK_NONBLOCK, 0)
	if err != nil {
		panic(err)
	}
	defer unix.Close(sockFd)
	// 为套接字绑定ip&port
	addr := unix.SockaddrInet4{
		Port: *port,
	}
	copy(addr.Addr[:], net.ParseIP(*ip).To4())
	err = unix.Bind(sockFd, &addr)
	if err != nil {
		panic(err)
	}
	// 监听文件描述符上的连接
	err = unix.Listen(sockFd, 5)
	if err != nil {
		panic(err)
	}
	// 将监听新描述符的套接字一起添加到epoll
	epollFd, err := unix.EpollCreate(65535)
	if err != nil {
		panic(err)
	}
	epfd = epollFd
	defer unix.Close(epollFd)
	// epoll事件
	event := unix.EpollEvent{
		Events: unix.EPOLLIN,
		Fd:     int32(sockFd),
	}
	err = unix.EpollCtl(epollFd, unix.EPOLL_CTL_ADD, sockFd, &event)
	if err != nil {
		panic(err)
	}
	epollEvents := make([]unix.EpollEvent, MAX_EVENTS)
	// 等待可读的epoll事件
	// 超时时间无限
	for {
		nEvent, err := unix.EpollWait(epollFd, epollEvents, -1)
		if err != nil {
			fmt.Println("epoll wait error: ", err)
			continue
		}
		for k, event := range epollEvents[:nEvent] {
			// accept
			if event.Fd == int32(sockFd) {
				fd, _, err := unix.Accept(int(event.Fd))
				if err != nil {
					panic(err)
				}
				// 将客户端连接的文件描述符设置为非阻塞模式
				oldOption, err := unix.FcntlInt(uintptr(fd), unix.F_GETFL, 0)
				if err != nil {
					panic(err)
				}
				newOption := oldOption | unix.O_NONBLOCK
				_, err = unix.FcntlInt(uintptr(fd), unix.F_SETFL, newOption)
				if err != nil {
					panic(err)
				}
				// 对客户端的连接监听事件设置为ET触发模式
				connEvent := unix.EpollEvent{
					Events: unix.EPOLLIN | unix.EPOLLET,
					Fd:     int32(fd),
				}
				// 将新连接的文件描述符放置到epoll的监听队列中
				err = unix.EpollCtl(epollFd, unix.EPOLL_CTL_ADD, fd, &connEvent)
				if err != nil {
					panic(err)
				}
			} else {
				// 非listen描述符的可读事件则处理数据
				threadFds[k%nWorkerThread()] <- event
			}
		}
	}
}

func main() {
	flag.Parse()
	initSlaveThreads()
	accept()
}
