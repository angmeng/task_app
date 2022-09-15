package stores

import (
	"fmt"

	"github.com/angmeng/task_app/config"
	"github.com/angmeng/task_app/pkg/helpers"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

func ConnectPG() db.Session {
	switch helpers.GetAppEnv() {
	case "development", "test":
		db.LC().SetLevel(db.LogLevelDebug)
	case "production":
		db.LC().SetLevel(db.LogLevelError)
	}

	sess, err := postgresql.Open(postgresql.ConnectionURL{
		Database: config.Val.DatabaseName,
		Host:     config.Val.DatabaseHost,
		User:     config.Val.DatabaseUser,
		Password: config.Val.DatabasePassword,
	})

	if err != nil {
		panic(fmt.Errorf("error open db: %v", err))
	}

	return sess
}
