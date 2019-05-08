package main

// 衣装インターフェース
type Costume interface {
	Wear(h *Human) bool
}

// ドレス
type Dress struct {
	Owner *Human
}

func NewDress(h *Human)Costume {
	return &Dress{
		Owner: h,
	}
}

// 所有者のみ着ることができる
func (d *Dress) Wear(h *Human) bool {
	return h == d.Owner && h.Gender == Woman
}

// 燕尾服
type Tailcoat struct {
	Owner *Human
}

func NewTailcoat(human *Human) Costume {
	return &Tailcoat{
		Owner: h,
	}
}

func (t *Tailcoat) Wear(h *Human) bool {
	return h == t.Owner && h.Gender == Man
}
