package database

import (
	"fmt"
	"os"
	"strings"
)

func Backup(sql string, backupFile string) {
	log.WithField("sql", sql).WithField("backupFilename", backupFile).Info("begin to backup sql")

	if !strings.Contains(sql, "limit") {
		if strings.Contains(sql, ";") {
			strings.Replace(sql, ";", " ", -1)
		}
		sql = fmt.Sprintf("%s %s", sql, "limit 10;")
	}

	sqlDB, err := GetDB().DB()
	defer func() {
		err = CloseDB()
		if err != nil {
			log.Fatalf("close database failed %s", dsn)
		}
	}()

	if err != nil {
		log.Fatalf("get sql db failed")
	}

	rows, err := sqlDB.Query(sql)
	if err != nil {
		log.Fatalf("query database failed", err)
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	var content string
	for rows.Next() {
		//columns, _ := rows.Columns()
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Fatalf("get data from rows failed ", err)
		}

		rowData := make([]string, len(columns))
		for i, val := range values {
			if val == nil {
				rowData[i] = "NULL"
			} else {
				rowData[i] = fmt.Sprintf("%s", val)
			}
		}

		result := strings.Join(rowData, ",")
		result += "\n"
		log.Info("the row result: %s", result)
		content += result

	}

	err = os.WriteFile(backupFile, []byte(content), 0644)
	if err != nil {
		log.Fatalf("backup file failed ", backupFile)
	}
}
