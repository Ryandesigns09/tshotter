// Code generated by "stringer -type=ResourceInvalidationKind"; DO NOT EDIT.

package sema

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ResourceInvalidationKindUnknown-0]
	_ = x[ResourceInvalidationKindMoveDefinite-1]
	_ = x[ResourceInvalidationKindMoveTemporary-2]
	_ = x[ResourceInvalidationKindDestroy-3]
}

const _ResourceInvalidationKind_name = "ResourceInvalidationKindUnknownResourceInvalidationKindMoveDefiniteResourceInvalidationKindMoveTemporaryResourceInvalidationKindDestroy"

var _ResourceInvalidationKind_index = [...]uint8{0, 31, 67, 104, 135}

func (i ResourceInvalidationKind) String() string {
	if i >= ResourceInvalidationKind(len(_ResourceInvalidationKind_index)-1) {
		return "ResourceInvalidationKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ResourceInvalidationKind_name[_ResourceInvalidationKind_index[i]:_ResourceInvalidationKind_index[i+1]]
}