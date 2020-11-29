// @Desc:
// @CreateTime: 2020/11/29
package server

import (
    "log"
    "my_go/Week02/dao"
)

type Response struct {
     Code int `json:"code,omitempty"` // 0 - fail, 1 - success
     Msg string `json:"msg,omitempty"`
     Data interface{} `json:"data,omitempty"`
}

// get user id
func GetUserId()  {
    var (
        userId int
        err error
    )
    userId, err = dao.SelectUserId()
    resp := &Response{Code: 1}
    if err != nil {
        resp.Code = 0
        resp.Msg = "您查询 user_id 不存在"
        log.Fatalf("select userId is failed, err: %#v", err)
    } else {
        resp.Data = map[string]int{"user_id": userId}
    }

    // TODO json.Marshal() 后并响应
}