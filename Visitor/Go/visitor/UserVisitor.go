package visitor

type UserVisitor interface {
  AddPointForVip(vip_user *VipUser)
  AddPointForNormal(normal_user *NormalUser)
}
