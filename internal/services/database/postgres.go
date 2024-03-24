package database

func TemplateInitSchema() string {
	res := "create table logs\n(\n    date    TIMESTAMP,\n    base_name varchar(25),\n    context varchar(400),\n    comment varchar(400),\n    handler varchar(400), level varchar(5)\n)\n   ;\n\n"
	return res
}

func TemplateCheckSchema() string {
	res := "SELECT * FROM logs LIMIT 1"
	return res
}

func TemplateAddLog() string {
	res := "INSERT INTO logs VALUES (now(), '%v', '%v', '%v', '%v', '%v')"
	return res
}

func TemplateNow() string {
	return "SELECT NOW()"
}

func GetLogs() string {
	res := `select date, base_name, context, comment, handler, level from logs order by date desc limit 200`
	return res
}
