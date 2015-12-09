package visitor

type ObjectStructure struct {
  user_list []UserApi
}

func (this *ObjectStructure) AddUser(user UserApi) {
  this.user_list = append(this.user_list, user)
}

func (this *ObjectStructure) HandleVisitor(visitor UserVisitor) {
  for _, user := range this.user_list {
    user.Accept(visitor)
  }
}
