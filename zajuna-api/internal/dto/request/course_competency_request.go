package request

type CreateCourseCompetencyRequest struct {
	CourseID     uint `json:"courseid" binding:"required"`
	CompetencyID uint `json:"competencyid" binding:"required"`
}

const (
	OUTCOME_NONE      = 0
	OUTCOME_EVIDENCE  = 1
	OUTCOME_RECOMMEND = 2
	OUTCOME_COMPLETE  = 3
)
