package day24

type evalNode interface {
	Evaluate() bool
}

type andNode struct {
	left, right evalNode
}

func (A *andNode) Evaluate() bool {
	return A.left.Evaluate() && A.right.Evaluate()
}

type orNode struct {
	left, right evalNode
}

func (O *orNode) Evaluate() bool {
	return O.left.Evaluate() || O.right.Evaluate()
}

type xorNode struct {
	left, right evalNode
}

func (X *xorNode) Evaluate() bool {
	return X.left.Evaluate() != X.right.Evaluate()
}

type valueNode struct {
	value bool
}

func (V *valueNode) Evaluate() bool {
	return V.value
}

type op uint8

const (
	OPERATION_ADD op = op(1)
	OPERATION_OR  op = op(2)
	OPERATION_XOR op = op(3)
)

type nilableNode struct {
	left, right evalNode
	opType      op
}

func (X *nilableNode) Evaluate() bool {
	if X.left == nil || X.right == nil {
		panic("nil node cannot be evaluated")
	}
	switch X.opType {
	case OPERATION_ADD:
		return X.left.Evaluate() && X.right.Evaluate()
	case OPERATION_OR:
		return X.left.Evaluate() || X.right.Evaluate()
	case OPERATION_XOR:
		return X.left.Evaluate() != X.right.Evaluate()
	default:
		panic("unknown operation")

	}
}
