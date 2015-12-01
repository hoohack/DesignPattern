package flyweight

import "strings"

type AuthorizationFactory struct {
  security_entity string
  permit string
}

func NewAuthorizationFactory(state string) *AuthorizationFactory {
  arr := strings.Split(state, ",")
  return &AuthorizationFactory{arr[0], arr[1]}
}

func (this *AuthorizationFactory) GetSecurityEntity() string {
  return this.security_entity
}

func (this *AuthorizationFactory) GetPermit() string {
  return this.permit
}

func (this *AuthorizationFactory) Match(security_entity string, permit string) bool {
  if this.security_entity == security_entity && this.permit == permit {
    return true
  }

  return false
}
