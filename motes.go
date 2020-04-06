package main

const (
	F Mote = iota // zero
	A
	B
	C
	D
	E
	G      // control char
	GSharp // separator
)

type (
	Mote byte
	Pack []Mote
)

func (m Mote) Freq() float64 {
	switch m {
	case A:
		return 440.0
	case B:
		return 493.883
	case C:
		return 523.251
	case D:
		return 587.329
	case E:
		return 659.255
	case F:
		return 698.456
	case G:
		return 783.990
	case GSharp:
		return 830.609
	}
	return 0
}

func (m Mote) String() string {
	switch m {
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	case F:
		return "F"
	case G:
		return "G"
	case GSharp:
		return "G#"
	}

	return "(#invalid#)"
}

func (p *Pack) Init() {
	(*p) = Pack{G}
}

func (p *Pack) Add(motes ...Mote) {
	last := GSharp
	for _, m := range motes {
		if m == last {
			(*p) = append(*p, GSharp)
			m = GSharp
		} else {
			(*p) = append(*p, m)
		}
		last = m
	}
	(*p) = append(*p, G)
}

func (p *Pack) Wrap() {
	(*p) = append(*p, G)
}

func PackFromBytes(data []byte) Pack {
	var p Pack
	p.Init()
	for _, b := range data {
		p.Add(byte2M(b)...)
	}
	p.Wrap()
	return p
}

func byte2M(number byte) []Mote {
	const base = 5
	var tmp []Mote

	for {
		if number <= base {
			tmp = append(tmp, Mote(number))
			break
		}
		tmp = append(tmp, Mote(number%base))
		number = number / base
	}

	i := 0
	l := len(tmp)
	result := make([]Mote, l)

	for l > 0 {
		l--
		result[i] = tmp[l]
		i++
	}

	return result
}
