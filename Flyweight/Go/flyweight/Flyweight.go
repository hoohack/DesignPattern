package flyweight

type Flyweight interface {
  Match(security_entity string, permit string) bool
}
