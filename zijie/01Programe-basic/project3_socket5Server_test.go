package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	_ "expvar"
	"fmt"
	"io"
	"log"
	"net"
	"testing"
)

//定义的版本字（十六进制表示）
const socks5Ver = 0x05 //默认socks的协议版本：5
const cmdBind = 0x01  //请求阶段的cmd字段：1表示连接
const atypIPV4 = 0x01   //请求阶段的atyp字段：1表示IPV4【4个字节】
const atypeHOST = 0x03  //请求阶段的atyp字段：3表示域名
const atypeIPV6 = 0x04  //请求阶段的atyp字段：4表示IPV6【6个字节】

func Test_Project3_socket5server(t *testing.T) {
	//监听本地的1080端口
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		//作为报告致命错误的一种方式，当某些不应该发生的场景发生时,我们就应该调用panic。【替代try catch】
		panic(err)
	}
	for  {
		//监听接收请求，拿到一个连接
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept faild %v", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	//函数执行完毕最后进行执行：关闭连接
	defer conn.Close()
	reader := bufio.NewReader(conn)  //创建带缓冲的只读流。【缓冲效果：减少底层系统调用次数，主要在这里是方便一个字节一个字节来读取，底层可能会合并几次大的读取】
	//1、认证操作（会进行认证结果响应）
	err := auth(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
	log.Println("auth success")

	//2、连接操作
	err = connect(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
}

//连接请求
func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 版本号，socks5的值为0x05
	// CMD 0x01表示CONNECT请求
	// RSV 保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名
	// DST.ADDR 一个可变长度的值
	// DST.PORT 目标端口，固定2个字节

	//1、读取前面四个字段，分别是VER、CMD、RSV、ATYP。之后会进行一个字段校验
	buf := make([]byte, 4)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("read header failed:%w", err)
	}
	ver, cmd, atyp := buf[0], buf[1], buf[3]  //测试时， curl --socks5 127.0.0.1:1080 -v http://www.qq.com，服务端这边拿到的type是1，也就是ipv4地址【http://www.qq.com => 183.942.238.19】
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not supported cmd:%v", ver)
	}
	addr := ""
	switch atyp {
		//若是IPV4，直接读取4个字节，使用上面的buf即可
		case atypIPV4:
			_, err = io.ReadFull(reader, buf)
			if err != nil {
				return fmt.Errorf("read atyp failed:%w", err)
			}
			//将四个字节转为一个字符串："183.942.238.19"
			addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
		//若是域名，会先读取1个字节，该字节表示域名的长度，接着来进行指定字节读取
		case atypeHOST:
			hostSize, err := reader.ReadByte()
			if err != nil {
				return fmt.Errorf("read hostSize failed:%w", err)
			}
			host := make([]byte, hostSize)
			_, err = io.ReadFull(reader, host)
			if err != nil {
				return fmt.Errorf("read host failed:%w", err)
			}
			addr = string(host)
		//对于IPV6暂不支持
		case atypeIPV6:
			return errors.New("IPv6: no supported yet")
		default:
			return errors.New("invalid atyp")
	}
	//2、读取DST.PORT，2个字节到buf数组中的[0,2)
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed:%w", err)
	}
	//将两个字节的端口号转为一个Uint16的一个属性
	port := binary.BigEndian.Uint16(buf[:2])  //对应80端口

	//3、拼接ip地址+port来进行发送tcp请求
	//dest也就是对应的connect连接
	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		return fmt.Errorf("dial dst failed:%w", err)
	}
	defer dest.Close()
	log.Println("dial", addr, port)

	//最后：进行响应结果，也就是返回一个包
	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER socks版本，这里为0x05
	// REP Relay field,内容取值如下 X’00’ succeeded
	// RSV 保留字段
	// ATYPE 地址类型
	// BND.ADDR 服务绑定的地址
	// BND.PORT 服务绑定的端口DST.PORT
	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write failed：%w", err)
	}
	//防止connect函数会立即返回，通过使用context机制，来用context连with cancel来创建一个context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//实现一个浏览器与下游服务器的双向数据转发：通过io.copy即可实现一个单向数据转发，双向转发需要启动两个goroutine
	//下面就是一个双向转发的实现
	go func() {
		_, _ = io.Copy(dest, reader)  //目标服务器返回结果 => 当前服务器读对象
		cancel()
	}()
	go func() {
		_, _ = io.Copy(conn, dest)  //客户端发送数据 => 目标服务器
		cancel()
	}()

	//一旦cancel被调用就会停止阻塞继续向下执行返回
	<- ctx.Done()
	return nil
}



//认证
/**
	reader：读取工具类
	conn：连接
 */
func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	//1、进行认证操作【读取三个字段】
	//1-1、读取第一个字节，进行判断是否为sockets5
	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}

	//1-2、读取认证的方法数量【这个数量为多少就表示第三个methods的字节长度为多少】
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	method := make([]byte, methodSize)

	//1-3、读取指定数量的第三个字段
	_, err = io.ReadFull(reader, method) // 当进行readfull读取reader数据时，reader中有一个隐形指针指向最后读到数据的后一位，
	// 当下一次进行readfull时从当前指针的位置进行读取，即前一次读取的数据不再进行读取，直接略过。【简而言之：读取剩下来的所有字节】
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}
	log.Println("ver", ver, "method", method) //打印一下版本号以及对应的方法

	//2、进行响应数据，响应的协议如下：
	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+
	_, err = conn.Write([]byte{socks5Ver, 0x00})  //第一个字段就是sockets的版本号，第二个字段就是选中的鉴传方法。当前0x00含义是不需要认证。
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}
	return nil  //若是整个过程中无任何异常，直接返回nil空
}