package maxicode

import "github.com/ingridhq/gg"

type SymbolGrid [30 * 33]bool

func (s *SymbolGrid) SetModule(row, column int, value bool) {
	s[30*row+column] = value
}

func (s *SymbolGrid) GetModule(row, column int) bool {
	return s[30*row+column]
}

func (s *SymbolGrid) Draw(dpmm float64) *gg.Context {
	centerX := 13.64 * dpmm
	centerY := 13.43 * dpmm

	innerRadius := 0.85 * dpmm
	centerRadius := 2.20 * dpmm
	outerRadius := 3.54 * dpmm

	dc := gg.NewContext(int(28*dpmm), int(26.8*dpmm))

	// Central bullseye patterns.
	dc.SetLineWidth(0.67 * dpmm)

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
	for row := range 33 {
		for column := range 30 {
			if !s.GetModule(row, column) {
				continue
			}

			rowOffset := 0.88
			if (row & 1) == 1 {
				rowOffset = 1.32
			}

			col := float64(column)
			row := float64(row)

			hexRectX := (float64(col*0.88) + rowOffset) * dpmm
			hexRectY := (float64(row*0.76) + 0.76) * dpmm
			hexRectW := 0.76 * dpmm
			hexRectH := 0.88 * dpmm

			dc.MoveTo(hexRectX+float64(hexRectW*0.5), hexRectY)
			dc.LineTo(hexRectX+hexRectW, hexRectY+float64(hexRectH*0.25))
			dc.LineTo(hexRectX+hexRectW, hexRectY+float64(hexRectH*0.75))
			dc.LineTo(hexRectX+float64(hexRectW*0.5), hexRectY+hexRectH)
			dc.LineTo(hexRectX, hexRectY+float64(hexRectH*0.75))
			dc.LineTo(hexRectX, hexRectY+float64(hexRectH*0.25))
			dc.Fill()
		}
	}

	return dc
}

func (s *SymbolGrid) SaveToPNG(multiplier float64, path string) error {
	return s.Draw(multiplier).SavePNG(path)
}
