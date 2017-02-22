package easy_middleware

import (
	"context"
	"database/sql"
	"net/http"
)

const dbContextKey = "db"

// DatabaseError represenst a error for bad database connection,
// it's occurs while a request process
type DatabaseError struct {
	s string
}

func (e *DatabaseError) Error() string {
	return e.s
}

// SqlDb is a middlware for creating a unique session for each incomming request
func SqlDb(databaseDriver string, dsn string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			db, err := Open(databaseDriver, dsn)
			if err != nil {
				r = r.WithContext(context.WithValue(r.Context(), dbContextKey, &DatabaseError{err.Error()}))
			} else {
				defer Close(db)
				r = r.WithContext(context.WithValue(r.Context(), dbContextKey, db))
			}

			h.ServeHTTP(rw, r)
		})
	}
}

var Open = open

func open(driver, dsn string) (*sql.DB, error) {
	return sql.Open(driver, dsn)
}

var Close = close

func close(Db *sql.DB) error {
	return Db.Close()
}
