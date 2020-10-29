package tipomagnitud

type TipoMagnitud int

const (
	Md = 1 + iota
	MbLg
	mb
	mbLg
	Mw
	M
)

var magnitudes = [...]string{
	"Md (M-MS)",
	"MbLg (M-MS)",
	"mb (V-C)",
	"mbLg (L)",
	"Mw",
	"M(mb)",
}