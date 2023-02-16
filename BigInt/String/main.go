package main

import (
	"fmt"
	"math/big"
)

func main() {
	s := big.NewInt(0)
	str := "怒发冲冠，凭栏处、潇潇雨歇。抬望眼，仰天长啸，壮怀激烈。三十功名尘与土，八千里路云和月。莫等闲，白了少年头，空悲切！靖康耻，犹未雪。臣子恨，何时灭！驾长车，踏破贺兰山缺。壮志饥餐胡虏肉，笑谈渴饮匈奴血。待从头、收拾旧山河，朝天阙。"
	fmt.Println(str)
	s.SetBytes([]byte(str))
	fmt.Println("s", s)

	i := big.NewInt(1)
	s.Add(s, i)
	fmt.Println("s", s)
	fmt.Println(string(s.Bytes()))
}
