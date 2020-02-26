package recsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//封装rpc server
func ServerRpc(host string,service interface{}) error {
	//注册了service 因此在客户端call的时候可以调用<service>.<Method>
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil{
		return err
	}
	log.Printf("Listening on %s",host)
	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Printf("accept error ;%v",err)
			continue
		}
		// 保持让接收servr连接的同时 还可以accept数据 ，后台运行处理连接
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

//封装rpc client
func ClientRpc(host string) (*rpc.Client,error) {
	conn, err := net.Dial("tcp", host)
	if err != nil{
		return nil,err
	}
	return jsonrpc.NewClient(conn),nil
}
