package userModel

import cnst "nta-blog/internal/constant"

type PrivateUser struct {
	Id    string            `json:"id"`
	Name  string            `json:"name"`
	Email string            `json:"email"`
	Role  cnst.TRoleAccount `json:"role"`
	Avt   string            `json:"avt"`
}
