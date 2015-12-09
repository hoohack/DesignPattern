package visitor

type VipUser struct {
  *User
}

func NewVipUser(name string) *VipUser {
  return &VipUser{&User{name, 0}}
}

func (this *VipUser) Accept(visitor UserVisitor) {
  visitor.AddPointForVip(this)
}
