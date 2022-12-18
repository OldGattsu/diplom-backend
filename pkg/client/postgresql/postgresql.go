package postgresql

import (
	"context"
	"fmt"
)

func newClient(ctx *context.Context, username, password, host, port, database string) {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)

}
