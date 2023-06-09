package markepplacerepo

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"ssr/internal/domain"
)

type MysqlRepository struct {
	client *sql.DB
}

type MysqlConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

const NftTable = "nft"
const UserTable = "user"

var _ domain.MarketplaceRepository = (*MysqlRepository)(nil)

func NewRepository(config MysqlConfig) MysqlRepository {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(localhost:33066)/%v", config.User, config.Password, config.DbName))
	if err != nil {
		panic(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("cannot connect to database")
		panic(err)
	}
	fmt.Println("database connected!")

	return MysqlRepository{
		client: db,
	}
}

func (m *MysqlRepository) findOne(ctx context.Context, id, table string) *sql.Row {
	query := fmt.Sprintf("select * from %s where id = '%s'", table, id)
	res := m.client.QueryRow(query)
	return res
}

func (m *MysqlRepository) findMany(ctx context.Context, ids, table string) (*sql.Rows, error) {
	query := fmt.Sprintf("select * from %s where id in (%s)", table, ids)
	res, err := m.client.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *MysqlRepository) findAll(ctx context.Context, table string) (*sql.Rows, error) {
	query := fmt.Sprintf("select * from %s", table)
	res, err := m.client.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *MysqlRepository) findAllWithPagination(ctx context.Context, next, took *int, table string) (*sql.Rows, error) {
	query := fmt.Sprintf("select * from %s", table)
	if next != nil {
		query = fmt.Sprintf("%v limit %v", query, *next)
		if took != nil {
			query = fmt.Sprintf("%v offset %v", query, *took)
		}
	}
	res, err := m.client.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *MysqlRepository) getCount(ctx context.Context, table string) (*sql.Rows, error) {
	query := fmt.Sprintf("select count(*) from %s", table)
	res, err := m.client.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}
