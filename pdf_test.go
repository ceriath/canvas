package canvas

import (
	"bytes"
	"image"
	"testing"

	"github.com/tdewolff/test"
)

func TestPDF(t *testing.T) {
	c := New(10, 10)
	c.DrawPath(0, 0, MustParseSVG("L10 0"))

	pdfCompress = false
	buf := &bytes.Buffer{}
	c.WritePDF(buf)
	test.T(t, buf.String(), `%PDF-1.7
1 0 obj
<< /Length 14 >> stream
0 0 m 10 0 l f
endstream
endobj
2 0 obj
<< /Type /Page /Contents 1 0 R /Group << /Type /Group /CS /DeviceRGB /I true /S /Transparency >> /MediaBox [0 0 10 10] /Parent 2 0 R /Resources << >> >>
endobj
3 0 obj
<< /Type /Pages /Count 1 /Kids [2 0 R] >>
endobj
4 0 obj
<< /Type /Catalog /Pages 3 0 R >>
endobj
xref
0 5
0000000000 65535 f
0000000009 00000 n
0000000073 00000 n
0000000241 00000 n
0000000298 00000 n
trailer
<< /Root 4 0 R /Size 4 >>
starxref
347
%%EOF`)
}

func TestPDFPath(t *testing.T) {
	buf := &bytes.Buffer{}
	pdf := newPDFWriter(buf).NewPage(210.0, 297.0)
	pdf.SetAlpha(0.5)
	pdf.SetFillColor(Red)
	pdf.SetStrokeColor(Blue)
	pdf.SetLineWidth(5.0)
	pdf.SetLineCap(RoundCapper)
	pdf.SetLineJoin(RoundJoiner)
	pdf.SetDashes(2.0, []float64{1.0, 2.0, 3.0})
	test.String(t, pdf.String(), " /A0 gs 1 0 0 rg /A1 gs 0 0 1 RG 5 w 1 J 1 j [1 2 3 1 2 3] 2 d")
}

func TestPDFText(t *testing.T) {
	dejaVuSerif := NewFontFamily("dejavu-serif")
	dejaVuSerif.LoadFontFile("./test/DejaVuSerif.ttf", FontRegular)

	ebGaramond := NewFontFamily("eb-garamond")
	ebGaramond.LoadFontFile("./test/EBGaramond12-Regular.otf", FontRegular)

	dejaVu8 := dejaVuSerif.Face(8.0*ptPerMm, Black, FontRegular, FontNormal)
	dejaVu12 := dejaVuSerif.Face(12.0*ptPerMm, Red, FontItalic, FontNormal, FontUnderline)
	dejaVu12sub := dejaVuSerif.Face(12.0*ptPerMm, Black, FontRegular, FontSubscript)
	garamond10 := ebGaramond.Face(10.0*ptPerMm, Black, FontBold, FontNormal)

	rt := NewRichText()
	rt.Add(dejaVu8, "dejaVu8")
	rt.Add(dejaVu12, " glyphspacing")
	rt.Add(dejaVu12sub, " dejaVu12sub")
	rt.Add(garamond10, " garamond10")
	text := rt.ToText(dejaVu12.TextWidth("glyphspacing")+float64(len("glyphspacing")-1), 100.0, Justify, Top, 0.0, 0.0)

	buf := &bytes.Buffer{}
	pdf := newPDFWriter(buf).NewPage(210.0, 297.0)
	text.WritePDF(pdf, Identity) // this actually gives coverage to PDF font embedding, which we don't test...
	test.String(t, pdf.String(), " BT /F0 8 Tf 0 -7.421875 Td[(\x00G\x00H\x00M\x00D\x009) 63 (\x00X\x00\x1B)]TJ 1 0 0 rg 1 0 .3 1 0 -20.453125 Tm 1 Tc[(\x00J\x00O\x00\\\x00S\x00K\x00V\x00S\x00D\x00F\x00L\x00Q\x00J)]TJ 0 g 1 0 0 1 0 -29.765625 Tm 0 Tc 2 Tr .27984 w[(\x00G\x00H\x00M\x00D\x009) 63 (\x00X\x00\x14\x00\x15\x00V\x00X\x00E)]TJ /F1 10 Tf 0 -8.734375 Td .4 w[(\x00H\x00B\x00S\x00B\x00N\x00P\x00O\x00E\x00\x12\x00\x11)]TJ ET 1 0 0 rg 0 -22.703125 m 91.71875 -22.703125 l 91.71875 -21.803125 l 0 -21.803125 l 0 -22.703125 l f")
}

func TestPDFImage(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 2, 2))

	buf := &bytes.Buffer{}
	pdf := newPDFWriter(buf).NewPage(210.0, 297.0)
	pdf.DrawImage(img, Lossless, Identity)
	test.String(t, pdf.String(), " q 2 0 0 2 0 0 cm /Im0 Do Q")
}
