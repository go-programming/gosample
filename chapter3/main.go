package main //定义package为main才能执行下面的main函数，因为main函数只能在package main 中执行

//简写版的import导入依赖的项目
import (
	"fmt"  //使用其下的Println函数
	"os"   //使用其下的Stdout常量定义
	"time" // 使用time包下的时间格式常量定义RFC3339

	"github.com/sirupsen/logrus"            //日志组件
	"github.com/souyunkutech/gosample/util" //自己写的工具包，下面有自定义的函数统一使用
)

//声明log变量是logrus.Logger的指针类型，使用时不需要带指针
var log *logrus.Logger

// 初始化函数，先于main函数执行
func init() {
	log = logrus.New()            //使用logrus包下的New()函数进行logrus组件的初始化
	log.Level = logrus.DebugLevel //将log变量中的Level字段设置为logrus下的DebugLevel
	log.Out = os.Stdout
	log.Formatter = &logrus.TextFormatter{ //因为log.Formatter被声明为指针类型，所以对其赋值也是需要使用‘&’关键字将其地址赋值给该字段
		TimestampFormat: time.RFC3339, //使用time包下的RFC3339常量，赋值时如果字段与大括号不在一行需要在赋值后面添加逗号，包括最后一个字段的赋值！！！
	}
}

//定义常量PI
const PI = 3.1415926

//定义Student结构体，可以统一使用该结构来生命学生信息
type Student struct {
	Name    string //姓名对外可见（首字母大写）
	age     int    //年龄不能随便让人知道，所以对外不可见
	classId int    //班级也是
}

//main函数，程序执行的入口
func main() {
	var hello = "hello world" //定义hello变量，省略了其类型string的声明
	fmt.Println(hello)        //使用fmt包下的Println函数打印hello变量

	//多个变量的定义和赋值，使用外部函数生成
	length, width, height := util.RandomShape() //使用其他package的函数

	//多个变量作为外部函数的参数
	size := util.CalSize(length, width, height)
	log.Infof("length=%d, width=%d, height=%d, size=%d", length, width, height, size) //使用日志组件logrus的函数进行打印长宽高和size

	var student = Student{Name: "小明", age: 18, classId: 1}                                         //声明学生信息，最后一个字段的赋值不需要添加逗号
	log.Debugf("学生信息: 学生姓名: %s, 学生年龄: %d, 学生班级号: %d ", student.Name, student.age, student.classId) //使用日志组件logrus的函数进行打印学生信息
}
