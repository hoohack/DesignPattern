package mediator

type Model struct {
  *Colleague
  user_list []string
}

func NewModel(mediator Mediator) *Model {
  return &Model{Colleague: &Colleague{mediator}}
}

func (this *Model) AddUser(user string) {
  this.user_list = append(this.user_list, user)
}

func (this *Model) GetUserList() []string {
  return this.user_list
}

func (this *Model) DeleteUser(user string) {
  for idx, u := range this.user_list {
    if u == user {
      this.user_list = append(this.user_list[:idx], this.user_list[idx+1:]...)
      this.Colleague.GetMediator().ShowResult("删除" + user + "成功")
      return
    }
  }

  this.Colleague.GetMediator().ShowResult("删除" + user + "失败，该用户不存在")
}
