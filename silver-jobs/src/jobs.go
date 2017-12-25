package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

// 公司结构体
type Company struct {
	Name            string   // 公司名称
	Address         string   // 公司地址
	PrivateKey      string   // 私钥
	PublicKey       string   // 公钥
	EmployeeAddress []string // 雇员地址列表
}

// 雇员结构体
type Employee struct {
	Name       string // 雇员名称
	Address    string // 雇员地址
	WorkInfoId []int  // 工作信息ID
}

// 雇员工作信息结构体
type WorkInfo struct {
	Id       int   // ID
	JoinTime int64 // 加入时间
	ExitTime int64 // 离职时间
	Status   int   // 0.在职 1.试用期 2.离职
}

// 记录结构体
type Record struct {
	Id              int    // 记录ID
	CompanyAddress  string // 公司地址
	EmployeeAddress string // 雇员地址
	CompanySign     string // 公司签名
	ModifyTime      int64  // 修改时间
	ModifyOperation string // 0.入职 1.转正 2.正常离职 3.试用期离职 4.劝退 5.开除
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)
	if function == "test1" { //自定义函数名称
		return t.test1(stub, args) //定义调用的函数
	}
	return shim.Error("Received unknown function invocation")
}

func (t *SimpleChaincode) test1(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success([]byte("Called test1"))
}

func main() {
	fmt.Print("Hello, silver jobs.\n")
}
