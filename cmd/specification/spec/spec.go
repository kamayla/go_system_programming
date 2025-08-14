package spec

type Specification interface {
	IsSatisfied() bool
}

// ヘルパー関数
func And(specs ...Specification) Specification {
	if len(specs) == 0 {
		return nil
	}
	if len(specs) == 1 {
		return specs[0]
	}

	result := specs[0]
	for i := 1; i < len(specs); i++ {
		result = &AndSpecification{left: result, right: specs[i]}
	}
	return result
}

func Or(specs ...Specification) Specification {
	if len(specs) == 0 {
		return nil
	}
	if len(specs) == 1 {
		return specs[0]
	}

	result := specs[0]
	for i := 1; i < len(specs); i++ {
		result = &OrSpecification{left: result, right: specs[i]}
	}
	return result
}

func Not(spec Specification) Specification {
	return &NotSpecification{spec: spec}
}

// And仕様
type AndSpecification struct {
	left, right Specification
}

func (a *AndSpecification) IsSatisfied() bool {
	return a.left.IsSatisfied() && a.right.IsSatisfied()
}

// Or仕様
type OrSpecification struct {
	left, right Specification
}

func (o *OrSpecification) IsSatisfied() bool {
	return o.left.IsSatisfied() || o.right.IsSatisfied()
}

// Not仕様
type NotSpecification struct {
	spec Specification
}

func (n *NotSpecification) IsSatisfied() bool {
	return !n.spec.IsSatisfied()
}
