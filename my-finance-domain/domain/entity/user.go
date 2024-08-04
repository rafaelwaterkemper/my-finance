package domain

type User struct {
	id             int
	identification string
	email          string
	cpf            string
}

func (u User) GetId() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

type UserBuilder struct {
	user User
}

func (b UserBuilder) New() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) Id(id int) *UserBuilder {
	b.user.id = id
	return b
}

func (b *UserBuilder) Identification(identification string) *UserBuilder {
	b.user.identification = identification
	return b
}

func (b *UserBuilder) Email(email string) *UserBuilder {
	b.user.email = email
	return b
}

func (b *UserBuilder) Cpf(cpf string) *UserBuilder {
	b.user.cpf = cpf
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}
