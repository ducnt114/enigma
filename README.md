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

Help:
```bash
./enigma help
Enigma is small tool to generate protobuf file from MySQL,
You don't need to create proto file manual

Usage:
  enigma [flags]
  enigma [command]

Available Commands:
  gen         Generate proto file
  help        Help about any command
  version     Print the version number of Enigma

Flags:
  -h, --help   help for enigma

Use "enigma [command] --help" for more information about a command.
```

or

```bash
./enigma gen --help
Generate proto file from MySQL

Usage:
  enigma gen [flags]

Flags:
  -d, --db string      db name (default "db_name")
  -f, --file string    proto file name (default "proto")
  -h, --help           help for gen
  -o, --host string    db host (default "localhost")
  -w, --pass string    db pass (default "secret")
  -p, --port int32     db port (default 3306)
  -t, --table string   db name (default "table_name")
  -g, --type string    db type: mysql, postgres (default "mysql")
  -u, --user string    db user (default "root")
```