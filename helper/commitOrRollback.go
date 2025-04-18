package helper

import (
	"database/sql"
	"fmt"
)

func CommitOrRollback(tx *sql.Tx, err *error) {
	if *err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			fmt.Println(rollbackErr)
		}
	} else {
		commitErr := tx.Commit()
		if commitErr != nil {
			fmt.Println(commitErr)
		}
	}
}
