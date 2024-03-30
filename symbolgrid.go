package maxicode

import "github.com/fogleman/gg"

type SymbolGrid [30 * 33]bool

func (s *SymbolGrid) SetModule(row, column int, value bool) {
	s[30*row+column] = value
}

func (s *SymbolGrid) GetModule(row, column int) bool {
	return s[30*row+column]
}

func (s *SymbolGrid) Draw(multiplier float64) *gg.Context {
	centerX := 13.64 * multiplier
	centerY := 13.43 * multiplier

	innerRadius := 0.85 * multiplier
	centerRadius := 2.20 * multiplier
	outerRadius := 3.54 * multiplier

	dc := gg.NewContext(int(28*multiplier), int(26.8*multiplier))

	// Central bullseye patterns.
	dc.SetLineWidth(0.67 * multiplier)

	dc.DrawCircle(centerX, centerY, outerRadius)
	dc.SetRGB(0, 0, 0)
	dc.Stroke()

	dc.DrawCircle(centerX, centerY, centerRadius)
	dc.SetRGB(0, 0, 0)
	dc.Stroke()

	dc.DrawCircle(centerX, centerY, innerRadius)
	dc.SetRGB(0, 0, 0)
	dc.Stroke()

	// Hexagons
	for row := 0; row < 33; row++ {
		for column := 0; column < 30; column++ {
			if s.GetModule(row, column) {
				rowOffset := 0.88
				if (row & 1) == 1 {
					rowOffset = 1.32
				}

				hexRectX := (float64(column)*0.88 + rowOffset) * multiplier
				hexRectY := (float64(row)*0.76 + 0.76) * multiplier
				hexRectW := 0.76 * multiplier
				hexRectH := 0.88 * multiplier

				dc.MoveTo(hexRectX+hexRectW*0.5, hexRectY)
				dc.LineTo(hexRectX+hexRectW, hexRectY+hexRectH*0.25)
				dc.LineTo(hexRectX+hexRectW, hexRectY+hexRectH*0.75)
				dc.LineTo(hexRectX+hexRectW*0.5, hexRectY+hexRectH)
				dc.LineTo(hexRectX, hexRectY+hexRectH*0.75)
				dc.LineTo(hexRectX, hexRectY+hexRectH*0.25)
				dc.Fill()
			}
		}
	}

	return dc
}

func (s *SymbolGrid) SaveToPNG(multiplier float64, path string) error {
	return s.Draw(multiplier).SavePNG(path)
}
