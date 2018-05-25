package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"fmt"
	"bufio"
	"strings"
)

// Usage: ./enigma <mysql_user> <mysql_password> <database_name> <table_name> <proto_file>
func main() {

	args := os.Args[1:]

	if len(args) != 5 {
		log.Fatal("Missing param, usage: ./enigma <mysql_user> <mysql_password> <database_name> <table_name> <proto_file>")
		return
	}

	mysqlUser := args[0]
	mysqlPass := args[1]
	mysqlDataBase := args[2]
	mysqlTable := args[3]
	outputFile := args[4]
	infomationSchemaTable := "information_schema"

	// example: root:pass@/database?parseTime=true
	dataSourceName := fmt.Sprintf("%s:%s@/%s", mysqlUser, mysqlPass, infomationSchemaTable)
	mysqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Can not connect to mysql, detail: ", err)
	}

	defer mysqlDB.Close()

	stmt, err := mysqlDB.Prepare("SELECT COLUMN_NAME, DATA_TYPE FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME =?")
	if err != nil {
		log.Println("Error when prepare query, detail: ", err)
		return
	}

	rows, err := stmt.Query(mysqlDataBase, mysqlTable)
	if err != nil {
		log.Println("Error when exec query, detail: ", err)
		return
	}

	// open output file
	fo, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			log.Println("Error when exec query, detail: ", err)
			return
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	// write message name
	if _, err := w.Write([]byte("syntax = \"proto3\";\n\n")); err != nil {
		log.Println("Error when write data to file, detail: ", err)
		return
	}

	if _, err := w.Write([]byte(fmt.Sprintf("message %s {\n", GetCamelCase(mysqlTable)))); err != nil {
		log.Println("Error when write data to file, detail: ", err)
		return
	}

	counter := 0
	for rows.Next() {
		var columnName, dataType string
		err = rows.Scan(&columnName, &dataType)
		if err != nil {
			log.Println("Error when scan rows, detail: ", err)
			return
		}
		//log.Println("ColumnName: ", columnName, ", DataType: ", dataType)
		counter++
		buf := []byte(fmt.Sprintf("\t%s %s = %d;\n", GetProtoDataType(dataType), columnName, counter))

		// write a chunk
		if _, err := w.Write(buf); err != nil {
			log.Println("Error when write data to file, detail: ", err)
			return
		}
	}

	if _, err := w.Write([]byte("}\n")); err != nil {
		log.Println("Error when write data to file, detail: ", err)
		return
	}

	if err = w.Flush(); err != nil {
		log.Println("Error when flush data to file, detail: ", err)
		return
	}

}

func GetProtoDataType(mysqlType string) string {
	switch mysqlType {
	case "varchar":
		return "string"
	case "bigint":
		return "int64"
	case "tinyint":
		return "bool"
	case "datetime":
		return "int64"
	default:
		return ""
	}
}

func GetCamelCase(input string) string {
	output := ""
	capNext := true
	for _, v := range input {
		if v >= 'A' && v <= 'Z' {
			output += string(v)
		}
		if v >= '0' && v <= '9' {
			output += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				output += strings.ToUpper(string(v))
			} else {
				output += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}

	return output
}
