package response

type CreateCompetencyPlanResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"fullname"`
	Message string `json:"message"`
}
