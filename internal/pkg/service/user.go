package service

//
//import (
//	"feynman-podcast/internal/pkg/model/user"
//	"feynman-podcast/internal/pkg/util"
//)
//
//func (c *Client) GetUser(email string) *user.Data {
//	return c.UserDataCollection.GetData(email)
//}
//
//func (c *Client) GetUserWithPassword(email, password string) *user.Data {
//	hash, _ := util.HashPassword(password)
//	return c.UserDataCollection.GetDataWithHash(email, hash)
//}
//
//func (c *Client) CreateUser(data *user.Data) error {
//	if h, err := util.HashPassword(data.PassWord); err != nil {
//		return err
//	} else {
//		d := &user.Data{
//			Email:    data.Email,
//			PassWord: h,
//			Name:     data.Name,
//		}
//		return c.UserDataCollection.CreateData(d)
//	}
//}
