//Package imports imports a package from the same module
package imports

import "github.com/erjillsison/testgo/nested"

func Imports() {
	nested.NestedFunc()
}
