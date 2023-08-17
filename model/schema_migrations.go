
package model

type SchemaMigrations struct { 
	// 
	Version int64 `json:"version"`  
	// 
	Dirty bool `json:"dirty"`  
}
