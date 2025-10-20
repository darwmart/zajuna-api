package models

type Sessions struct {
	ID           uint   `gorm:"column:id;primaryKey" json:"id"`
	State        int64  `gorm:"column:state" json:"state"`
	SID          string `gorm:"column:sId" json:"sid"`
	UserID       int    `gorm:"column:userId" json:"userid"`
	SessData     string `gorm:"column:sessdata" json:"sessdata"`
	TimeCreated  int64  `gorm:"column:timecreated" json:"timecreated"`
	TimeModified int64  `gorm:"column:timemodified" json:"timemodified"`
	FirstIp      int    `gorm:"column:firstip" json:"firstip"`
	LastIp       int    `gorm:"column:lastip" json:"lastip"`
}

func (Sessions) TableName() string {
	return "mdl_sessions"
}
