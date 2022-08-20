package mariadb

import (
	atomic "sync/atomic"
	time "time"

	sql "database/sql"
	// Mariadb Driver
	_ "github.com/go-sql-driver/mysql"

	sys "interview-test-free-fair/pkg/sys"
)

var (
	db      *sql.DB
	healthy int32
)

// Connect - Connect at Mariadb
func Connect() {
	db = createClient()
	go monitor()
}

// IsAlive - Return Mariadb session status
func IsAlive() bool {
	return atomic.LoadInt32(&healthy) == 1
}

// RetrieveClient - Return a Mariadb session
func RetrieveClient() *sql.DB {
	return db
}

// Disconnect - Disconnect at Mariadb
func Disconnect() {
	if db != nil {
		db.Close()
		sys.LogInfo("[Mariadb session closed]")
	}
}

func createClient() *sql.DB {

	db, err := sql.Open("mysql", sys.Properties.Mariadb)

	if err != nil {
		sys.LogError("[Could not create Mariadb client] err:%+v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		setStatusDown()
		sys.LogWarn("[Could create a Mariadb session] err:%+v", err)
	} else {
		var version string
		db.QueryRow("SELECT VERSION()").Scan(&version)
		sys.LogInfo("[Mariadb connected with version: %s]", version)
		setStatusUp()
	}

	return db
}

func monitor() {
	for {

		time.Sleep(30 * time.Second)

		if db == nil || db.Ping() != nil {
			setStatusDown()
			sys.LogWarn("[Mariadb session is not active, trying to reconnect]")
		} else {
			setStatusUp()
			sys.LogInfo("[Mariadb session it's alive]")
		}
	}
}

func setStatusUp() {
	atomic.StoreInt32(&healthy, 1)
}

func setStatusDown() {
	atomic.StoreInt32(&healthy, 0)
}
