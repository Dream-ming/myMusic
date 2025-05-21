package user

// import (
//     "errors"

//     "github.com/Dream-ming/myMusic/initialize"
// )

// func GetUserInfo(userID int) (*model.User, error) {
// 	var u model.User
//     row := initialize.DB.QueryRow("SELECT id, username, age FROM users WHERE id = ?", userID)
//     err := row.Scan(&u.ID, &u.Username, &u.Age)
//     if err != nil {
//         return nil, err
//     }
//     return &u, nil
// }