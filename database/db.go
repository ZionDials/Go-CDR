// Copyright (c) 2023 Zion Dials <me@ziondials.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package database

import (
	"fmt"

	"github.com/ziondials/go-cdr/config"
	"github.com/ziondials/go-cdr/logger"
	"github.com/ziondials/go-cdr/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"

	// "gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

// DataService is a wrapper for the GORM database.
type DataService struct {
	Session *gorm.DB
	Config  *config.DatabaseConfig
}

// Provides a pointer to a databse connection for a given configuration.
func InitDB() *DataService {

	dbConfig := config.GetDatabaseFromGlobalConfig()

	switch dbConfig.Driver {

	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=%scharset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.SSL)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: glogger.Default.LogMode(glogger.Silent),
		})
		if err != nil {
			logger.Fatal("Database Connection Error: %s\n", err)
		}
		logger.Info("Connected to MySQL database.\n")
		if dbConfig.AutoMigrate {
			migrate(db)
		}
		return &DataService{Session: db, Config: dbConfig}

	case "mssql":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
			Logger: glogger.Default.LogMode(glogger.Silent),
		})
		if err != nil {
			logger.Fatal("Database Connection Error: %s\n", err)
		}
		logger.Info("Connected to Microsoft SQL Server database.\n")
		if dbConfig.AutoMigrate {
			migrate(db)
		}
		return &DataService{Session: db, Config: dbConfig}

	case "postgres":
		dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Database, dbConfig.Password, dbConfig.SSL)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: glogger.Default.LogMode(glogger.Silent),
		})
		if err != nil {
			logger.Fatal("Database Connection Error: %s\n", err)
		}
		logger.Info("Connected to PostgreSQL database.\n")
		if dbConfig.AutoMigrate {
			migrate(db)
		}
		return &DataService{Session: db, Config: dbConfig}
	/* case "sqlite":
	db, err := gorm.Open(sqlite.Open(dbConfig.Path), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
	})
	if err != nil {
		logger.Panic("Database Connection %s\n", err)
	}
	logger.Info("Connected to SQLite database.")
	if dbConfig.AutoMigrate {
		migrate(db)
	}
	return &DataService{Session: db, Config: dbConfig}
	*/
	case "":
		logger.Fatal("No database driver specified.\n")
		return nil

	default:
		logger.Fatal("Database Connection Error: invalid driver %s.\n", dbConfig.Driver)
		return nil
	}
}

// This method migrates all tables in the database
func migrate(db *gorm.DB) {
	logger.Info("Migrating database...\n")
	err := db.AutoMigrate(&models.CubeCDR{}, &models.CucmCdr{}, &models.CucmCmr{})
	if err != nil {
		logger.Fatal("Database Migration Error: %s\n", err)
	}
	logger.Info("Database migration complete.\n")
}
