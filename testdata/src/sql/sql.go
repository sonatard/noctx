package http_request

import (
	"context"
	"database/sql"
)

func _() {
	ctx := context.Background()

	db, _ := sql.Open("noctx", "noctx://")

	db.Exec("select * from testdata") // want `\(\*database/sql\.DB\)\.Exec must not be called. use \(\*database/sql\.DB\)\.ExecContext`
	db.ExecContext(ctx, "select * from testdata")

	db.Ping() // want `\(\*database/sql\.DB\)\.Ping must not be called. use \(\*database/sql\.DB\)\.PingContext`
	db.PingContext(ctx)

	db.Prepare("select * from testdata") // want `\(\*database/sql\.DB\)\.Prepare must not be called. use \(\*database/sql\.DB\)\.PrepareContext`
	db.PrepareContext(ctx, "select * from testdata")

	db.Query("select * from testdata") // want `\(\*database/sql\.DB\)\.Query must not be called. use \(\*database/sql\.DB\)\.QueryContext`
	db.QueryContext(ctx, "select * from testdata")

	db.QueryRow("select * from testdata") // want `\(\*database/sql\.DB\)\.QueryRow must not be called. use \(\*database/sql\.DB\)\.QueryRowContext`
	db.QueryRowContext(ctx, "select * from testdata")

	// transactions

	tx, _ := db.Begin()
	tx.Exec("select * from testdata") // want `\(\*database/sql\.Tx\)\.Exec must not be called. use \(\*database/sql\.Tx\)\.ExecContext`
	tx.ExecContext(ctx, "select * from testdata")

	tx.Prepare("select * from testdata") // want `\(\*database/sql\.Tx\)\.Prepare must not be called. use \(\*database/sql\.Tx\)\.PrepareContext`
	tx.PrepareContext(ctx, "select * from testdata")

	tx.Query("select * from testdata") // want `\(\*database/sql\.Tx\)\.Query must not be called. use \(\*database/sql\.Tx\)\.QueryContext`
	tx.QueryContext(ctx, "select * from testdata")

	tx.QueryRow("select * from testdata") // want `\(\*database/sql\.Tx\)\.QueryRow must not be called. use \(\*database/sql\.Tx\)\.QueryRowContext`
	tx.QueryRowContext(ctx, "select * from testdata")

	_ = tx.Commit()
}
