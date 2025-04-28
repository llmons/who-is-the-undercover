package entity

import "time"

type WordPair struct {
	ID        int64     `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`

	CivWord string  `xorm:"not null default '' VARCHAR(255) civ_word"`
	UcWord  string  `xorm:"not null default '' VARCHAR(255) uc_word"`
	Score   float64 `xorm:"not null default 0 FLOAT score"`
	Source  string  `xorm:"not null default 'manual' VARCHAR(255) source"`
}
