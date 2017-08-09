package runtime

import (
	env "github.com/ta2gch/iris/runtime/environment"
	"github.com/ta2gch/iris/runtime/ilos"
	"github.com/ta2gch/iris/runtime/ilos/class"
	"github.com/ta2gch/iris/runtime/ilos/instance"
)

func tagbodyGo(args ilos.Instance, local *env.Environment, global *env.Environment) (ilos.Instance, ilos.Instance) {
	// args must be a instance of Cons, not Null, and ends with nil
	if !ilos.InstanceOf(args, class.Cons) || !UnsafeEndOfListIsNil(args) || UnsafeListLength(args) != 1 { // Checked at the head of test
		return nil, instance.NewWrongNumberOfArguments(instance.NewSymbol("GO"), args)
	}
	car := instance.UnsafeCar(args) // Checked at the top of this function
	if _, ok := local.TagBodyTag.Get(car); !ok {
		return nil, instance.NewSimpleError(instance.NewString("%v is not defined as the tag"), car)
	}
	return nil, instance.NewGo(car)
}
