# Go-CDR

Parses Cisco Collaboration Manager (CUCM/CCM) CDR/CMR files as well as Cisco UBE (CUBE) CDR files and inserts them into a database.
Inserts CDR/CMR records in bulk to improve performance, and utilizes UTC time for insertion. If the files are not in UTC time, the time will be converted to UTC time.

## Usage

``` bash
go-cdr --directory "C:\CDR" --cdr_type (cucm|cube) --config "config.yaml"
```

## Limitations

* Only supports CUCM/CCM and CUBE CDR/CMR files
* Only supports PostgreSQL, MySQL, Microsoft SQL Server, and SQLite databases
* Only supports CDR/CMR files in CSV format
* Will parse directories in single-threaded mode, i.e. one directory at a time

## Example Config

``` yaml
database:
  autoMigrate: true # Migrate the database schema on startup
  database: cdr # Database name
  driver: postgres # Database driver (mysql|mssql|postgres|sqlite)
  host: localhost # Database host
  limit: 100 # Maximum number of records to insert in bulk
  password: 012345abc # Database password
  path: ./go_cdr.db # Path to the database file (sqlite only)
  port: 5432 # Database port
  username: postgres # Database username
  SSL: disable # Database SSL mode (disable|require|verify-ca|verify-full)
logging:
  compress: false # Compress log files
  level: info # Logging level (debug|info|warn|error|fatal|panic)
  maxAge: 30 # Maximum age of log files in days
  maxSize: 100 # Maximum size of log files in megabytes
  name: go-cdr.log # Name of the log files
  path: ./logs # Path to store log files
parser:
  parseInterval: 30 # Interval in minutes to parse files
  directories:
  - input: D:\CDR\cube_cdr\home\cubecdr\ftp # Path to the CDR files
    output: D:\CDR\cube_cdr\home\cubecdr\ftp\processed # Path to move the CDR files after parsing
    type: cube # Type of CDR files (cucm|cube)
    deleteOriginal: false # Delete original files after parsing
  - input: D:\CDR\cucm_cdr\home\cucmcdr\ftp # Path to the CDR files
    output: D:\CDR\cucm_cdr\home\cucmcdr\ftp\processed # Path to move the CDR files after parsing
    type: cucm # Type of CDR files (cucm|cube)
    deleteOriginal: false # Delete original files after parsing
```
