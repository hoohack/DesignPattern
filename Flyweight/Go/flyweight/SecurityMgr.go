package flyweight

import "strings"
import "fmt"

type SecurityMgr struct {
  map_user map[string][]*AuthorizationFactory
}

var s_instance *SecurityMgr

var user_data = [2]string{"hhq,money,look", "tom,money,get"}

func GetSInstance() *SecurityMgr {
  if instance == nil {
    s_instance = &SecurityMgr{}
  }

  return s_instance
}

func (this *SecurityMgr) Login(user string) {
  col := this.queryUser(user)
  if this.map_user == nil {
    this.map_user = make(map[string][]*AuthorizationFactory)
  }
  this.map_user[user] = col
}

func (this *SecurityMgr) HasPermit(user string, security_entity string, permit string) bool {
  col := this.map_user[user]
  if col == nil || len(col) == 0 {
    fmt.Println(user + "没有登录或是没有被分配到任何权限")
    return false
  }

  for _, fm := range col {
    if fm.Match(security_entity, permit) {
      return true
    }
  }

  return false
}

func (this *SecurityMgr) queryUser(user string) []*AuthorizationFactory{
  var col []*AuthorizationFactory
  for _, data := range user_data {
    str := strings.Split(data, ",")
    if str[0] == user {
      fm := GetInstance().GetFlyweight(str[1] + "," + str[2])
      col = append(col, fm)
    }
  }

  return col
}
