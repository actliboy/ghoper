package b

import "test/struct/mixture/a"

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/1
 * @description：
 */
type A struct {
	I int
	J int
}

type B struct {
	a.A
	a.B
}
