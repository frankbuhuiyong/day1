package main

import "fmt"

//切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。

func main() {
	//var a []string
	//var b []int
	//var c = []bool{false, true}
	//
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	////	基于数组定义切片
	//a := [5]int{55, 56, 57, 58, 59}
	//b := a[1:4]
	//fmt.Println(b)
	//fmt.Printf("%T\n", b)
	//// 切片后再切片
	//c := b[0:len(b)]
	//fmt.Println(c)
	//fmt.Printf("%T\n", c)
	////	make 函数构造切片（指定长度和容量）
	//d := make([]int, 5, 10)
	//fmt.Println(d)
	////通过len()和cap()获取切片的长度和容量
	//fmt.Println(len(d))
	//fmt.Println(cap(d))

	//	切片的本质是对底层数组的封装，它包含了三个信息：底层数组的指针，切片的长度和切片的容量。
	//	该指针指向第一个包含的元素的地址。
	// 切片不能直接比较

	//var a[]int //声明int类型切片
	//var b = []int{} //声明{}代表初始化
	//c := make([]int,0)
	//
	//if a == nil {
	//	fmt.Println("a is a nil slice")
	//}
	//fmt.Println(a, len(a), cap(a))
	//if b == nil {
	//	fmt.Println("b is a nil slice")
	//} else {
	//	fmt.Println("b is Not a nil slice")
	//}
	//
	//fmt.Println(b, len(b), cap(b))
	//
	//if c == nil {
	//	fmt.Println("c is a nil slice")
	//}else {
	//	fmt.Println("c is Not a nil slice")
	//}
	//fmt.Println(c, len(c), cap(c))

	//	切片的赋值和拷贝
	//	a := make([]int, 3) //[0 0 0]
	//	b := a //共用一个底层数组
	//	b[0] = 100
	//	fmt.Println(a)
	//	fmt.Println(b)

	////	切片的遍历
	//a := []int{1, 2, 3, 4, 5}
	//
	////通过索引遍历
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(i, a[i])
	//}
	//fmt.Println()
	////通过for range遍历
	//for i, value := range a {
	//	fmt.Println(i, value)
	//}

	//append（）方法和切片的扩容
	//var a []int //此时它并没有申请内存,panic: runtime error: index out of range [0] with length 0
	//一定要初始化，才能赋值
	//a[0] = 100
	//fmt.Println(a)

	//a = append(a, 10) //append相当于申请内存。
	//fmt.Println(a)

	//var a []int
	//for i := 0; i < 10; i++ {
	//	a = append(a, i)
	//	fmt.Printf("a value:%v, a len:%d, a cap:%d, a addr:%p\n", a, len(a), cap(a), &a)
	//	fmt.Printf("a value:%v, a len:%d, a cap:%d, a addr:%p\n", a, len(a), cap(a), a)
	//}

	//var a []int
	//a = append(a, 1, 2, 3, 4, 5)
	//fmt.Println(a)
	//b := []int {11, 22, 33, 44}
	//a = append(a, b...)
	//fmt.Println(a)

	//通过copy复制切片
	//a := []int{1, 2, 3, 4, 5}
	//b := make([]int, 5, 5)
	//copy(a, b) //复制一个切片到另外一个地址空间
	//fmt.Println(a)
	//fmt.Println(b)
	//b[0] = 100
	//fmt.Println(a)
	//fmt.Println(b)

	//切片元素的删除
	a := []string{"BeiJing","ZhuHai","ShenZhen","Guangzhou","ShiJiaZhuang"}
	//a = append(a[0:2], a[3:]...)
	//fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Printf(" a[%d] addr:%p, a addr:%p\n", i, &a[i],&a)
	}


}
