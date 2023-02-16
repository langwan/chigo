package main

import (
	"fmt"
	"math/big"
)

func main() {
	m := big.NewInt(0)
	n := big.NewInt(1)
	str := "怒发冲冠，凭栏处、潇潇雨歇。抬望眼，仰天长啸，壮怀激烈。三十功名尘与土，八千里路云和月。莫等闲，白了少年头，空悲切！靖康耻，犹未雪。臣子恨，何时灭！驾长车，踏破贺兰山缺。壮志饥餐胡虏肉，笑谈渴饮匈奴血。待从头、收拾旧山河，朝天阙。"
	for i := 0; i < 10000; i++ {
		fmt.Println("-s-")
		str += str
		fmt.Println("-e-")
		m.SetBytes([]byte(str))
		m.Add(m, n)
		fmt.Println(len(m.String()))
	}

}
