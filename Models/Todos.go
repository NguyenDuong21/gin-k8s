package Models

import (
)

func GetAllTodos() ([]Todo) {
	var toDos = []Todo{
		{ID: "1", Title: "Blue Train", Description: "John Coltrane"},
		{ID: "2", Title: "Jeru", Description: "Gerry Mulligan"},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Description: "Sarah Vaughan"},
	}
	
	return toDos;
}

// func CreateATodo(todo *Todo) (err error) {
// 	if err = Config.DB.Create(todo).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetATodo(todo *Todo, id string) (err error) {
// 	if err := Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func UpdateATodo(todo *Todo, id string) (err error) {
// 	fmt.Println(todo)
// 	Config.DB.Save(todo)
// 	return nil
// }

// func DeleteATodo(todo *Todo, id string) (err error) {
// 	Config.DB.Where("id = ?", id).Delete(todo)
// 	return nil
// }
