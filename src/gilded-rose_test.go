package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	- 一旦销售期限过期，品质`Quality`会以双倍速度加速下降
	- 物品的品质`Quality`永远不会为负值
	- "Aged Brie"的品质`Quality`会随着时间推移而提高
	- 物品的品质`Quality`永远不会超过50
	- 传奇物品"Sulfuras"永不到期，也不会降低品质`Quality`
	- "Backstage passes"与aged brie类似，其品质`Quality`会随着时间推移而提高；当还剩10天或更少的时候，品质`Quality`每天提高2；当还剩5天或更少的时候，品质`Quality`每天提高3；但一旦过期，品质就会降为0

	* "Conjured"物品的品质`Quality`下降速度比正常物品快一倍
 */
func Test_GildedRose(t *testing.T) {
	assert.True(t,true)
}
