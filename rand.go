package golib

type Rand struct {
	seed int64
}

func (r *Rand) Max() int64 {
	return 233280
}

func (r *Rand) Seed(seed int64) {
	r.seed = seed
}

func (r *Rand) Next() int64 {
	r.seed = (r.seed*9301 + 49297) % 233280
	return r.seed
}

func (r *Rand) Float() float64 {
	return float64(r.Next()) / 233280
}

func New(seed int64) *Rand {
	return &Rand{seed}
}
