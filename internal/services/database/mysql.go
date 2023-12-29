package database

func TemplateInitSchema() string {
	res := "create table logs\n(\n    id              int auto_increment\n        primary key,\n    datetime        datetime default (now()) not null,\n    baseID          int                      null,\n    baseName        char(255)                null,\n    context         char(255)                null,\n    internalContext char(255)                null,\n    comment         char(255)                not null,\n    level           char(10)                 null,\n    constraint logs_id_uindex\n        unique (id)\n)\n    comment 'Логирование ошибок и информационных записей работы программы прокси 1С' collate = utf8mb4_0900_ai_ci;\n\ncreate index logs_baseID_index\n    on logs (baseID);\n\n"
	return res
}

func TemplateCheckSchema() string {
	res := "SELECT * FROM logs"
	return res
}

func TemplateAddLog() string {
	res := "INSERT INTO logs VALUES (id, now(), '%v', '%v', '%v', '%v', '%v', '%v')"
	return res
}

func TemplateNow() string {
	return "SELECT NOW()"
}
