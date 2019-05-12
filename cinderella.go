stepMother := NewHuman("StepMother", 52, Woman)
sisterA := NewHuman("SisterA", 23, Woman)
sisterB := NewHuman("SisterB", 20, Woman)
cinderella := NewHuman("ella", 18, Woman)

stepMother.Say("ネコに餌をやれ")
sisterA.Say("洗濯しておいて")
sisterB.Say("お腹空いた")
cinderella.Say("...")

// 舞踏会は19時から27時まで開催される
ball := NewBall(19, 27)

// 舞踏会用のドレスを用意する
// シンデレラのドレスは用意しない
dressRoom := NewDressRoom()
dressRoom.Store(stepMother, sisterA, sisterB)

// 継母のドレスはある
dressRoom.GetDress(stepMother)

// 舞踏会参加
ball.Entry(stepMother)

// 姉 A のドレスもある
dressRoom.GetDress(sisterA)
ball.Entry(sisterA)

// 姉 B のドレスもある
dressRoom.GetDress(sisterB)
ball.Entry(sisterB)

// シンデレラのドレスだけない
dressRoom.GetDress(cinderella)
ball.Entry(cinderella)

magic := NewMagic(cinderella)
cinderella.SetCostume(magic.GeneraeDress())
cinderella.SetShoes(magic.GenerateGlassShoes())

limit := make(chan int, 1)
go magic.Limit(limit)
// シンデレラが舞踏会に参加する
ball.Entry(cinderella)

// 王子登場
prince := NewHuman("王子", 18, Man)
tailcoat := NewTailcoat(prince)
prince.SetCostume(tailcoat)
ball.Entry(prince)
