package FSM

import (
	"strings"

	ut "github.com/PlayerR9/MyGoLib/CustomData/Tray"
	ub "github.com/PlayerR9/MyGoLib/Utility/Debugging"
	lustr "github.com/PlayerR9/lib_units/strings"
)

// DebugPrintTray is a helper function that prints the contents of a tray.
//
// Parameters:
//   - tray: The tray to print.
//
// Returns:
//   - []string: A slice of strings that represent the contents of the tray.
func DebugPrintTray[T any](tray ut.Trayer[T]) []string {
	if tray == nil {
		return nil
	}

	originalPos := tray.GetLeftDistance()

	tray.ArrowStart()

	var values []string

	for {
		val, err := tray.Read()
		if err != nil {
			break
		}

		str := lustr.GoStringOf(val)
		values = append(values, str)

		remaining := tray.Move(1)
		if remaining != 0 {
			break
		}
	}

	tray.ArrowStart()
	tray.Move(originalPos)

	return []string{
		strings.Join(values, " "),
		ub.PrintPointer(originalPos),
	}
}
