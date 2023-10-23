package database

import "go-server/internal/models"

// GetUserByEmail - get user by email address
func (db *DB) GetUserByEmail(email string) (*models.User, error) {

	stmt, err := db.Client.Prepare("SELECT * FROM users WHERE email = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)

	var user models.User
	if err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Position); err != nil {
		return nil, err
	}

	return &user, nil
}
