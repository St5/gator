package commands

import "context"

func CallbackReset(state State, params ...string) error {

	err := state.Db.ClearUesrs(context.Background())

	if err != nil {
		return err
	}

	return nil
}