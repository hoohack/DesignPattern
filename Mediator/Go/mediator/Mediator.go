package mediator

type Mediator interface {
  AddUser(user string)
  GetUserList() []string
  DeleteUser(user string)
  ShowResult(result string)
}
