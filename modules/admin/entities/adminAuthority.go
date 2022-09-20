package entities

type AdminAuthority struct {
	AdminId              uint   `gorm:"column:admin_id"`
	AuthorityAuthorityId string `gorm:"column:authority_authority_id"`
}

func (s *AdminAuthority) TableName() string {
	return "admin_authority"
}
