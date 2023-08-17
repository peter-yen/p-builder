
package model

type PetHouse struct { 
	// 
	Id int32 `json:"id"`  
	// 
	Title string `json:"title"`  
	// 
	AnimalType int16 `json:"animal_type"`  
	// 
	AnimalBreed int32 `json:"animal_breed"`  
	// 
	Status bool `json:"status"`  
	// 
	Address string `json:"address"`  
	// 
	City string `json:"city"`  
	// 
	Contacts interface {} `json:"contacts"`  
	// 
	TwlicenseValiddate time.Time `json:"twlicense_validdate"`  
	// 
	TwlicenseNum string `json:"twlicense_num"`  
	// 
	OwnerName string `json:"owner_name"`  
}
