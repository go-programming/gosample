package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/souyunkutech/gosample/util"
	"os"
	"time"
)

var log *logrus.Logger

func init() {
	log = logrus.New() //使用logrus包下的New()函数进行logrus组件的初始化
	log.Level = logrus.DebugLevel
	log.Out = os.Stdout
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.RFC3339, //使用time包下的RFC3339常量
	}
}

const PI = 3.1415926



type Student struct {
	Name    string
	age     int //年龄不能随便让人知道，所以对外不可见
	classId int //班级也是
}
func main() {
	var hello = "hello world"
	fmt.Println(hello)
	length, width, height := util.RandomShape() //使用其他package的函数

	size := util.CalSize(length, width, height)
	log.Infof("length=%d, width=%d, height=%d, size=%d", length, width, height, size)

	var student = Student{Name: "小明", age: 18, classId: 1}
	log.Debugf("学生信息: 学生姓名: %s, 学生年龄: %d, 学生班级号: %d ", student.Name, student.age, student.classId)
}
