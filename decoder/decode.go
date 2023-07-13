package decoder

type MsgpackDecoder struct {
	data []byte
}

func NewMsgpackDecoder(data []byte) *MsgpackDecoder {
	return &MsgpackDecoder{
		data: data,
	}
}
