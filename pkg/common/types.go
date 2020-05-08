package common

const (
	OutputTypeMySQL = "mysql"
	OutputTypeCSV   = "csv"
	OutputDbBook    = "book"
	OutputMongodb   =  "mongodb"
	OutputYaml  =  "yaml"
)

type MTS struct {
	ID     int
	Status TaskStatus
}
