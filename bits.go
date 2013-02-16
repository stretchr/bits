package bits

// bits provides a simple interface for interacting with an integer in a 
// bitwise fashion. The integer is an int64, meaning you have 64 buckets
// of usable data.
type Bits struct {
	bits uint64
}

func MakeBits(initialBits uint64) *Bits {
	b := new(Bits)
	b.bits = initialBits
	return b
}

// Bits gets the underlying uint64
func (b *Bits) Bits() uint64 {
	return b.bits
}

// Reset sets all bits to zero
func (b *Bits) Reset() *Bits {
	b.bits = 0
	return b
}

// Set sets the bit at "bit" offset
func (b *Bits) Set(bit uint64) *Bits {
	b.bits |= 1 << bit
	return b
}

// SetMask sets the bits that are in "mask"
func (b *Bits) SetMask(mask uint64) *Bits {
	b.bits |= mask
	return b
}

// SetBits sets the bits that are in the Bits object
func (b *Bits) SetBits(bits *Bits) *Bits {
	b.bits |= bits.bits
	return b
}

// Test determines if the bit at "bit" offset is set
func (b *Bits) Test(bit uint64) bool {
	return (b.bits & (1 << bit)) != 0
}

// TestMask determines if all bits in "mask" are set
func (b *Bits) TestMask(mask uint64) bool {
	return (b.bits & mask) == mask
}

// TestBits determines if all bits in the Bits object are set
func (b *Bits) TestBits(bits *Bits) bool {
	return (b.bits & bits.bits) == bits.bits
}

// Clear clears the bit at "bit" offset
func (b *Bits) Clear(bit uint64) *Bits {
	b.bits &^= 1 << bit
	return b
}

// ClearMask clears the bits set in "mask"
func (b *Bits) ClearMask(mask uint64) *Bits {
	b.bits &^= mask
	return b
}

// ClearBits clears the bits set in the Bits object
func (b *Bits) ClearBits(bits *Bits) *Bits {
	b.bits &^= bits.bits
	return b
}

// Toggle toggles the bit at "bit" offset
func (b *Bits) Toggle(bit uint64) *Bits {
	b.bits ^= 1 << bit
	return b
}

// ToggleMask toggles the bits in "mask"
func (b *Bits) ToggleMask(mask uint64) *Bits {
	b.bits ^= mask
	return b
}

// ToggleBits toggles the bits in the Bits object
func (b *Bits) ToggleBits(bits *Bits) *Bits {
	b.bits ^= bits.bits
	return b
}
