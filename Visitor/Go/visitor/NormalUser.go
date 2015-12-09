package visitor

type NormalUser struct {
  *User
}

func NewNormalUser(name string) *NormalUser {
  return &NormalUser{&User{name, 0}}
}

func (this *NormalUser) Accept(visitor UserVisitor) {
  visitor.AddPointForNormal(this)
}
