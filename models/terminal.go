package models

import "time"

type Terminal struct {
	Id        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Num       int32     `gorm:"BIGINT UNSIGNED; NOT NULL AUTO_INCREMENT UNIQUE" json:"num"`
	Pass      string    `gorm:"varchar(100);not null" json:"pass"`
	Token     string    `gorm:"varchar(100); not null" json:"token"`
	Address   string    `gorm:"varchar(255); not null"json:"address"`
	AgentId   int       `gorm:"varchar(100); not null" json:"agent_id"`
	AgentName string    `gorm:"varchar(255);not null" json:"agent_name"`
	Mcode     string    `gorm:"varchar(100);default ''" json:"mcode"`
	Status    bool      `gorm:"boolean;default false" json:"status"`
	Regdate   time.Time `gorm:"default:current_timestamp()" json:"regdate" `
}

func NewTerminal(terminal Terminal) error {
	db := Connect()
	defer db.Close()
	err := db.Create(&terminal).Error
	return err

}

//id uuid NOT NULL DEFAULT uuid_generate_v4(),
//num SERIAL,
//pass VARCHAR(100) NOT NULL,
//token VARCHAR(64) DEFAULT '',
//info jsonb DEFAULT '{}',
//agentid uuid NOT NULL,
//agentname VARCHAR(255) NOT NULL,
//mcode VARCHAR(100) DEFAULT '',
//status VARCHAR(20) DEFAULT 'new',
//regdate TIMESTAMP DEFAULT now(),
//PRIMARY KEY (id)
//);
//ALTER SEQUENCE terminals_num_seq RESTART WITH 32451 INCREMENT BY 63;
//ALTER TABLE terminals ADD COLUMN local_db_pass VARCHAR(64) DEFAULT 'masterkey';
