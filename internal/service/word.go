package service

import (
	time
	json
)

type Word struct {
	Id                 string    `json:"id"`
	Term               string    `json:"term"`
	PersonalDefinition string    `json:"personalDefinition"`
	OfficialDefinition string    `json:"officialDefinition"`
	Examples           []string  `json:"examples"`
	DateAdded          time.Time `json:"dateAdded"`
	DateReviewed       time.Time `json:"dateReviewed"`
	Understanding      int       `json:"Understanding"`
};

