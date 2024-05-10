package database

import "rakamin-submission/models"

type Migration struct {
	Migration interface{}
}

func MigrationDB() []Migration {
	return []Migration{
		{Migration: models.User{}},
		{Migration: models.Photo{}},
	}
}
