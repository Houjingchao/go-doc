package sql

//create by littleM
// 2017-6-27
import (
	"github.com/jmoiron/sqlx"
	"database/sql"
)

type Student struct {
	Id   int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Sex  string `db:"sex" json:"sex"`
}

type StudentDao struct {
	ColId   string `sm:"-"`
	ColName string `sm:"-"`
	ColSex  string `sm:"-"`
	Columns []string `sm:"-"`
	Table   string `sm:"-"`
	DB      *sqlx.DB `sm:"@.(.)"`
}

func NewStudentDao(db *sqlx.DB) *StudentDao {
	dao := &StudentDao{DB: db}
	dao.Init()
	return dao
}

func (dao *StudentDao) Init() {
	dao.Columns = []string{
		"id",
		"name",
		"sex",
	}
	dao.ColId = "id"
	dao.ColName = "name"
	dao.ColSex = "sex"
	dao.Table = "student"
}

func (dao *StudentDao) Fields() string {
	return "`id` ,`name` ,`sex` "
}

func (dao *StudentDao) FindByID(ID int) (one *Student, err error) {
	one = &Student{}
	err = dao.DB.Get(one, "select "+dao.Fields()+" from `student` where `id`=? limit 1", ID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (dao *StudentDao) Page(page, pageSize int) (list []*Student, err error) {
	list = []*Student{}
	err = dao.DB.Select(&list, "select "+dao.Fields()+" from `student` limit ?,?", pageSize*(page-1), pageSize)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (dao *StudentDao) WherePage(whereSql string, page, pageSize int, args ...interface{}) (list []*Student, err error) {
	list = []*Student{}
	args = append(args, pageSize*(page-1), pageSize)
	err = dao.DB.Select(list, "select "+dao.Fields()+" from `student` where "+whereSql+" limit ?,?", args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return list, nil
}
