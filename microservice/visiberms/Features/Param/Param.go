package visiberms_param

import "database/sql"

type Param struct {
	Postgres *sql.DB
	DBPrefix string
}

func (self *Param) Init() (err error) {

	err = self.DBConnect()

	return
}

func (self *Param) Close() {
	self.Postgres.Close()
}
