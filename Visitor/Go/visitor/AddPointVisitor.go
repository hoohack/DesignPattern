package visitor

import "fmt"

type AddPointVisitor struct {
}

func (this *AddPointVisitor) AddPointForVip(vip_user *VipUser) {
  vip_user.AddPoint(10)
  fmt.Printf("给VIP%s加10分，总分%d分\n", vip_user.GetName(), vip_user.GetPoint())
}

func (this *AddPointVisitor) AddPointForNormal(normal_user *NormalUser) {
  normal_user.AddPoint(5)
  fmt.Printf("给普通用户%s加5分，总分%d分\n", normal_user.GetName(), normal_user.GetPoint())
}
