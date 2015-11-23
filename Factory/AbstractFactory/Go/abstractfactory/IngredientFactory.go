package abstractfactory

type IngredientFactory interface {
  createOS(name string)
  createSpecial(content string)
}
