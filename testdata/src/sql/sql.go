package http_request

import (
	"context"
	"database/sql"
)

func _() {
	ctx := context.Background()

	db, _ := sql.Open("noctx", "noctx://")

	// database/sql.DB methods

	db.Begin() // want `\(\*database/sql\.DB\)\.Begin must not be called. use \(\*database/sql\.DB\)\.BeginTx`
	db.BeginTx(ctx, nil)

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

	// database/sql.Stmt methods
	stmt, _ := db.PrepareContext(context.Background(), "select * from testdata where id = ?")

	stmt.Query("1") // want `\(\*database/sql\.Stmt\)\.Query must not be called. use \(\*database/sql\.Conn\)\.QueryContext`
	stmt.QueryContext(ctx, "1")

	stmt.QueryRow("1") // want `\(\*database/sql\.Stmt\)\.QueryRow must not be called. use \(\*database/sql\.Conn\)\.QueryRowContext`
	stmt.QueryRowContext(ctx, "1")

	stmt.Exec("1") // want `\(\*database/sql\.Stmt\)\.Exec must not be called. use \(\*database/sql\.Conn\)\.ExecContext`
	stmt.ExecContext(ctx, "1")

	// database/sql.Tx methods
	tx, _ := db.BeginTx(ctx, nil)
	tx.Exec("select * from testdata") // want `\(\*database/sql\.Tx\)\.Exec must not be called. use \(\*database/sql\.Tx\)\.ExecContext`
	tx.ExecContext(ctx, "select * from testdata")

	tx.Prepare("select * from testdata") // want `\(\*database/sql\.Tx\)\.Prepare must not be called. use \(\*database/sql\.Tx\)\.PrepareContext`
	tx.PrepareContext(ctx, "select * from testdata")

	tx.Query("select * from testdata") // want `\(\*database/sql\.Tx\)\.Query must not be called. use \(\*database/sql\.Tx\)\.QueryContext`
	tx.QueryContext(ctx, "select * from testdata")

	tx.QueryRow("select * from testdata") // want `\(\*database/sql\.Tx\)\.QueryRow must not be called. use \(\*database/sql\.Tx\)\.QueryRowContext`
	tx.QueryRowContext(ctx, "select * from testdata")

	tx.Stmt(stmt) // want `\(\*database/sql\.Tx\)\.Stmt must not be called. use \(\*database/sql\.Tx\)\.StmtContext`
	tx.StmtContext(ctx, stmt)

	_ = tx.Commit()

	// database/sql.Conn are safe, they only have context-aware methods
	// these lines are just to show that they are not flagged
	conn, _ := db.Conn(ctx)
	conn.ExecContext(ctx, "select * from testdata")
	conn.QueryContext(ctx, "select * from testdata")
	conn.QueryRowContext(ctx, "select * from testdata")
}
