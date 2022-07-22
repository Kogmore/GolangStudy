package main

import (
	//1、"GolangStudy/5-init/lib1"
	//2、"GolangStudy/5-init/lib2"
	//3、
	_ "GolangStudy/5-init/lib1" //匿名导包方式   只能调用该包的init函数
	//4、
	mylib2 "GolangStudy/5-init/lib2" //别名导包方式  能调用该包的非私有函数
	//5、. "GolangStudy/5-init/lib2" //把lib2包下的所有非私有函数 导入到当前包  可直接使用
)

func main() {
	//1、 lib1.Lib1Test()
	//2、 lib2.Lib2Test()
	//4、
	mylib2.Lib2Test()
	//5、Lib2Test()
}
