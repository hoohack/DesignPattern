package facade

type ServiceFacade struct {
  *StoreDao
  *GoodsDao
}

func (sf *ServiceFacade) GoodsDetail() [2]string {
  return [2]string{sf.StoreDao.Info(), sf.GoodsDao.Info()}
}
