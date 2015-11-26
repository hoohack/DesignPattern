package proxy

type DBQueryProxy struct {
  real *DBQuery
}

func (this *DBQueryProxy) Request() {
  if this.real == nil {
    this.real = NewDBQuery()
  }

  this.real.Request()
}
