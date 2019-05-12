package main

import (
	"fmt"
	"time"
)

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

// 舞踏会が開催
ball.Start()
finished := make(chan int, 1)
go func() {
	for !ball.IsFinished() {
		<- time.After(1 * time.Second):
		ball.Dancing()
		// 24時になったら limit に 1 が送信される
		if ball.Clock == 24 {
			limit <- 1
		}
	}
	ball.Finish()
	finished <- 1
}()

<-magic.Broken
// シンデレラが急いで帰る
ball.Exit(cinderella)
// ガラスの靴を落としてしまう！！
falledShoes := cinderella.Shoes
cinderella.Shoes = nil

// 王子、ガラスの靴を見つける
foundShoes := falledShoes

// 舞踏会終了
<-finished

// 靴の持ち主を舞踏会の参加者の中から探している
for _, h := range ball.Entries {
	if h.Gender == Woman {
		if foundShoes.Wear(h) {
			fmr.Println("見つけた！")
		} else {
			// ドレスルームにドレスがある人が h.Name に入ってる
			fmt.Printf("%v: %vさんの靴ではない\n", price.Name, h.Name)
		}
	}
}

// 靴がぴったりなシンデレラ
if foundShoes.Wear(cinderella) {
	fmt.Println("見つけた！")
}
