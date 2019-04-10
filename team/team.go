package team

import (
	"github.com/pattack/rcss"
)

type janstun struct {
	name string
	side rcss.Side
	match rcss.Match
}

func (j janstun) Name() string {
	return j.name
}

func (j *janstun) Init(match rcss.Match, side rcss.Side, unum rcss.UniformNumber, mode rcss.PlayMode) {
	j.side = side
	j.match = match

	// fmt.Printf("Uniform Number: %d, Mode: %s\n", unum, mode)
}

func NewTeam(name string) rcss.Team {
	return &janstun{
		name: name,
	}
}
