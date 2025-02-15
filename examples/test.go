package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/tdewolff/canvas"
)

var dejaVuSerif *canvas.FontFamily

func main() {
	dejaVuSerif = canvas.NewFontFamily("dejavu-serif")
	dejaVuSerif.Use(canvas.CommonLigatures)
	if err := dejaVuSerif.LoadLocalFont("Amiri", canvas.FontRegular); err != nil {
		panic(err)
	}

	c := canvas.New(220, 100)
	draw(c)
	c.Fit(1.0)

	////////////////

	svgFile, err := os.Create("test.svg")
	if err != nil {
		panic(err)
	}
	defer svgFile.Close()
	c.WriteSVG(svgFile)

	////////////////

	pngFile, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}
	defer pngFile.Close()

	img := c.WriteImage(10.0)
	err = png.Encode(pngFile, img)
	if err != nil {
		panic(err)
	}

	pdfFile, err := os.Create("test.pdf")
	if err != nil {
		panic(err)
	}
	defer pdfFile.Close()

	err = c.WritePDF(pdfFile)
	if err != nil {
		panic(err)
	}
}

func drawStrokedPath(c *canvas.Canvas, x, y, d float64, path string) {
	fmt.Println("----------")
	c.SetFillColor(canvas.Black)
	p, err := canvas.ParseSVG(path)
	if err != nil {
		panic(err)
	}
	c.DrawPath(x, y, p)

	p = p.Stroke(d, canvas.ButtCapper, canvas.MiterClipJoiner(canvas.RoundJoiner, d))
	c.SetFillColor(color.RGBA{128, 0, 0, 128})
	c.DrawPath(x, y, p)

	p = p.Stroke(0.2, canvas.RoundCapper, canvas.RoundJoiner)
	c.SetFillColor(color.RGBA{255, 0, 0, 255})
	c.DrawPath(x, y, p)
}

func draw(c *canvas.Canvas) {
	c.SetFillColor(color.RGBA{0, 0, 128, 128})
	p := canvas.Rectangle(4.0, 2.0)
	c.DrawPath(10.0, 10.0, p)
	c.DrawPath(10.0, 10.0, p.Transform(canvas.Identity.Shear(2.0, -1.0)))
	//s := "ﻙﺎﻨﺗ ﺎﻠﺴﻣﺍﺀ ﺹ \n\n ﺎﻔﻳﺓ ﻢﻧ"

	//face := dejaVuSerif.Face(12.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	//frame := canvas.Rectangle(28.0, -10.0)
	//rt := canvas.NewRichText()
	//rt.Add(face, s)
	//text := rt.ToText(28.0, 0.0, canvas.Justify, canvas.Top, 0.0, 0.0)

	//p, _ := canvas.ParseSVG(fmt.Sprintf("A20.0 10.0 0.0 0 0 40.0 0.0z"))
	//f := p.Flatten().Stroke(1.0, canvas.ButtCapper, canvas.MiterJoiner)
	//c.SetFillColor(color.RGBA{0, 0, 128, 128})

	//c.DrawPath(10.0, 10.0, f)
	//c.DrawPath(10.0, 10.0, frame)
	//c.DrawText(10.0, 10.0, text)

	//c.SetView(canvas.Identity.Rotate(45).Scale(2.0, 1.0))

	//c.DrawPath(20.0, 20.0, f)
	//c.DrawPath(20.0, 20.0, frame)
	//c.DrawText(20.0, 20.0, text)

	//c.SetFillColor(color.RGBA{0, 128, 0, 128})
	//for _, marker := range p.Markers(canvas.Rectangle(-2, -2, 4, 4), canvas.Circle(2), canvas.RegularPolygon(6, 2, true), true) {
	//	c.DrawPath(10.0, 10.0, marker)
	//}

	//p = p.Stroke(1.0, canvas.ButtCapper, canvas.MiterJoiner)
	//c.SetFillColor(color.RGBA{128, 0, 0, 128})
	//c.DrawPath(10.0, 10.0, p)

	//p, _ = canvas.ParseSVG(fmt.Sprintf("A60.0 30.0 45.0 1 0 80.0 0.0"))
	//f = p.Flatten().Stroke(1.0, canvas.ButtCapper, canvas.MiterJoiner)
	//c.SetFillColor(color.RGBA{0, 0, 128, 128})
	//c.DrawPath(30.0, 10.0, f)

	//p = p.Stroke(1.0, canvas.ButtCapper, canvas.MiterJoiner)
	//c.SetFillColor(color.RGBA{128, 0, 0, 128})
	//c.DrawPath(30.0, 10.0, p)

	// test Filling
	//canvas.FillRule = canvas.EvenOdd
	//p, _ := canvas.ParseSVG(fmt.Sprintf("M100 0V100H-100V-100H100zM50 0V-50H-50V50H50z"))
	//p, _ = canvas.ParseSVG(fmt.Sprintf("M100 0V100H-100V-100H100zM50 0V50H-50V-50H50z"))
	//fmt.Println(p.Filling())
	//p = p.Offset(1.0)
	//c.DrawPath(110, 110, p)

	//ellipse, _ := canvas.ParseSVG(fmt.Sprintf("A20 40 0 0 0 40 0z"))
	//c.SetFillColor(canvas.Red)
	//c.DrawPath(10.0, 10.0, ellipse)

	//ellipse = ellipse.Flatten()
	//ps := ellipse.SplitAt(10.0)
	//ellipse = ps[0]
	//fmt.Println(ellipse)
	//ellipse = ellipse.Stroke(2.0, canvas.RoundCapper, canvas.RoundJoiner)
	//c.SetFillColor(color.RGBA{0, 0, 0, 128})
	//c.DrawPath(10.0, 10.0, ellipse)

	//drawStrokedPath(c, 30, 40, 2.0, "M0 0L50 0L50 -5")
	//drawStrokedPath(c, 30, 30, 2.0, "M-25 -25A25 25 0 0 1 0 0A25 25 0 0 1 25 -25z")
	//drawStrokedPath(c, 80, 30, 2.0, "M-35.35 -14.65A50 50 0 0 0 0 0A50 50 0 0 0 35.35 -14.65L-35.35 -14.65z")
	//drawStrokedPath(c, 140, 35, 2.0, "M-25 -30A50 50 0 0 1 0 0A50 50 0 0 1 25 -30L-25 -30z")
	//drawStrokedPath(c, 30, 70, 2.0, "M0 -25A25 25 0 0 1 0 0A25 25 0 0 1 0 -25z") // CCW
	//drawStrokedPath(c, 60, 70, 2.0, "M0 -25A25 25 0 0 0 0 0A25 25 0 0 0 0 -25z") // CW
	//drawStrokedPath(c, 90, 65, 2.0, "M0 -25A25 25 0 0 0 0 0A25 25 0 0 1 20 -25z")
	//drawStrokedPath(c, 140, 50, 4.0, "M0 0A20 20 0 0 0 40 0A10 10 0 0 1 20 0z")
	//drawStrokedPath(c, 170, 20, 2.0, "C10 -13.33 10 -13.33 20 0z")
	//drawStrokedPath(c, 170, 30, 2.0, "C10 13.33 10 13.33 20 0z")

	// c.SetColor(canvas.LightGrey)
	// c.DrawPath(20.0, 20.0, 0.0, canvas.Rectangle(0.0, 0.0, 160.0, 40.0))
	// c.SetColor(canvas.Black)

	// ff := dejaVuSerif.Face(8.0)
	// text := canvas.NewTextBox(ff, "Lorem ipsum dolor sid amet, confiscusar patria est gravus repara sid ipsum. Apare tu garage.", 160.0, 40.0, canvas.Justify, canvas.Justify, 140.0)
	// c.DrawText(20, 60-ff.Metrics().Ascent, 0.0, text)

	//p, _ := canvas.ParseSVG("C0 10 10 10 100 0")
	//ps := p.SplitAt(52.5)
	//c.DrawPath(110, 20, 0, ps[0])
	//c.DrawPath(110, 20, 0, ps[1])
	//c.DrawPath(110, 30, 0, p.Dash(1.0, 1.0).Stroke(1.0, canvas.ButtCapper, canvas.RoundJoiner))

	//p, _ = canvas.ParseSVG("Q0 10 100 0")
	//ps = p.SplitAt(52.5)
	//c.DrawPath(110, 50, 0, ps[0])
	//c.DrawPath(110, 50, 0, ps[1])
	//c.DrawPath(110, 60, 0, p.Dash(1.0, 1.0).Stroke(1.0, canvas.ButtCapper, canvas.RoundJoiner))

	//p, _ = canvas.ParseSVG("A100 10 0 0 1 -100 10")
	//ps = p.SplitAt(52.5)
	//c.DrawPath(105, 40, 0, ps[0])
	//c.DrawPath(105, 40, 0, ps[1])
	//c.DrawPath(105, 50, 0, p.Dash(1.0, 1.0).Stroke(1.0, canvas.ButtCapper, canvas.RoundJoiner))
}
