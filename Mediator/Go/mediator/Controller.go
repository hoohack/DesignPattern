package mediator

type Controller struct {
  model *Model
  view *View
}

func (this *Controller) SetColleague(model *Model, view *View) {
  this.model = model
  this.view = view
}

func (this *Controller) AddUser(user string) {
  this.model.AddUser(user)
}

func (this *Controller) GetUserList() []string {
  return this.model.GetUserList()
}

func (this *Controller) DeleteUser(user string) {
  this.model.DeleteUser(user)
}

func (this *Controller) ShowResult(result string) {
  this.view.ShowResult(result)
}
