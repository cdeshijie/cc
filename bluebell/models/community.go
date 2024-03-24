package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

// omitempty的意思是字段为空就不展示了
type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
