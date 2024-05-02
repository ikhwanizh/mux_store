package helper

import (
	"context"
)

func GetUserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value("ID").(int)
	return userID, ok
}
