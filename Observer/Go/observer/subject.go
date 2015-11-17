package observer

type Subject interface
{
  RegisterObserver(ob Observer)
  Notify()
}
