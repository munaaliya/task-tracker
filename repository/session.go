package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"time"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	filebasedDb *filebased.Data
}

func NewSessionsRepo(filebasedDb *filebased.Data) *sessionsRepo {
	return &sessionsRepo{filebasedDb}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	return u.filebasedDb.AddSession(session) // TODO: replace this
}

func (u *sessionsRepo) DeleteSession(token string) error {
	return u.filebasedDb.DeleteSession(token) // TODO: replace this
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	return u.filebasedDb.UpdateSession(session)
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	sessionEmail, err := u.filebasedDb.SessionAvailEmail(email)

	if err != nil {
		return model.Session{}, err
	}

	return sessionEmail, nil
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	sessionToken, err := u.filebasedDb.SessionAvailToken(token)

	if err != nil {
		return model.Session{}, err
	}

	return sessionToken, nil
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
