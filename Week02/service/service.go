package service

import "Week02/dao"

func GetID(id int) (int, error) {
	return dao.QueryID(id)
}
func GetName(name string) (string, error) {
	return dao.QueryName(name)
}
