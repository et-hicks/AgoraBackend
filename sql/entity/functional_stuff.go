package entity

type FunctionalStuff struct {
	BEFSJson string
	FEFSJson string
	BEFS map[string]interface{} `gorm:"-:all"`
	FEFS map[string]interface{} `gorm:"-:all"`
}
