// Code generated by "stringer -type=CommandType -output=strings.go"; DO NOT EDIT

package ami

import "fmt"

const _CommandType_name = "ActionResponseEvent"

var _CommandType_index = [...]uint8{0, 6, 14, 19}

func (i CommandType) String() string {
	i -= 1
	if i < 0 || i >= CommandType(len(_CommandType_index)-1) {
		return fmt.Sprintf("CommandType(%d)", i+1)
	}
	return _CommandType_name[_CommandType_index[i]:_CommandType_index[i+1]]
}
