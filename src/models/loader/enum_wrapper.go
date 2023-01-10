package loader

import (
	"fmt"
	"github.com/hexya-erp/hexya/src/models/types"
)

type EnumWrapper interface {
	fmt.Stringer
	Values() types.Selection
}
