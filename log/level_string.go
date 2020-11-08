// Code generated by "stringer -type level -linecomment"; DO NOT EDIT.

package log

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[debug-0]
	_ = x[info-1]
	_ = x[warn-2]
	_ = x[err-3]
}

const _level_name = "DEBUGINFOWARNERROR"

var _level_index = [...]uint8{0, 5, 9, 13, 18}

func (i level) String() string {
	if i >= level(len(_level_index)-1) {
		return "level(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _level_name[_level_index[i]:_level_index[i+1]]
}
