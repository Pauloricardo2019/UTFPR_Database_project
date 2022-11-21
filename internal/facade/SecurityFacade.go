package facade

import "utfpr_db/dto"

type SecurityFacade struct {
}

func NewSecurityFacade() *SecurityFacade {
	return &SecurityFacade{}
}

func (s *SecurityFacade) ValidateToken(*dto.ValidateTokenRequest) (*dto.ValidateTokenResponse, error) {
	return nil, nil
}
