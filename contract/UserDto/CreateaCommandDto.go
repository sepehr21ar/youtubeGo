package userdto

type UserCommandDto struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
}
