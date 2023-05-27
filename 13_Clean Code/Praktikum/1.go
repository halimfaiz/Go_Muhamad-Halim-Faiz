package main

// Dalam struct user :
// id lebih baik uint
//'username' dan 'password' umumnya teks jadi harus string tidak seharusnya integer
type user struct {
	id       int
	username int
	password int
}

//penamaan lebih baik camelcase 'userService'
type userservice struct {
	t []user
}

//penamaan lebih baik camelcase 'getAllUsers'
func (u userservice) getallusers() []user {
	return u.t
}

//penamaan lebih baik camelcase 'getUserById'
//lebih baik parameternya diganti 'id uint'
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
