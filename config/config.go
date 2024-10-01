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

package config

import (
	"log"

	"github.com/spf13/viper"
)

type GlobalConfig struct {
	Database *DatabaseConfig
	Logging  *LoggingConfig
	Parser   *ParserConfig
}

type DatabaseConfig struct {
	AutoMigrate bool
	Database    string
	Driver      string
	Host        string
	Limit       uint32
	Password    string
	Path        string
	Port        int
	Username    string
	SSL         string
}

type LoggingConfig struct {
	Compress bool
	Level    string
	MaxAge   uint32
	MaxSize  uint32
	Name     string
	Path     string
}

type ParserConfig struct {
	Directories   []DirectoryConfig `mapstructure:"directories"`
	ParseInterval int
}

type DirectoryConfig struct {
	Input          string `mapstructure:"input"`
	Output         string `mapstructure:"output"`
	Type           string `mapstructure:"type"`
	DeleteOriginal bool   `mapstructure:"deleteOriginal"`
}

func SetDefaults() {
	// Set defaults for the LoggingConfig
	viper.SetDefault("logging.compress", true)
	viper.SetDefault("logging.level", "debug")
	viper.SetDefault("logging.maxAge", 7)
	viper.SetDefault("logging.maxSize", 10)
	viper.SetDefault("logging.name", "go-cdr.log")
	viper.SetDefault("logging.path", "./go-cdr/log")

	// Set defaults for the DatabaseConfig
	viper.SetDefault("database.autoMigrate", true)
	// viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.path", "./go-cdr/db/go-cdr.db")
	viper.SetDefault("database.limit", 100)

}

func GetLoggerFromGlobalConfig() *LoggingConfig {
	loggerConfig := viper.Sub("logging")
	if loggerConfig == nil {
		log.Fatalf("No log settings found in config file")
		return nil
	}
	return &LoggingConfig{
		Compress: loggerConfig.GetBool("compress"),
		Level:    loggerConfig.GetString("level"),
		MaxAge:   loggerConfig.GetUint32("maxAge"),
		MaxSize:  loggerConfig.GetUint32("maxSize"),
		Name:     loggerConfig.GetString("name"),
		Path:     loggerConfig.GetString("path"),
	}
}

func GetParserFromGlobalConfig() *ParserConfig {
	parserConfig := viper.Sub("parser")
	if parserConfig == nil {
		log.Fatalf("No parser settings found in config file")
		return nil
	}
	return &ParserConfig{
		ParseInterval: parserConfig.GetInt("parseInterval"),
	}
}

func GetDirectoriesFromGlobalConfig() []DirectoryConfig {

	var directories []DirectoryConfig

	viper.UnmarshalKey("parser.directories", &directories)

	return directories
}

func GetDatabaseFromGlobalConfig() *DatabaseConfig {
	databaseConfig := viper.Sub("database")
	if databaseConfig == nil {
		log.Fatalf("No database settings found in config file")
		return nil
	}
	return &DatabaseConfig{
		AutoMigrate: databaseConfig.GetBool("autoMigrate"),
		Database:    databaseConfig.GetString("database"),
		Driver:      databaseConfig.GetString("driver"),
		Host:        databaseConfig.GetString("host"),
		Limit:       databaseConfig.GetUint32("limit"),
		Password:    databaseConfig.GetString("password"),
		Path:        databaseConfig.GetString("path"),
		Port:        databaseConfig.GetInt("port"),
		Username:    databaseConfig.GetString("username"),
	}
}

func GetGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		Logging:  GetLoggerFromGlobalConfig(),
		Database: GetDatabaseFromGlobalConfig(),
	}
}
