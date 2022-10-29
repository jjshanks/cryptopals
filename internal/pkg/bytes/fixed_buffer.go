package bytes

type FixedBuffer struct {
	Fix    byte
	Length int
}

func (b *FixedBuffer) Len() int {
	return b.Length
}

func (b *FixedBuffer) ReadByte() (byte, error) {
	return b.Fix, nil
}
