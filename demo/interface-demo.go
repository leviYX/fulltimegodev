package demo

type ReturnCode int

const (
	OK  ReturnCode = 200
	BAD ReturnCode = 500
)

type DataSync interface {
	syncAll() (ReturnCode, error)
}

type GangLianData struct {
	name string
}

func (glData GangLianData) syncAll() (ReturnCode, error) {
	// todo 同步钢联数据
	return OK, nil
}

//func main() {
//	gl := GangLianData{
//		name: "Rex",
//	}
//	var dataSync DataSync = gl
//	code, err := dataSync.syncAll()
//	if err != nil {
//		panic(err)
//	}
//	if ReturnCode(code) == OK {
//		fmt.Println("同步成功")
//	} else {
//		fmt.Println("同步失败")
//	}
//
//	var a int = 0
//	if a = 1; 1 != 2 {
//		fmt.Println(a)
//	}
//}
