# Enigma

A small tool to generate protobuf file from MySQL

Requirement:
- A open connection to MySQL server, default port is 3306
- MySQL user with `read` permission on table `information_schema`

Run:
```bash
cd cmd/enigma
go build
./enigma <mysql_user> <mysql_password> <database_name> <table_name> <proto_file>
```

Example:
```bash
./enigma my_user secret_pass my_db_name my_tbl my_tbl.proto
```