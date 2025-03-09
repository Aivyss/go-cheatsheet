package types

type Pair[F, S any] struct {
	first  F
	second S
}

func (p Pair[F, S]) First() F {
	return p.first
}

func (p Pair[F, S]) Second() S {
	return p.second
}

func PairOf[F, S any](first F, second S) Pair[F, S] {
	return Pair[F, S]{
		first:  first,
		second: second,
	}
}
