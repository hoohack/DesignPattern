package observer

type Observer interface {
  Update(s *Subject) string
}
