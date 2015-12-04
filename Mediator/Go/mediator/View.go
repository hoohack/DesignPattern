package mediator

import "fmt"

type View struct {
  *Colleague
}

func NewView(media Mediator) *View {
  return &View{&Colleague{media}}
}

func (this *View) MakeAddRequest(user string) {
  fmt.Println("发起增加用户" + user + "请求")
  this.Colleague.GetMediator().AddUser(user)
}

func (this *View) ShowUser() {
  fmt.Println("当前用户如下：\n--------")
  user_list := this.Colleague.GetMediator().GetUserList()
  for _, u := range user_list {
    fmt.Println(u)
  }
  fmt.Println("--------")
}

func (this *View) MakeDeleteRequest(user string) {
  fmt.Println("发起删除用户" + user + "请求")
  this.Colleague.GetMediator().DeleteUser(user)
}

func (this *View) ShowResult(result string) {
  fmt.Println(result)
}
