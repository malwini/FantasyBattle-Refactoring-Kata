package codingdojo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

// | Where     | What            | Base Damage | Damage Modifier |
// |-----------|-----------------|-------------|-----------------|
// | right hand|  flashy sword of danger | 10  | 1.0             |
// | right hand|  excalibur              | 20  | 1.5             |
// | left hand |  round shield           |  0  | 1.4             |
// | feet      |  ten league boots       |  0  | 0.1             |
// | head      |  helmet of swiftness    |  0  | 1.2             |
// | chest     |  breastplate of steel   |  0  | 1.4             |

// choose this one if you are familiar with mocks
type MockItem struct {
}

func (m MockItem) GetBaseDamage() int32 {
	return 20
}

func (m MockItem) GetDamageModifier() float64 {
	return 1.0
}

type MockBuff struct {
}

func (buff MockBuff) SoakModifier() float64 {
	return 1.0
}

func (buff MockBuff) DamageModifier() float64 {
	return 1.0
}

type MockArmor struct {
}

func (armor MockArmor) GetDamageSoak() int32 {
	return 5
}

func TestDamageCalculationsWithMocks(t *testing.T) {

	mockitem := MockItem{}
	equipment := MakeEquipment(mockitem, mockitem, mockitem, mockitem, mockitem)
	inventory := MakeInventory(equipment)

	// (equipment)
	//inventory := mockInventory()
	stats := MakeStats(10)
	mockBuffs := []Buff{MockBuff{}}
	target := MakeSimpleEnemy(MockArmor{}, mockBuffs)

	damage := MakePlayer(inventory, stats).CalculateDamage(target)

	assert.EqualValues(t, 590, damage.GetAmount())
}

// choose this one if you are not familiar with mocks
func TestDamageCalculations(t *testing.T) {
	t.Skip("Test is not finished yet")
	// inventory := MakeInventory(/* TODO */)
	// stats := MakeStats(0)
	// target := MakeSimpleEnemy(/* TODO */, /* TODO */)

	// damage := MakePlayer(inventory, stats).CalculateDamage(target)

	// assert.EqualValues(t, 10, damage.GetAmount())
}
