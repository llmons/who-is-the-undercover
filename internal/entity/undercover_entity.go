package entity

type UndercoverEntity struct {
	ID        string `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt string `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP created_at"`
	UpdatedAt string `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP updated_at"`

	CivWord string `xorm:"not null default '' VARCHAR(255) civ_word"`
	UcWord  string `xorm:"not null default '' VARCHAR(255) uc_word"`
	Score   int    `xorm:"not null default 0 INT(11) score"`
	Source  string `xorm:"not null default 'manul' VARCHAR(255) source"`
}
