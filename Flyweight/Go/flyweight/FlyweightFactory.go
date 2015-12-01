package flyweight

type FlyweightFactory struct {
  fmMap map[string]*AuthorizationFactory
}

var instance *FlyweightFactory

func GetInstance() *FlyweightFactory {
  if instance == nil {
    instance = &FlyweightFactory{}
  }

  return instance
}

func (this *FlyweightFactory) GetFlyweight(key string) *AuthorizationFactory {
  if this.fmMap == nil {
    this.fmMap = make(map[string]*AuthorizationFactory)
  }

  if _, ok := this.fmMap[key]; !ok {
    this.fmMap[key] = NewAuthorizationFactory(key)
  }

  return this.fmMap[key]
}
