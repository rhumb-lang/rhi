package vm

type Code uint8

const OPCODE_INDEX_MASK uint8 = 0b11100000
const TAG_IDX_EXTENSION uint8 = 0b00000000
const TAG_VALUE_LITERAL uint8 = 0b00100000
const TAG_LOCAL_REQUEST uint8 = 0b01000000
const TAG_INNER_REQUEST uint8 = 0b01100000
const TAG_UNDER_REQUEST uint8 = 0b10000000
const TAG_OUTER_REQUEST uint8 = 0b10100000
const TAG_EVENT_REQUEST uint8 = 0b11000000
const TAG_REPLY_REQUEST uint8 = 0b11100000
const MAX_IDX_SIZE uint32 = 31

func (c Code) String() (result string) {
	switch c.Tag() {
	case TAG_IDX_EXTENSION:
		result = "TAG_IDX_EXTENSION"
	case TAG_VALUE_LITERAL:
		result = "TAG_VALUE_LITERAL"
	case TAG_LOCAL_REQUEST:
		result = "TAG_LOCAL_REQUEST"
	case TAG_INNER_REQUEST:
		result = "TAG_INNER_REQUEST"
	case TAG_UNDER_REQUEST:
		result = "TAG_UNDER_REQUEST"
	case TAG_OUTER_REQUEST:
		result = "TAG_OUTER_REQUEST"
	case TAG_EVENT_REQUEST:
		result = "TAG_EVENT_REQUEST"
	case TAG_REPLY_REQUEST:
		result = "TAG_REPLY_REQUEST"
	}
	return
}

func generateCodes(tag uint8, idx uint32) []Code {
	codeWithExtensions := make([]Code, (idx/MAX_IDX_SIZE)+1)
	remaining := idx
	for i := range codeWithExtensions {
		if remaining > MAX_IDX_SIZE {
			codeWithExtensions[i] =
				Code(TAG_IDX_EXTENSION | uint8(MAX_IDX_SIZE))
		} else {
			codeWithExtensions[i] =
				Code(tag | uint8(remaining))
		}
		remaining = remaining - MAX_IDX_SIZE
	}
	return codeWithExtensions
}

func (c Code) Index() uint8 {
	return uint8(c) & uint8(MAX_IDX_SIZE)
}

func (c Code) Tag() uint8 {
	return uint8(c) & OPCODE_INDEX_MASK
}

func (c Code) Pair() (uint8, uint8) {
	return c.Tag(), c.Index()
}

func (c Code) IsIndexExtension() bool {
	return c.Tag() == TAG_IDX_EXTENSION
}

func NewValueLiteral(idx uint32) []Code {
	return generateCodes(TAG_VALUE_LITERAL, idx)
}

func (c Code) IsValueLiteral() bool {
	return c.Tag() == TAG_VALUE_LITERAL
}

func NewLocalRequest(idx uint32) []Code {
	return generateCodes(TAG_LOCAL_REQUEST, idx)
}

func (c Code) IsLocalRequest() bool {
	return c.Tag() == TAG_LOCAL_REQUEST
}

func NewInnerRequest(idx uint32) []Code {
	return generateCodes(TAG_INNER_REQUEST, idx)
}

func (c Code) IsInnerRequest() bool {
	return c.Tag() == TAG_INNER_REQUEST
}

func NewUnderRequest(idx uint32) []Code {
	return generateCodes(TAG_UNDER_REQUEST, idx)
}

func (c Code) IsUnderRequest() bool {
	return c.Tag() == TAG_UNDER_REQUEST
}

func NewOuterRequest(idx uint32) []Code {
	return generateCodes(TAG_OUTER_REQUEST, idx)
}

func (c Code) IsOuterRequest() bool {
	return c.Tag() == TAG_OUTER_REQUEST
}

func NewEventRequest(idx uint32) []Code {
	return generateCodes(TAG_EVENT_REQUEST, idx)
}

func (c Code) IsEventRequest() bool {
	return c.Tag() == TAG_EVENT_REQUEST
}

func NewReplyRequest(idx uint32) []Code {
	return generateCodes(TAG_REPLY_REQUEST, idx)
}

func (c Code) IsReplyRequest() bool {
	return c.Tag() == TAG_REPLY_REQUEST
}
