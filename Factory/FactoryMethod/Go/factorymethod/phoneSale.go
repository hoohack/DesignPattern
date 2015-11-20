package factorymethod

type phoneSale interface {
  CreatePhone(ptype string) Phone
}
