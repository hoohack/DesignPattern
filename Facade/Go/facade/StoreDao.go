package facade

type StoreDao struct {
}

func (sd *StoreDao) Info() string {
  return "store_name: uniqlo, store_address: guangzhou"
}
