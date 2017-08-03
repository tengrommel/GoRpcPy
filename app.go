package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"fmt"
)

var ii int

func main() {
	println(ii)
	// 监听端口
	lis, err := net.Listen("tcp", "0.0.0.0:9789")
	if err != nil{
		return
	}
	// 注册关闭
	defer lis.Close()

	// 启动新的服务
	srv := rpc.NewServer()
	// 注册类型名称
	if err := srv.RegisterName("Json_type", new(Json_type)); err != nil{
		return
	}
	if err := srv.RegisterName("JJ", new(JJ)); err != nil{
		return
	}

	for  {
		conn, err := lis.Accept()
		if err != nil{
			log.Fatalf("lis.Accept(): %v\n", err)
			continue
		}
		// 创建新的goroutine提供rpc服务
		go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

type Json_type struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

// Json结构体的方法
func (c *Json_type)Get(args Json_type, result *Json_type) error {
	if args.Name == "红雀"{
		log.Println(args.Name,"GET操作")
		*result = Json_type{"红雀", 34}
		return nil
	}else {
		log.Printf("不明类型 %v GET操作", args.Name)
		*result = Json_type{"其他的类型",56}
		return nil
	}
	return errors.New("Input Name")
}

func (c *Json_type)Find(args Json_type, result *Json_type) error {
	if args.Name == "红雀"{
		log.Println(args.Name,"Find操作")
		*result = Json_type{"红雀", 34}
		return nil
	}else {
		log.Printf("不明类型 %v Find操作", args.Name)
		*result = Json_type{"其他的类型",56}
		return nil
	}
	return errors.New("Input Name")
}

type JJ struct {
	Name string `json:"name"`
	Number int `json:"number"`
	Counter int `json:"counter"`
}

func (c *JJ)Inc_Number(args JJ, result *JJ) error {
	if name, ok := interface{}(args.Name).(string); ok{
		fmt.Printf("你好，%v ", name)
		if number, ok := interface{}(args.Number).(int); ok{
			fmt.Printf("number，%v ", number)
			if counter, ok := interface{}(args.Counter).(int); ok{
				fmt.Printf("counter, %v\n", counter)
				args.Number++
				*result = args
				return nil
			}else {return errors.New("JJ counter error")}
		} else {return errors.New("JJ number error")}
	}else {return errors.New("JJ name error")}
}

func (c *JJ)Inc_Counter(args JJ, result *JJ) error {
	if name, ok := interface{}(args.Name).(string); ok{
		fmt.Printf("你好，%v ", name)
		if number, ok := interface{}(args.Number).(int); ok{
			fmt.Printf("number，%v ", number)
			if counter, ok := interface{}(args.Counter).(int); ok{
				fmt.Printf("counter, %v\n", counter)
				args.Counter++
				*result = args
				go func() {for ii:=0;ii<10000;ii++{ii++;fmt.Printf("Inc: %v\n",ii)}}()
				return nil
			}else {return errors.New("JJ counter error")}
		} else {return errors.New("JJ number error")}
	}else {return errors.New("JJ name error")}
}

func (c *JJ)Find(args JJ, result *JJ) error {
	if _,ok :=interface{}(args).(JJ);ok{
		go func() {fmt.Println("异步执行")}()
		go func() {for ii:=0;ii<10000;ii++{ii++; fmt.Printf("find: %v\n",ii)}}()
		*result = args
		return nil
	}
	return errors.New("JJ error")
}