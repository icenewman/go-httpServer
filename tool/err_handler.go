package tools

import (
	"log"
	"os"
)

//主函数panic捕获恢复，输出到日志
func PanicHandler() {
	exeName := os.Args[0] //获取程序名称

	pid := os.Getpid() //获取进程ID
	// now := time.Now()  //获取当前时间
	// time_str := now.Format("20060102150405")                          //设定时间格式
	// fname := fmt.Sprintf("%s-%d-%s-dump.log", exeName, pid, time_str) //保存错误信息文件名:程序名-进程ID-当前时间（年月日时分秒）
	// log.Println("dump to file ", fname)

	//f, err := os.Create(fname)
	//if err != nil {
	//	return
	//}
	//defer f.Close()
	//
	if err := recover(); err != nil {
		//f.WriteString(fmt.Sprintf("%v\r\n", err)) //输出panic信息
		log.Printf("app:%s, pid:%s, %v\r\n", exeName, pid, err)
		//f.WriteString("========\r\n")
		//log.Println("========\r\n")
	}
	//
	//f.WriteString(string(debug.Stack())) //输出堆栈信息
}

