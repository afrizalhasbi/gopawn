package iface

import (
	"database/sql"
)

type Service interface {
	*sql.DB
	DataClass
}
