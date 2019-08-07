package dbUtil

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"HFish/error"
)

// 连接数据库
func conn() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/hfish.db")
	error.Check(err, "连接数据库失败")
	return db
}

// 插入数据
func Insert(sql string, args ...interface{}) int64 {
	/*
	参数说明：

		sql  insert 语句
		args insert value 参数

	使用案例：

		sql := `
		INSERT INTO coot_tasks (
			task_name,
			task_explain,
			task_id,
			task_time_type,
			task_time,
			last_exec_time,
			is_plug_script,
			script_type,
			script_path,
			alert_type,
			create_time
		)
		VALUES
			(?,?,?,?,?,?,?,?,?,?,?);
		`
		dbUtil.Insert(sql, "插入任务测试", "测试说明", "", 1, "2", "", "1", "shell", "/scripts/myscript/test.sh", "1", "2019-07-10 16:12")
	*/
	db := conn()
	stmt, _ := db.Prepare(sql)

	res, err := stmt.Exec(args...)
	error.Check(err, "插入数据失败")

	defer stmt.Close()

	id, err := res.LastInsertId()
	error.Check(err, "获取插入ID失败")

	defer db.Close()

	// 返回 自增长 ID
	return id
}

// 更新数据
func Update(sql string, args ...interface{}) int64 {
	/*
	参数说明：

		sql  update 语句
		args update 参数

	使用案例：

		sql := `
		UPDATE coot_tasks
		SET task_name = ?
		WHERE
			id = ?;
		`
		dbUtil.Update(sql, "任务更新测试", 1)
	*/
	db := conn()
	stmt, _ := db.Prepare(sql)

	res, err := stmt.Exec(args...)

	error.Check(err, "更新数据失败")
	defer stmt.Close()

	affect, err := res.RowsAffected()
	error.Check(err, "获取影响行数失败")

	defer db.Close()

	return affect
}

// 查询数据
func Query(sql string, args ...interface{}) []map[string]interface{} {
	/*
	参数说明：

		sql  select 语句
		args select 参数

	使用案例：

		sql := `select * from coot_tasks where id=?;`
		result := dbUtil.Query(sql, 1)
	*/

	db := conn()

	rows, err := db.Query(sql, args ...)
	error.Check(err, "查询数据失败")

	defer rows.Close()

	columns, err := rows.Columns()
	error.Check(err, "查询表名失败")

	count := len(columns)

	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	defer db.Close()

	return tableData
}

// 删除数据
func Delete(sql string, args ...interface{}) int64 {
	/*
	参数说明：

		sql  delete 语句
		args delete 参数

	使用案例：

		sql := `delete from coot_tasks where id=?;`
		dbUtil.Delete(sql, 2)
	*/
	db := conn()

	stmt, _ := db.Prepare(sql)

	res, err := stmt.Exec(args...)
	error.Check(err, "删除数据失败")
	defer stmt.Close()

	affect, err := res.RowsAffected()
	error.Check(err, "获取影响行数失败")

	defer db.Close()

	return affect
}
