package util

import (
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

func CreatePgTypeFloat(n int64) pgtype.Numeric {
	number := new(big.Int)
	number.SetInt64(n)
	pgNumber := pgtype.Numeric{
		Int:   number,
		Exp:   -2,
		Valid: true,
	}
	return pgNumber
}
