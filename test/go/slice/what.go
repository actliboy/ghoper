package main

import "fmt"

/*
在有些时候，我们还可以在方括号中放入第三个正整数，如下所示：

numbers3[1:4:4]
这第三个正整数被称为容量上界索引。它的意义在于可以把作为结果的切片值的容量设置得更小。换句话说，它可以限制我们通过这个切片值对其底层数组中的更多元素的访问。下面举个例子。让我们先来回顾下在上一节讲到的numbers3和slice1。针对它们的赋值语句是这样的：

var numbers3 = [5]int{1, 2, 3, 4, 5}
var slice1 = numbers3[1:4]
这时，变量slice1的值是[]int{2, 3, 4}。但是我们可以通过如下操作将其长度延展得与其容量相同：

slice1 = slice1[:cap(slice1)]
通过此操作，变量slice1的值变为了[]int{2, 3, 4, 5}，且其长度和容量均为4。现在，numbers3的值中的索引值在[1,5)范围内的元素都被体现在了slice1的值中。这是以numbers3的值是slice1的值的底层数组为前提的。这意味着，我们可以轻而易举地通过切片值访问其底层数组中对应索引值更大的更多元素。如果我们编写的函数返回了这样一个切片值，那么得到它的程序很可能会通过这种技巧访问到本不应该暴露给它的元素。这是确确实实是一个安全隐患。

如果我们在切片表达式中加入了第三个索引（即容量上界索引），如：

var slice1 = numbers3[1:4:4]
那么在这之后，无论我们怎样做都无法通过slice1访问到numbers3的值中的第五个元素。因为这超出了我们刚刚设定的slice1的容量。如果我们指定的元素上界索引或容量上界索引超出了被操作对象的容量，那么就会引发一个运行时恐慌（程序异常的一种），而不会有求值结果返回。因此，这是一个有力的访问控制手段。

虽然切片值在上述方面受到了其容量的限制，但是我们却可以通过另外一种手段对其进行不受任何限制地扩展。这需要使用到内建函数append。append会对切片值进行扩展并返回一个新的切片值。使用方法如下：

slice1 = append(slice1, 6, 7)
通过上述操作，slice1的值变为了[]int{2, 3, 4, 6, 7}。注意，一旦扩展操作超出了被操作的切片值的容量，那么该切片的底层数组就会被自动更换。这也使得通过设定容量上界索引来对其底层数组进行访问控制的方法更加严谨了。
*/
func main()  {
	s:=[]int{1,2,3,4,5,6}
	s =s[0:1:2]//切片设置容量上限操作
	fmt.Printf("%v",s)
}
