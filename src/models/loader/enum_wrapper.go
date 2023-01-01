package loader

import "github.com/hexya-erp/hexya/src/models/types"

type EnumWrapper interface {
	Values() types.Selection
}
