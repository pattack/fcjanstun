package team

import (
	"github.com/pattack/rcss"
)

type janstun struct {
	name  string
	side  rcss.Side
	match rcss.Match
}

func (j janstun) Name() string {
	return j.name
}
func (j *janstun) SetSide(side rcss.Side) {
	// j.side = side
}
func (j *janstun) Invite(m rcss.Match, unum rcss.UniformNumber) {
	j.match = m
	j.match.Move(0, 0)
}
func (j *janstun) SetPlayMode(mode rcss.PlayMode) {

	// fmt.Printf("Side: %c, Uniform Number: %d, Mode: %s\n", side, unum, mode)
}

func (j *janstun) ServerParam(sp rcss.ServerParameters) {
	//fmt.Printf("Server Parameters: %#v\n", sp)
}

func (j *janstun) PlayerParam(pp rcss.PlayerParameters) {
	//fmt.Printf("Player Parameters: %#v\n", pp)
}

func (j *janstun) PlayerType(pt rcss.PlayerType) {
	//fmt.Printf("Player Type: %#v\n", pt)
}

func (j *janstun) See() {

}

func (j *janstun) Hear() {

}

func (j *janstun) SenseBody() {

}

func (j *janstun) Score() {

}

func (j *janstun) Kickoff() {

}

func NewTeam(name string) rcss.Team {
	return &janstun{
		name: name,
	}
}
