package rpc

import (
	"net"
	"reflect"
)

//RPC客户端的实现

//实现通用的RPC客户端
//绑定RPC访问的方法
//传入访问的函数名
//函数具体实现在Server端，Client只有函数原型 (函数原型是指由函数定义中抽取出来的能代表函数应用特征的部分，包括函数的数据类型、函数名称、形式参数说明。)
//使用MakeFunc() 完成原型到函数的调用
//注：Client端不会直接调用Server端。

//声明客户端
type Client struct {
	conn net.Conn
}

//创建一个客户端对象
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

//实现通用的RPC客户端
//绑定RPC访问的方法 -(callRPC)
//传入访问的函数名 -(rpcName)
//函数具体实现在Server端，Client只有函数原型 -(fPtr)
//fPtr: 指向函数原型(函数原型是指由函数定义中抽取出来的能代表函数应用特征的部分，包括函数的数据类型、函数名称、形式参数说明。)
//调用：xxx.callRPC("queryUser",&query)
func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	//通过反射，获取fPtr未初始化的函数原型-fn
	fn := reflect.ValueOf(fPtr).Elem()
	//定义另一个函数f，作用是对第一个函数的参数操作
	//f完成与Server的交互
	f := func(args []reflect.Value) []reflect.Value {
		//处理输入的参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		//创建连接
		cliSession := NewSession(c.conn)
		//编码数据
		reqRPC := RPCData{Name: rpcName, Args: inArgs}
		b, err := encode(reqRPC)
		if err != nil {
			panic(err)
		}
		//写出数据
		err = cliSession.Write(b)
		if err != nil {
			panic(err)
		}
		//读取响应数据
		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}
		//解码数据
		respRPC, err := decode(respBytes)
		if err != nil {
			panic(err)
		}
		//处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for i, arg := range respRPC.Args {
			//必须进行nil
			if arg == nil {
				//必须填充一个真正的类型，不能是nil
				//reflect.Zero(fn.Type().Out(i)) 返回该类型零值的Value
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}

	//参数1：一个未初始化函数的方法值，类型是reflect.Type
	//参数2：另一个函数，作用是对第一个函数参数操作
	//返回 reflect.Value 类型
	//MakeFunc 使用传入的函数原型，创建一个绑定参数2的新函数
	v := reflect.MakeFunc(fn.Type(), f)
	//为函数fPtr赋值
	fn.Set(v)
}
