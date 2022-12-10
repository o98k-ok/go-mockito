package sqlmocked

import "github.com/DATA-DOG/go-sqlmock"

type Expect struct {
	Mocker sqlmock.Sqlmock
}

func (e Expect) ExecOK(match string) *sqlmock.ExpectedExec {
	return e.Mocker.ExpectExec(match).WillReturnError(nil).WillReturnResult(sqlmock.NewResult(1, 1))
}

func (e Expect) ExecOKWithRes(match string, id, rowCount int64) *sqlmock.ExpectedExec {
	return e.Mocker.ExpectExec(match).WillReturnError(nil).WillReturnResult(sqlmock.NewResult(id, rowCount))
}

func (e Expect) ExecFailed(match string, err error) *sqlmock.ExpectedExec {
	return e.Mocker.ExpectExec(match).WillReturnError(err)
}

func (e Expect) QueryOK(match string, row *sqlmock.Rows) *sqlmock.ExpectedQuery {
	return e.Mocker.ExpectQuery(match).WillReturnError(nil).WillReturnRows(row)
}

func (e Expect) QueryFailed(match string, err error) *sqlmock.ExpectedQuery {
	return e.Mocker.ExpectQuery(match).WillReturnError(err)
}
