# Enigma

A small tool to generate protobuf file from database

Currently support database: MySQL

Requirement:
- A open connection to database server server, default port is 3306
- MySQL user with `read` permission on database `information_schema`

Run:
```bash
go build
./enigma gen --type=<...> --host=<...> --port=<...> --user=<...> --pass=<...> --db=<...> --table=<..> --file=<..>
```

Example:
```bash
./enigma gen --type=mysql --host=localhost --port=3306 --user=root --pass=secret --db=staging --table=apps --file=app.proto
```