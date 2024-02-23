package database

func TemplateInitSchema() string {
	res := "create table logs\n(\n    date    DateTime,\n    base_name Nullable(String),\n    context String,\n    comment String,\n    handler Nullable(String)\n)\n    engine = Memory;\n\n"
	return res
}

func TemplateCheckSchema() string {
	res := "SELECT * FROM logs"
	return res
}

func TemplateAddLog() string {
	res := "INSERT INTO logs VALUES (now(), '%v', '%v', '%v', '%v', '%v')"
	return res
}

func TemplateNow() string {
	return "SELECT NOW()"
}
