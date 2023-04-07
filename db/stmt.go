package db

// 獲取 table 列表
const getTableList = "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'"

// 獲取備註
const getCommentStmt = `
SELECT COALESCE(d.description, '') 
FROM pg_attribute a 
LEFT JOIN pg_description d ON a.attrelid = d.objoid AND a.attnum = d.objsubid 
JOIN pg_class c ON c.oid = a.attrelid 
WHERE c.relname = '%s' AND a.attname = '%s';
`
