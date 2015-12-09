package main

import "./visitor"

func main() {
  add_point_visitor := &visitor.AddPointVisitor{}
  os := &visitor.ObjectStructure{}
  os.AddUser(visitor.NewVipUser("vip_hector"))
  os.AddUser(visitor.NewNormalUser("continue"))
  user_tom := visitor.NewNormalUser("tom")
  os.AddUser(user_tom)
  os.AddUser(user_tom)
  os.HandleVisitor(add_point_visitor)
}
