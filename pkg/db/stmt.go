package db

// 獲取 postgres table 列表
const postgresTableStmt = "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'"

// 獲取 mysql table 列表
const mysqlTableStmt = "SHOW TABLES"

// 獲取 postgres 備註
const postgresCommentStmt = `
SELECT COALESCE(d.description, '') 
FROM pg_attribute a 
LEFT JOIN pg_description d ON a.attrelid = d.objoid AND a.attnum = d.objsubid 
JOIN pg_class c ON c.oid = a.attrelid 
WHERE c.relname = '%s' AND a.attname = '%s';
`

// 獲取 mysql 備註
const mysqlCommentStmt = `
SELECT COLUMN_NAME, COLUMN_COMMENT, DATA_TYPE 
FROM INFORMATION_SCHEMA.COLUMNS 
WHERE TABLE_NAME = '%s';
`
