package main

import "fmt"

var c, python, java = true, false, "no!"

func main() {
	// 一般声明
	var name string
	name = "ssr"

	// 赋值声明
	var name2 = 22123

	// 简短声明
	name3 := 225

	// 多变量声明
	var ssr1, ssr2, ssr3 string
	ssr1, ssr2, ssr3 = "v1", "v2", "v3"

	// 多变量赋值声明
	var sst1, sst2, sst3 = "v1", "v2", "v3"

	// 集合类型
	var (
		ttd = "v1"
		tts = "v4"
	)

	println(name, name2, name3)
	println(ssr1, ssr2, ssr3)
	println(sst1, sst2, sst3)
	println(ttd, tts)

	dinfeng()
}

func dinfeng() {
	k := 3
	c, python, java := false, true, "yes!"
	// 注意检查，是使用的局部变量，还是修改全局变量，该方式容易造成错误！
	fmt.Println(c, python, java, k)
}
