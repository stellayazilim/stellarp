package Entity

import "StellaRP/modules/identity/domain/realm/ValueObject"

type IRealmUser interface {
	IsEqual(other IRealmUser) bool
	GetId() ValueObject.IRealmUserId
	GetEmail() string
	SetEmail(email string)
	GetFirstName() string
	SetFirstName(firstName string)
	GetLastName() string
	SetLastName(lastName string)
	GetPassword() string
	SetPassword(password string)
}

type realmUser struct {
	id        ValueObject.IRealmUserId
	email     string
	firstName string
	lastName  string
	password  string
	realmId   string
}

func (r *realmUser) IsEqual(other IRealmUser) bool {
	id := r.GetId()
	return id.IsEqual(other.GetId())
}

func (r *realmUser) GetId() ValueObject.IRealmUserId {
	return r.id
}

func (r *realmUser) GetEmail() string {
	return r.email
}
func (r *realmUser) SetEmail(e string) {
	r.email = e
}

func (r *realmUser) GetFirstName() string {
	return r.firstName
}
func (r *realmUser) SetFirstName(n string) {
	r.firstName = n
}

func (r *realmUser) GetLastName() string {
	return r.lastName
}

func (r *realmUser) SetLastName(l string) {
	r.lastName = l
}

func (r *realmUser) GetPassword() string {
	return r.password
}

func (r *realmUser) SetPassword(p string) {
	r.password = p
}

func NewRealmUser(id ValueObject.IRealmUserId, email string, firstName string, lastName string, password string) IRealmUser {
	return &realmUser{
		id:        id,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
		password:  password,
	}
}
