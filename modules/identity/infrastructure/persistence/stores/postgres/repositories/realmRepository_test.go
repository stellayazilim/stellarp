package repositories

import (
	"StellaRP/modules/identity/domain/realm"
	"StellaRP/modules/identity/domain/realm/ValueObject"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	. "github.com/onsi/gomega"
	"github.com/pashagolub/pgxmock/v4"
	"regexp"
	"testing"
)

func TestRealmStore_Create_ShouldFailStartTransaction(t *testing.T) {
	g := NewGomegaWithT(t)

	mock, err := pgxmock.NewConn()

	if err != nil {
		g.Fail("Fail creating mock connection")
	}

	r := realmStore{
		conn: mock,
	}

	mock.ExpectBegin().WillReturnError(errors.New("start transaction error"))

	err = r.Create(context.Background(), realm.NewRealm(ValueObject.NewRealmId(), "", nil))

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("start transaction error"))
}

func TestRealmStore_Create_ShouldFailExec(t *testing.T) {
	g := NewGomegaWithT(t)
	mock, _ := pgxmock.NewConn()

	r := realmStore{
		conn: mock,
	}
	nr := realm.NewRealm(ValueObject.NewRealmId(), "stella", nil)

	// Begin transaction
	mock.ExpectBegin()

	// Expect Exec with named parameters (@id, @name)
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO realms (id, name) VALUES (@id, @name)`)).
		WithArgs(pgx.NamedArgs{
			"id":   nr.GetId().GetValue(),
			"name": nr.GetName(),
		}).
		WillReturnError(errors.New("exec fail"))

	// Expect the rollback to succeed after exec fails
	mock.ExpectRollback().WillReturnError(nil)

	// Execute the Create method
	err := r.Create(context.Background(), nr)

	// We expect an error because the exec failed
	g.Expect(err).To(HaveOccurred())

	// Ensure the error is related to exec failure, not rollback failure
	g.Expect(err.Error()).NotTo(ContainSubstring("rollback failed"))
	g.Expect(err.Error()).To(ContainSubstring("exec fail"))

	// Ensure all expectations were met
	g.Expect(mock.ExpectationsWereMet()).To(Succeed())

}

func TestRealmStore_Create_ShouldFailRollback(t *testing.T) {
	g := NewGomegaWithT(t)
	mock, _ := pgxmock.NewConn()

	r := realmStore{
		conn: mock,
	}
	nr := realm.NewRealm(ValueObject.NewRealmId(), "stella", nil)

	// Begin transaction
	mock.ExpectBegin()

	// Expect Exec with named parameters (@id, @name)
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO realms (id, name) VALUES (@id, @name)`)).
		WithArgs(pgx.NamedArgs{
			"id":   nr.GetId().GetValue(),
			"name": nr.GetName(),
		}).
		WillReturnError(errors.New("exec fail"))

	// Expect the rollback to succeed after exec fails
	mock.ExpectRollback().WillReturnError(errors.New("rollback fail"))

	// Execute the Create method
	err := r.Create(context.Background(), nr)

	// We expect an error because the exec failed
	g.Expect(err).To(HaveOccurred())

	// Ensure the error is related to exec failure, not rollback failure
	g.Expect(err.Error()).To(ContainSubstring("rollback failed"))
	g.Expect(err.Error()).To(ContainSubstring("exec fail"))

	// Ensure all expectations were met
	g.Expect(mock.ExpectationsWereMet()).To(Succeed())

}

func TestRealmStore_Create_ShouldCommit(t *testing.T) {
	g := NewGomegaWithT(t)
	mock, _ := pgxmock.NewConn()

	r := realmStore{
		conn: mock,
	}
	nr := realm.NewRealm(ValueObject.NewRealmId(), "stella", nil)

	// Begin transaction
	mock.ExpectBegin()

	// Expect Exec with named parameters (@id, @name), returning success result
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO realms (id, name) VALUES (@id, @name)`)).
		WithArgs(pgx.NamedArgs{
			"id":   nr.GetId().GetValue(),
			"name": nr.GetName(),
		}).
		// Simulate a successful insert affecting 1 row
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	// Expect commit to succeed
	mock.ExpectCommit().WillReturnError(nil)

	// Execute the Create method
	err := r.Create(context.Background(), nr)

	// Expect no error since the exec and commit succeed
	g.Expect(err).To(BeNil())

	// Ensure all expectations were met
	g.Expect(mock.ExpectationsWereMet()).To(Succeed())
}
