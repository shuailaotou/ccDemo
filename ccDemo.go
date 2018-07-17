package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleChainCode struct {
}

//在链码实例化或者升级时被自动调用
func (t *SimpleChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	//获取参数
	args := stub.GetStringArgs()
	//判断参数长度是否为2
	if len(args) != 2 {

		return shim.Error("指定了错误的参数个数")

	}
	err:=stub.PutState(args[1],[]byte(args[2]))
	if err != nil {
		fmt.Printf("存储数据错误：%v",err)
	}
	fmt.Println("数据保存成功")

	return shim.Success(nil)

}

//在链码进行调用时自动执行Invoke方法
func (t *SimpleChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	//获取调用链码是传递的参数内容
	fun,args := stub.GetFunctionAndParameters()

	if fun == "query" {
		return query(stub,args)
	}

	return shim.Error("非法操作，指定功能不能实现")
}

func query(stub shim.ChaincodeStubInterface,args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("只能指定相应的Key")
	}

	result,err:= stub.GetState(args[1])
	if err != nil {
		return shim.Error("根据指定的"+args[1]+"查询数据是发生错误")
	}
	if result == nil {
		shim.Error("根据指定的"+args[1]+"没有查询到数据")
	}

	return shim.Success(result)


}
func main() {

	err := shim.Start(new(SimpleChainCode))
	if err != nil {
		fmt.Printf("%v链码启动失败", err)
	}

}
