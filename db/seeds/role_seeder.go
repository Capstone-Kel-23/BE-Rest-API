package seeds

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/role/utils"
)

func (s Seed) RoleSeed() {
	s.db.FirstOrCreate(&domain.Role{}, domain.Role{Name: utils.Admin.String(), ID: 1})
	s.db.FirstOrCreate(&domain.Role{}, domain.Role{Name: utils.User.String(), ID: 2})
}
