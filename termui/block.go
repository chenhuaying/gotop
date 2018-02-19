package termui

import (
	"image"
)

// Block is a base struct for all other upper level widgets,
// consider it as css: display:block.
// Normally you do not need to create it manually.
type Block struct {
	Grid     image.Rectangle
	X        int
	Y        int
	XOffset  int
	YOffset  int
	Label    string
	BorderFg Attribute
	BorderBg Attribute
	LabelFg  Attribute
	LabelBg  Attribute
	Bg       Attribute
	Fg       Attribute
}

// NewBlock returns a *Block which inherits styles from current theme.
func NewBlock() *Block {
	return &Block{
		Fg:       Theme.Fg,
		Bg:       Theme.Bg,
		BorderFg: Theme.BorderFg,
		BorderBg: Theme.BorderBg,
		LabelFg:  Theme.LabelFg,
		LabelBg:  Theme.LabelBg,
	}
}

// Buffer draws a box border.
func (b *Block) drawBorder(buf *Buffer) {
	x := b.X + 1
	y := b.Y + 1

	// draw lines
	buf.Merge(NewFilledBuffer(0, 0, x, 1, Cell{HORIZONTAL_LINE, b.BorderFg, b.BorderBg}))
	buf.Merge(NewFilledBuffer(0, y, x, y+1, Cell{HORIZONTAL_LINE, b.BorderFg, b.BorderBg}))
	buf.Merge(NewFilledBuffer(0, 0, 1, y+1, Cell{VERTICAL_LINE, b.BorderFg, b.BorderBg}))
	buf.Merge(NewFilledBuffer(x, 0, x+1, y+1, Cell{VERTICAL_LINE, b.BorderFg, b.BorderBg}))

	// draw corners
	buf.SetCell(0, 0, Cell{TOP_LEFT, b.BorderFg, b.BorderBg})
	buf.SetCell(x, 0, Cell{TOP_RIGHT, b.BorderFg, b.BorderBg})
	buf.SetCell(0, y, Cell{BOTTOM_LEFT, b.BorderFg, b.BorderBg})
	buf.SetCell(x, y, Cell{BOTTOM_RIGHT, b.BorderFg, b.BorderBg})
}

func (b *Block) drawLabel(buf *Buffer) {
	r := MaxString(b.Label, (b.X-3)-1)
	buf.SetString(3, 0, r, b.LabelFg, b.LabelBg)
	if b.Label == "" {
		return
	}
	c := Cell{' ', b.Fg, b.Bg}
	buf.SetCell(2, 0, c)
	if len(b.Label)+3 < b.X {
		buf.SetCell(len(b.Label)+3, 0, c)
	} else {
		buf.SetCell(b.X-1, 0, c)
	}
}

// Resize computes Height, Width, XOffset, and YOffset
func (b *Block) Resize(termWidth, termHeight, termCols, termRows int) {
	b.X = int((float64(b.Grid.Dx())/float64(termCols))*float64(termWidth)) - 2
	b.Y = int((float64(b.Grid.Dy())/float64(termRows))*float64(termHeight)) - 2
	b.XOffset = int(((float64(b.Grid.Min.X) / float64(termCols)) * float64(termWidth)))
	b.YOffset = int(((float64(b.Grid.Min.Y) / float64(termRows)) * float64(termHeight)))
}

func (b *Block) SetGrid(c0, r0, c1, r1 int) {
	b.Grid = image.Rect(c0, r0, c1, r1)
}

func (b *Block) GetXOffset() int {
	return b.XOffset
}

func (b *Block) GetYOffset() int {
	return b.YOffset
}

// Buffer implements Bufferer interface and draws background and border
func (b *Block) Buffer() *Buffer {
	buf := NewBuffer()
	buf.SetAreaXY(b.X+2, b.Y+2)
	buf.Fill(Cell{' ', ColorDefault, b.Bg})

	b.drawBorder(buf)
	b.drawLabel(buf)

	return buf
}