package services

type RoleCapabilityServiceInterface interface {
	HasCapabilities(capabilities []string, userID uint) (bool, error)
}
