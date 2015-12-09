package visitor

type User struct {
  name string
  point int
}

func NewUser(name string) *User {
  return &User{name, 0}
}

func (this *User) AddPoint(point int) {
  this.point += point
}

func (this *User) GetName() string {
  return this.name
}

func (this *User) GetPoint() int {
  return this.point
}
