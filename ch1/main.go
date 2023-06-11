// package 声明：
// 1.字母+下划线组合
// 2. 可以和文件夹名字不同
// 3. 同一个文件夹下的文件不能存在不同的包名，声明必须一致
package main

// 引入包语法形式：
// import [alias] [package_name]
// 包引入不用，会报错
import (
	"errors"
	"fmt"
	"unicode/utf8"

	// 匿名引入，为了执行包中的init方法，一般用作接口实现等操作
	_ "math"
)

func main() {
	fmt.Println("hello GO!")

	// string声明：
	// 一般用于短的，不用换行的，不含双引号
	fmt.Println("he said :\" u are cool! \"")
	// 长的，复杂的，用 json 中的，避免有过多转义符
	fmt.Println(`i said :" yes, i am cool!
	   i can sing,
	   i can skip,
	   i can also rap! `)

	// string长度：
	// 字节长度，和编码无关，用len()获取
	fmt.Println(len("你好")) // 6
	// 字符长度，和编码有关，用编码库获取，例如， utf8
	fmt.Println(utf8.RuneCountInString("你好"))   // 2
	fmt.Println(utf8.RuneCountInString("你好nb")) // 4

	// rune
	// rune 直观理解，就是字符
	// rune 本质int32 type rune = int32， 一个rune 4个字节
	// rune 在很多余语言里是没有的，与之对应的，golang中没有char类型，rune不是数字，不是char，不是byte
	// 实际开发中不常用
	var x rune // ctrl 点进去看下rune type定义
	x = 1323
	fmt.Println(x)

	// 变量声明
	var a int = 13  // 这里int可以省略
	var b uint = 15 // uint 绝对不可以省略，否则默认为uint
	// if a == b {     // 这里不会做类型转换

	// }
	fmt.Printf("a:%d , b: %d\n", a, b)

	// 常量声明
	// 驼峰命名
	const (
		constA = "this is const a" // 包内可访问
		ConstB = "this is const b" // 包外可访问
	)
	// 支持类型推断
	// 常量声明无法修改值

}

// 方法声名
func getStudent(name string, age int, others ...interface{}) (student Student, err error) {
	s := new(Student)
	s.name = name
	s.age = age
	for _, v := range others {
		switch t := v.(type) {
		case Gender:
			s.gender = t
		case uint:
			s.score = t
	}
	return *s, nil
}

type Gender int

const (
	man   Gender = 1
	woman Gender = 2
)

type Student struct {
	name   string
	age    int
	gender Gender
	score  uint
}