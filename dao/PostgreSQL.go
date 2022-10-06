package dao

import (
	"FIFA/core"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init 连接到pgsql
func Init() (err error) {
	pgDsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		core.Conf.PgSQLConfig.Host, core.Conf.PgSQLConfig.Port, core.Conf.PgSQLConfig.User,
		core.Conf.PgSQLConfig.Password, core.Conf.PgSQLConfig.Database)
	DB, err = gorm.Open(postgres.Open(pgDsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("Postgres connect failed", zap.Error(err))
		return
		
	}
	//DB.AutoMigrate(
	//	&model.Admin{},
	//	&model.AwardPlayer{},
	//	&model.Match{},
	//	&model.Player{},
	//	&model.Team{},
	//	&model.Group{},
	//	&model.TeamStatisticWithDisciplinary{},
	//	&model.TeamStatisticWithAttempts{},
	//	&model.PlayersStatisticWithDisciplinary{},
	//	&model.PlayersStatisticWithGoalsScored{},
	//	&model.TeamStatisticWithTopGoal{},
	//)
	return nil
}
