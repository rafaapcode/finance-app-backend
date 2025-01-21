package model

type User struct {
	Id         string       `json:"id"`
	Nome       string       `json:"nome"`
	Email      string       `json:"email"`
	PhotoUrl   string       `json:"photoUrl"`
	Income     []Income     `json:"income"`
	Outcome    []Outcome    `json:"outcome"`
	Investment []Investment `json:"investment"`
	Goals      []Goals      `json:"goals"`
}
