package bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeBits(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {
		assert.Equal(t, b.bits, bitMask)
	}
}

func TestBits(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {
		assert.Equal(t, b.Bits(), bitMask)
	}
}

func TestReset(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {

		b.Reset()
		assert.Condition(t, func() bool {
			return b.bits == 0
		})
	}

}

func TestSet(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {

		assert.NotEqual(t, b.bits, bitMask|1<<4)

		b.Set(4)

		assert.Equal(t, b.bits, bitMask|1<<4)

	}
}

func TestSetMask(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<2 | 1<<3
	var newMask uint64 = 1<<1 | 1<<4
	var resultMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3 | 1<<4

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {

		assert.NotEqual(t, b.bits, newMask)

		b.SetMask(newMask)

		assert.Equal(t, b.bits, resultMask)

	}
}

func TestSetBits(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<2 | 1<<3
	var newMask uint64 = 1<<1 | 1<<4
	var resultMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3 | 1<<4

	b := MakeBits(bitMask)
	b2 := MakeBits(newMask)

	if assert.NotNil(t, b) {

		assert.NotEqual(t, b.bits, newMask)

		b.SetBits(b2)

		assert.Equal(t, b.bits, resultMask)

	}
}

func TestTest(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {

		assert.Equal(t, b.Test(2), true)

	}
}

func TestTestMask(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3
	var newMask uint64 = 1<<1 | 1<<3
	var badMask uint64 = 1<<1 | 1<<4

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {

		assert.True(t, b.TestMask(newMask))
		assert.False(t, b.TestMask(badMask))

	}
}

func TestTestBits(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3
	var newMask uint64 = 1<<1 | 1<<3
	var badMask uint64 = 1<<1 | 1<<4

	b := MakeBits(bitMask)
	b2 := MakeBits(newMask)
	b3 := MakeBits(badMask)

	if assert.NotNil(t, b) {

		assert.True(t, b.TestBits(b2))
		assert.False(t, b.TestBits(b3))

	}
}

func TestClear(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3
	var targetbitMask uint64 = 1<<0 | 1<<1 | 1<<2

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {

		assert.NotEqual(t, b.bits, targetbitMask)

		b.Clear(3)

		assert.Equal(t, b.bits, targetbitMask)

	}
}

func TestClearMask(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3
	var clearMask uint64 = 1<<1 | 1<<3
	var newMask uint64 = 1<<0 | 1<<2

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {

		b.ClearMask(clearMask)
		assert.Equal(t, b.bits, newMask)

	}
}

func TestClearBits(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3
	var clearMask uint64 = 1<<1 | 1<<3
	var newMask uint64 = 1<<0 | 1<<2

	b := MakeBits(bitMask)
	b2 := MakeBits(clearMask)

	if assert.NotNil(t, b) {

		b.ClearBits(b2)
		assert.Equal(t, b.bits, newMask)

	}
}

func TestToggle(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3
	var targetbitMask uint64 = 1<<0 | 1<<2 | 1<<3

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {

		assert.NotEqual(t, b.bits, targetbitMask)

		b.Toggle(1)

		assert.Equal(t, b.bits, targetbitMask)

	}
}

func TestToggleMask(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3
	var toggleMask uint64 = 1<<1 | 1<<3 | 1<<4
	var newMask uint64 = 1<<0 | 1<<2 | 1<<4

	b := MakeBits(bitMask)

	if assert.NotNil(t, b) {

		b.ToggleMask(toggleMask)
		assert.Equal(t, b.bits, newMask)

	}
}

func TestToggleBits(t *testing.T) {

	var bitMask uint64 = 1<<0 | 1<<1 | 1<<2 | 1<<3
	var toggleMask uint64 = 1<<1 | 1<<3 | 1<<4
	var newMask uint64 = 1<<0 | 1<<2 | 1<<4

	b := MakeBits(bitMask)
	b2 := MakeBits(toggleMask)

	if assert.NotNil(t, b) {

		b.ToggleBits(b2)
		assert.Equal(t, b.bits, newMask)

	}
}
