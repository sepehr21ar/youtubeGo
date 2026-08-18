package userdto

type UserCommandDto struct {
	Name   string `json:"name" binding:"required"`
	Family string `json:"family" binding:"required"`
}
