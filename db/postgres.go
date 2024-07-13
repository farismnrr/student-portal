/**
 * Package db provides the database connection and operations for PostgreSQL.
 *
 * Structs:
 * - Postgres: Represents the PostgreSQL database connection.
 *
 * Functions:
 * - Connect: Establishes a connection to the PostgreSQL database using the provided credentials.
 * - NewDB: Constructor for creating a new instance of the Postgres struct.
 */

package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"a21hc3NpZ25tZW50/model"
)

type Postgres struct{}

func (p *Postgres) Connect(creds *model.Credential) (*gorm.DB, error) {
	Host := creds.HostAlternative
	if Host == "" {
		Host = creds.Host
	}

	Username := creds.UsernameAlternative
	if Username == "" {
		Username = creds.Username
	}

	Password := creds.PasswordAlternative
	if Password == "" {
		Password = creds.Password
	}

	DatabaseName := creds.DatabaseNameAlternative
	if DatabaseName == "" {
		DatabaseName = creds.DatabaseName
	}

	Port := creds.PortAlternative
	if Port == 0 {
		Port = creds.Port
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Jakarta", Host, Username, Password, DatabaseName, Port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func NewDB() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Reset(db *gorm.DB, table string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("TRUNCATE " + table).Error; err != nil {
			return err
		}

		if err := tx.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1").Error; err != nil {
			return err
		}

		return nil
	})
}
