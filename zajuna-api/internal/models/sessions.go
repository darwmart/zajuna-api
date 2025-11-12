package models

type Sessions struct {
	ID           uint    `gorm:"column:id;primaryKey" json:"id"`
	State        int64   `gorm:"column:state" json:"state"`
	SID          string  `gorm:"column:sid" json:"sid"`
	UserID       uint    `gorm:"column:userid" json:"userid"`
	SessData     *string `gorm:"column:sessdata" json:"sessdata"`
	TimeCreated  int64   `gorm:"column:timecreated" json:"timecreated"`
	TimeModified int64   `gorm:"column:timemodified" json:"timemodified"`
	FirstIp      string  `gorm:"column:firstip" json:"firstip"`
	LastIp       string  `gorm:"column:lastip" json:"lastip"`
}

func (Sessions) TableName() string {
	return "mdl_sessions"
}
