package cmd

import (
	"gomysql/models"
	"strings"
)

type Value models.Value

func (v *Value) Delete(tbl string) *Value {
	v.Data += "DELETE FROM" + tbl + " "
	return v

}
func (v *Value) Select(cond string) *Value {
	v.Data += "SELECT " + cond + " "
	return v
}
func (v *Value) From(tbl string) *Value {
	v.Data += "FROM " + tbl + " "
	return v
}
func (v *Value) Where(cond string) *Value {
	v.Data += "WHERE " + cond + " "
	return v
}
func (v *Value) InsertInto(tbl string) *Value {
	v.Data += "INSERT INTO " + tbl + " "
	return v
}
func (v *Value) Fields(fields ...string) *Value {
	var newField = "("
	for _, fields := range fields {
		newField += fields + ", "
	}
	v.Data += newField[:len(newField)-2] + ") "
	return v
}
func (v *Value) Values(val ...string) *Value {
	var newVal = ""
	for _, val := range val {
		newVal += val + ","
	}
	v.Data += "VALUES " + newVal[:len(newVal)-1]
	return v
}
func (v *Value) Update(tbl string) *Value {
	v.Data += "UPDATE" + tbl + " "
	return v
}
func (v *Value) Set(conditions ...string) *Value {
	var newVal = ""
	for _, val := range conditions {
		newVal += val + ","
	}
	v.Data += "VALUES " + newVal[:len(newVal)-1]
	return v
}
func (v *Value) GenerateQuery() string {
	return strings.TrimSpace(v.Data) + ";"
}
