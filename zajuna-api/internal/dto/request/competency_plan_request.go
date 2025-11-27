package request

type CreateCompetencyPlanRequest struct {
	Name              string `json:"name" binding:"required"`
	Description       string `json:"description"`
	DescriptionFormat int16  `json:"descriptionformat" binding:"omitempty,min=1,oneof=0 1 2 4"`
	UserID            uint   `json:"user_id" binding:"required"`
	TemplateID        *uint  `json:"template_id"`
	OrigTemplateID    *uint  `json:"orig_template_id"`
	Status            int16  `json:"status" binding:"omitempty,oneof=0 1 2 3 4"`
	DueDate           int64  `json:"due_date"`
	ReviewerID        *uint  `json:"reviewer_id"`
}

const (
	STATUS_DRAFT              = 0
	STATUS_ACTIVE             = 1
	STATUS_COMPLETE           = 2
	STATUS_WAITING_FOR_REVIEW = 3
	STATUS_IN_REVIEW          = 4
)
