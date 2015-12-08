package visitor

type UserApi interface {
  Accept(visitor UserVisitor)
}
