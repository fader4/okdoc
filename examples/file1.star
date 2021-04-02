load("foo", "bar")

load("a", b = "c", "qw")

# inline comment

foo = module(
	"bar"
)

foo = module(
	"bar",
	a = "b",
	aff = not a,
	uniq = [1,2,3].uniq(),
	a = "b".format(),
	assert = 1 + (2 + 3),
	ar = f(**dict(c=7, obj=dict(c=7, a=2, b=3), a=2, b=3, fn=f(**dict(c=7, a=2, b=3)))),
	dic = dict(c=7, a=2, b=3),
	dic = {"a": "b", "c": 123, "d": {"a": "b", "c": 123}, "fn": f(**dict(c=7, obj=dict(c=7, a=2, b=3), a=2, b=3, fn=f(**dict(c=7, a=2, b=3))))}
)

"""
Free comment
"""

"""
@Param(ctx, "ctx with fields")
@Param(ctx.foo, description = "multi
line
description
with\"escaped\" strings
")
@Param(name = ctx.bar, description = "ctx with fields")
@Param(name = ctx.bar, description = "ctx with fields", example = 123)
@Param(name = ctx.bar, description = "ctx with fields", example = 123.123)
@Param(name = ctx.bar, description = "ctx with fields", example = True)
@Param(name = ctx.bar, description = "ctx with fields", example = ident.a.b.c)
@Param(name = ctx.bar, description = "ctx with fields", example = ["a", 123, ["a", b], {a = "b"}])
@Param(name = ctx.bar, description = "ctx with fields", example = {a = "b"})
@Param(name = ctx.bar, description = "ctx with fields",
    example  = "repeat_same_field",
    example  = "repeat_same_field",
    example  = "repeat_same_field",
    example  = "repeat_same_field"
)
"""
def main(ctx=(1,2,3) ,  obj={"a": "b"}, **kwargs):
	"""
	Inline comment into method
	"""

	if False:
		# @Return("Some case")
		return False
	# @Return("Happy path")
	return True
