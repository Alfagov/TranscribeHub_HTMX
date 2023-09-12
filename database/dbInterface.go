package database

import (
	"TranscribeHub_HTMX/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
)

type Dao interface {
	RegisterUser(user models.RegisterUser) error
	LoginUser(user models.LoginUser) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UsernameExists(username string) (bool, error)

	GetNews() ([]*models.News, error)

	GetDefaultCountersByUserId(userId string) (*models.SubscriptionCounter, error)
	GetTranscriptionSubscriptionByUserId(userId string) (*models.SubscriptionCounter, error)
	GetAssistSubscriptionByUserId(userId string) (*models.SubscriptionCounter, error)
	GetProjectSubscriptionByUserId(userId string) (*models.SubscriptionCounter, error)
}

type DaoImpl struct {
	*pgx.Conn
}

func NewDao(db *pgx.Conn) Dao {
	return &DaoImpl{db}
}

func (d *DaoImpl) GetDefaultCountersByUserId(userId string) (*models.SubscriptionCounter, error) {
	var subscriptionCounter models.SubscriptionCounter
	projectCounter := models.GetDefaultProjectCounter()
	assistCounter := models.GetDefaultAssistCounter()
	transcriptionCounter := models.GetDefaultTranscriptionCounter()
	sqlString := "SELECT subscriptions.name, user_subscriptions.current_usage_projects, subscriptions.max_projects, user_subscriptions.current_usage_assists, subscriptions.max_assists, user_subscriptions.current_usage_transcriptions, subscriptions.max_transcriptions FROM users INNER JOIN user_subscriptions ON users.id = user_subscriptions.user_id INNER JOIN subscriptions ON user_subscriptions.subscription_id = subscriptions.id WHERE users.id = $1;"
	err := d.QueryRow(context.Background(), sqlString, userId).Scan(&subscriptionCounter.Name, &projectCounter.Value, &projectCounter.Max, &assistCounter.Value, &assistCounter.Max, &transcriptionCounter.Value, &transcriptionCounter.Max)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	subscriptionCounter.Counters = []*models.Counter{projectCounter, assistCounter, transcriptionCounter}

	return &subscriptionCounter, nil
}

func (d *DaoImpl) GetTranscriptionSubscriptionByUserId(userId string) (*models.SubscriptionCounter, error) {
	var subscriptionCounter models.SubscriptionCounter
	counter := models.GetDefaultTranscriptionCounter()
	sqlString := `SELECT
		subscriptions.name,
		user_subscriptions.current_usage_transcriptions, 
		subscriptions.max_transcriptions
		FROM
	users
	INNER JOIN user_subscriptions ON users.id = user_subscriptions.user_id
	INNER JOIN subscriptions ON user_subscriptions.subscription_id = subscriptions.id
	WHERE users.id = $1;`
	err := d.QueryRow(context.Background(), sqlString, userId).Scan(&subscriptionCounter.Name, &counter.Value, &counter.Max)
	if err != nil {
		return nil, err
	}

	subscriptionCounter.Counters = append(subscriptionCounter.Counters, counter)
	return &subscriptionCounter, nil
}

func (d *DaoImpl) GetAssistSubscriptionByUserId(userId string) (*models.SubscriptionCounter, error) {
	var subscriptionCounter models.SubscriptionCounter
	counter := models.GetDefaultAssistCounter()
	sqlString := `SELECT
		subscriptions.name,
		user_subscriptions.current_usage_assists, 
		subscriptions.max_assists
		FROM
	users
	INNER JOIN user_subscriptions ON users.id = user_subscriptions.user_id
	INNER JOIN subscriptions ON user_subscriptions.subscription_id = subscriptions.id
	WHERE users.id = $1;`
	err := d.QueryRow(context.Background(), sqlString, userId).Scan(&subscriptionCounter.Name, &counter.Value, &counter.Max)
	if err != nil {
		return nil, err
	}

	subscriptionCounter.Counters = append(subscriptionCounter.Counters, counter)
	return &subscriptionCounter, nil
}

func (d *DaoImpl) GetProjectSubscriptionByUserId(userId string) (*models.SubscriptionCounter, error) {
	var subscriptionCounter models.SubscriptionCounter
	counter := models.GetDefaultProjectCounter()
	sqlString := `SELECT
		subscriptions.name,
		user_subscriptions.current_usage_projects, 
		subscriptions.max_projects
		FROM
	users
	INNER JOIN user_subscriptions ON users.id = user_subscriptions.user_id
	INNER JOIN subscriptions ON user_subscriptions.subscription_id = subscriptions.id
	WHERE users.id = $1;`
	err := d.QueryRow(context.Background(), sqlString, userId).Scan(&subscriptionCounter.Name, &counter.Value, &counter.Max)
	if err != nil {
		return nil, err
	}

	subscriptionCounter.Counters = append(subscriptionCounter.Counters, counter)
	return &subscriptionCounter, nil
}

func (d *DaoImpl) GetNews() ([]*models.News, error) {
	rows, err := d.Query(context.Background(), "SELECT id, title, svgname, date, text, downloadlink, downloadtext FROM news WHERE active = true;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models.News
	for rows.Next() {
		var nw models.News
		var svgTemp string
		err = rows.Scan(&nw.Id, &nw.Title, &svgTemp, &nw.Date, &nw.Text, &nw.DownloadLink, &nw.DownloadText)
		if err != nil {
			return nil, err
		}

		svg, ok := models.SvgMap[svgTemp]
		if !ok {
			return nil, errors.New("svg not found")
		}

		nw.SvgNotification = svg

		out = append(out, &nw)
	}

	return out, nil
}

func (d *DaoImpl) RegisterUser(user models.RegisterUser) error {
	_, err := d.Exec(context.Background(), "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4);", user.Id, user.Username, user.Email, user.Password)
	return err
}

func (d *DaoImpl) LoginUser(user models.LoginUser) (*models.User, error) {
	var out models.User
	err := d.QueryRow(context.Background(), "SELECT id, username, email FROM users WHERE email = $1 AND password = $2;", user.Email, user.Password).Scan(&out.Id, &out.Username, &out.Email)
	return &out, err
}

func (d *DaoImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := d.QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1;", email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	return &user, err
}

func (d *DaoImpl) UsernameExists(username string) (bool, error) {
	err := d.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1;", username).Scan(&username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
