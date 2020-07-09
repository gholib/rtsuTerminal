package models

type Service struct {
	ID               int32  `gorm:"primary_key;auto_increment" json:"id"`
	NameOfService    string `gorm:"varchar(255);not null" json:"name_of_service"`
	Parent           int32  `gorm:"integer; not null" json:"parent"`
	Logo             string `gorm:"varchar(255); not null" json:"logo"`
	Regex            string `gorm:"varchar(255);not null" json:"regex"`
	Top              int32  `gorm:"integer;not null" json:"top"`
	Style            string `gorm:"varchar(255); not null" json:"style"`
	Gate             string `gorm:"varchar(255); not null"json:"gate"`
	Title            string `gorm:"varchar(255); not null" json:"title"`
	Sample           string `gorm:"varchar(255) not null "json:"sample"`
	Limiting         string `gorm:"varchar(255) not null" json:"limiting"`
	Prefix           string `gorm:"varchar(255) not null "json:"prefix"`
	ExtraInfo        string `gorm:"varchar(255) not null "json:"extra_info"`
	Info             bool   `gorm:"boolean;default false" json:"info"`
	Currency         string `gorm:"varchar(10); not null" json:"currency"`
	CanCorrect       bool   `gorm:"boolean;default false" json:"can_correct"`
	CanCancel        bool   `gorm:"boolean;default false " json:"can_cancel"`
	CorrectionPeriod int32  `gorm:"integer; not null" json:"correction_period"`
	Enabled          bool   `gorm:"boolean; default false" json:"enabled"`
}

func NewService(service Service) error {
	db := Connect()
	defer db.Close()
	err := db.Create(&service).Error
	return err

}
