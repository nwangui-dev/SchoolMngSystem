// db/tables.go
package db

import "github.com/nwangui-dev/SchoolMngSystem/models"

var Tables = []interface{}{
	&models.Student{},
	&models.User{},
}