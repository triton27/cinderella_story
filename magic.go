package main

import "fmt"

type Magic struct {
	Target *Human
	Broken chan int
}

func NewMagic(h *Human) *Magic {
	return &Magic{
		Target: h,
		Broken: make(chan int, 1),
	}
}

// %v は全ての型に使えるデフォルトフォーマット verb(意味: 動詞)
// 値のデフォルトのフォーマットでの表現を出力する
func (m *Magic) GenerateDress() Costume {
	fmt.Printf("%v は魔法のドレスを作ってもらった！\n", m.Target.Name)
	return NewDress(m.Target)
}

func (m *Magic) GenerateGlassShoes() *Shoes {
	fmt.Printf("%v は魔法のガラスの靴を作ってもらった！\n", m.Target.Name)
	return NewShoes(m.Target)
}
