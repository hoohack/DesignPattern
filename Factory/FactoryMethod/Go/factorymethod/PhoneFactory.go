package factorymethod

type PhoneFactory interface {
  createPhone() Phone
}
