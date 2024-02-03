package types

type Repository interface {
	Select(sql string) interface{}
	Update(sql string)
	Delete(sql string)
	Insert(sql string)
}

type SqlRepository struct {
}

func (sql *SqlRepository) Select(sqlTxt string) interface{} {
	return "Query executed successfully"
}

func (sql *SqlRepository) Update(sqlText string) {

}

func (sql *SqlRepository) Delete(sqlText string) {

}

func (sql *SqlRepository) Insert(sqlText string) {

}
