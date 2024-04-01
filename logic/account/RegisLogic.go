package logicacc

import (
	"context"
	"database/sql"
	"thuchanh_go/banana"
	"thuchanh_go/database"
	"thuchanh_go/logic"
	"thuchanh_go/models"
	"thuchanh_go/types/req"

	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
)

type AccRegisterLogic struct {
	sql *database.Sql
}

func NewAccRegisterLogic(sql *database.Sql) logic.AccLogic {
	return &AccRegisterLogic{
		sql: sql,
	}
}
func (a *AccRegisterLogic) Insert(ctx context.Context, user models.Account) (models.Account, error) {
	statement := `INSERT INTO users (id, name, phone, email, username, password)
				  VALUES(:id, :name, :phone, :email, :username, :password)`

	_, err := a.sql.Db.NamedExecContext(ctx, statement, user)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		return user, banana.RegisFail
	}

	return user, nil

}
func (a *AccRegisterLogic) Select(ctx context.Context, req req.UserLoginReq) (models.Account, error) {
	var user = models.Account{}
	statement := "SELECT * FROM users WHERE username=$1"
	err := a.sql.Db.GetContext(ctx, &user, statement, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFoud
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}
