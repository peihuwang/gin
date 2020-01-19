/**
 * Created by Goland.
 * User: TinyStar
 * Date: 2019/11/11
 * Time: 10:56
 */
package entities

type UserId struct {
	Id  int64 `json:"id" form:"id"`
}

type UserInfo struct {
	Id  int64 `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	IsDelete  string `json:"is_delete" form:"is_delete"`
}