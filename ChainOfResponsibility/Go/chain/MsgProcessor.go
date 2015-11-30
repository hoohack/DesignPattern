package chain

type MsgProcessor struct {
  msg string
  filter_chain *FilterChain
}

func NewMsgProcessor(msg string, filter_chain *FilterChain) *MsgProcessor {
  return &MsgProcessor{msg, filter_chain}
}

func (this *MsgProcessor) Process() string {
  return this.filter_chain.HandleFilter(this.msg)
}
