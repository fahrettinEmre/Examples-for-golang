/* // tank struct
name
location x y
bu tanka move metodu ekle
x ve y koordinatı alacak
game table tanımla yani bir game struct olacak
bu game struct içinde tanks diye değişken olacak
game start dediğinde game in start metodu olacak
içindeki tankları random bir kere hareket ettirecek */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tank struct {
	name      string
	locationx int
	locationy int
}

type gameTable struct {
	tanks []Tank
}

func newTable() gameTable {
	var game1 gameTable
	return game1
}

func (g gameTable) ekle(t Tank) {
	g.tanks = append(g.tanks, t)
}

func (g gameTable) tanklarıhareket() {
	for _, t := range g.tanks {
		t.move()
	}
}

func (t *Tank) move() {
	rand.Seed(time.Now().UnixNano())
	rx := rand.Intn(100)
	rand.Seed(time.Now().UnixNano())
	ry := rand.Intn(100)

	t.locationx += rx
	t.locationy += ry
}

func newTank(s string) Tank {
	var tank Tank
	tank.name = s
	tank.move()
	return tank
}

func main() {
	table := newTable()
	tank1 := newTank("TankA")
	tank2 := newTank("TankB")

	table.ekle(tank1)
	table.ekle(tank2)

	fmt.Println(tank1, tank2)

	table.tanklarıhareket()
	fmt.Println(tank1, tank2)

}
