package order

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"

	"core-data/business"
	orderBusiness "core-data/business/order"
)

type MySQL struct {
	db *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) *MySQL {
	return &MySQL{
		db,
	}
}

func (repo *MySQL) InsertOrder(insertSpec orderBusiness.OrderSpec) error {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return err
	}
	insertQuery := `INSERT INTO orders (
		teknisi_id,
		kerusakan_id,
		jenis_hp,
		jenis_platform,
		status_service,
		lama_pengerjaan,
		version,
		version_teknisi)
	VALUES
	(?,?,?,?,?,?,?,?)`

	_, err = tx.Exec(insertQuery, insertSpec.TeknisiID, insertSpec.KerusakanID, insertSpec.JenisHP, insertSpec.JenisPlatform, insertSpec.StatusService, insertSpec.LamaPengerjaan, 1, insertSpec.VersionTeknisi+1)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New("error duplicate entry data")
			return err
		}

		err = errors.New("resource error")
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return err
	}

	return nil
}

func (repo *MySQL) UpdateOrderStatus(updateSpec orderBusiness.UpdateOrderStatus) (bool, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	insertQuery := `UPDATE orders 
		SET
		status_service = ?,
		version = ?
		WHERE
		id = ?`

	_, err = tx.Exec(insertQuery, updateSpec.StatusService, updateSpec.Version+1, updateSpec.ID)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New("error duplicate entry data")
			return false, err
		}

		err = errors.New("resource error")
		return false, err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return false, err
	}

	return true, nil
}

func (repo *MySQL) FindOrderByID(id int) (*orderBusiness.Order, error) {
	var order orderBusiness.Order

	selectQuery := `SELECT * FROM orders WHERE id = ?`

	err := repo.db.QueryRowx(selectQuery, id).StructScan(&order)
	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, business.ErrNotFound
		}
		err = errors.New("resource error")
		return nil, err
	}

	return &order, nil
}
