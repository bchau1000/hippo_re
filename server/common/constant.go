package common

var Error = struct {
	DecodeJson string
	ConvertSql string
	ExecuteSql string
}{
	DecodeJson: "Error occurred while decoding JSON",
	ConvertSql: "Error occurred while stringifying SQL",
	ExecuteSql: "Error occurred while executing SQL",
}
