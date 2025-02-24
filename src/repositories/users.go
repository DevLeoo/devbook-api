package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, nick, email, pwd) values (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		// log.Fatal(err)
		return 0, nil
	}

	lastInsertedID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

func (repository Users) GetAll(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := repository.db.Query("SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? OR nick LIKE ?", nameOrNick, nameOrNick)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) GetByID(id uint64) (models.User, error) {

	lines, err := repository.db.Query("SELECT id, name, nick, email, createdAt FROM users WHERE id = ?", id)

	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err := lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}

	}

	return user, nil
}

func (repository Users) Update(id uint64, user models.User) error {

	statement, err := repository.db.Prepare("UPDATE users SET name = ?, nick = ? , email= ? WHERE id = ?")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (repository Users) Remove(id uint64) error {

	statement, err := repository.db.Prepare("DELETE from users WHERE id = ?")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository Users) SearchByEmail(email string) (models.User, error) {

	lines, err := repository.db.Query("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err := lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) Follow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare("INSERT ignore INTO followers (user_id, follower_id) VALUES (?,?)")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository Users) Unfollow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM followers where user_id=? AND follower_id=?")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository Users) GetFollowers(userId uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		SELECT u.id , u.name, u.nick, u.email, u.createdAt FROM users u inner join followers s on
		u.id = follower_id where s.user_id = ?
	`, userId)

	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (repository Users) GetFollowing(userId uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		SELECT u.id , u.name, u.nick, u.email, u.createdAt FROM users u inner join followers s on
		u.id = follower_id where s.follower_id = ?
	`, userId)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (repository Users) FindPassword(userId uint64) (string, error) {
	lines, err := repository.db.Query(`
		SELECT pwd from users where id = ?
	`, userId)

	if err != nil {
		return "", err
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err := lines.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil

}

func (repository Users) UpdatePassword(userId uint64, password string) error {
	statement, err := repository.db.Prepare(`
		UPDATE users set pwd = ? where id = ?
	`)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(userId, password); err != nil {
		return err
	}

	return nil
}
