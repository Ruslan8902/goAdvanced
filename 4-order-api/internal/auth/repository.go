package auth

import (
	"orderApiStart/pkg/db"

	"gorm.io/gorm/clause"
)

type SessionRepository struct {
	Database *db.Db
}

func NewSessionRepository(database *db.Db) *SessionRepository {
	return &SessionRepository{
		Database: database,
	}
}

func (repo *SessionRepository) Create(session *Session) (*Session, error) {
	result := repo.Database.DB.Create(session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

func (repo *SessionRepository) Update(session *Session) (*Session, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

func (repo *SessionRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Session{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *SessionRepository) GetBySessionId(id string) (*Session, error) {
	var session Session
	result := repo.Database.DB.First(&session, "SessionID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}

func (repo *SessionRepository) GetByUserId(id string) (*Session, error) {
	var session Session
	result := repo.Database.DB.First(&session, "UserID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}
