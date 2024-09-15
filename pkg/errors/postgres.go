package errors

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jackc/pgconn"
)

// HandleError はエラーハンドリングを行います
func HandleDBError(err error) *Error {
	if err == nil {
		return nil
	}

	// sql.ErrNoRows の場合、401を返す
	if errors.Is(err, sql.ErrNoRows) {
		return New(http.StatusUnauthorized, err)
	}

	var pgErr *pgconn.PgError
	// PostgreSQLのエラーを検出
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // 重複キーエラー 409を返す
			return New(http.StatusConflict, err)
		default:
			return New(http.StatusInternalServerError, err)
		}
	}

	// それ以外は内部サーバーエラーを返す
	return New(http.StatusInternalServerError, err)
}
