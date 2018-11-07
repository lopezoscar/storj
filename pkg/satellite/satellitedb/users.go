package satellitedb

import "storj.io/satest"

type Users struct {
	db *dbx.DB
}

func (users *Users) Get(ctx context.Context, id uuid.UUID) (*satellite.User, error) {
	user, err :=  users.db.Get_User_By_Id(u.ctx, User_Id(id.String()))
	return convertUser(user), err
}

func (users *Users) GetByCredentials(ctx context.Context, password, email string) (*satellite.User, error) {
	user, err :=  users.db.Get_User_By_Email_And_PasswordHash(ctx,
		User_Email(email), User_PasswordHash(password))
	return convertUser(user), err
}

func (users *Users) Insert(ctx context.Context, user *satellite.User) (error) {
	return users.db.Create_User(u.ctx, 
		User_Id(user.ID.String()),
		User_FirstName(user.FirstName),
		User_LastName(user.LastName),
		User_Email(user.Email),
		User_PasswordHash(user.PasswordHash),
	)
}

func (users *Users) Delete(ctx context.Context, id uuid.UUID) (error) {
	_, err := users.db.Delete_User_By_Id(u.ctx, User_Id(id.String()))
	return err
}

func (users *Users) Update(ctx context.Context, user *satellite.User) (error) {
	userId := User_Id(user.Id().String())
	_, err := users.db.Update_User_By_Id(u.ctx, userId, 
		User_Update_Fields{
			User_FirstName(user.FirstName),
			User_LastName(user.LastName),
			User_Email(user.Email),
			User_PasswordHash(user.Password),
		})

	return err
}