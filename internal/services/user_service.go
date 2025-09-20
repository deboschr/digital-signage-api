package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers() ([]dto.GetSummaryUserResDTO, error)
	GetUser(id uint) (dto.GetDetailUserResDTO, error)
	CreateUser(req dto.CreateUserReqDTO) (dto.GetSummaryUserResDTO, error)
	UpdateUser(req dto.UpdateUserReqDTO) (dto.GetSummaryUserResDTO, error)
	DeleteUser(id uint) error
	Authenticate(username string, password string) (dto.GetSummaryUserResDTO, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUsers() ([]dto.GetSummaryUserResDTO, error) {

	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.GetSummaryUserResDTO
	for _, u := range users {
		res = append(res, dto.GetSummaryUserResDTO{
			UserID:   u.UserID,
			Username: u.Username,
			Role:     u.Role,
		})
	}

	return res, nil
}

func (s *userService) GetUser(id uint) (dto.GetDetailUserResDTO, error) {

	user, err := s.repo.FindByID(id)
	
	if err != nil {
		return dto.GetDetailUserResDTO{}, err
	}

	var airport *dto.GetSummaryAirportResDTO
	if user.Airport != nil {
		airport = &dto.GetSummaryAirportResDTO{
			AirportID: user.Airport.AirportID,
			Name:      user.Airport.Name,
			Code:      user.Airport.Code,
			Address:   user.Airport.Address,
		}
	}

	return dto.GetDetailUserResDTO{
		UserID:    user.UserID,
		Username:  user.Username,
		Role:      user.Role,
		Airport:   airport,
	}, nil
}

func (s *userService) CreateUser(req dto.CreateUserReqDTO) (dto.GetSummaryUserResDTO, error) {
	
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.GetSummaryUserResDTO{}, err
	}

	user := models.User{
		AirportID: req.AirportID,
		Username:  req.Username,
		Password:  string(hashed),
		Role:      req.Role,
	}

	if err := s.repo.Create(&user); err != nil {
		return dto.GetSummaryUserResDTO{}, err
	}

	return dto.GetSummaryUserResDTO{
		UserID:    user.UserID,
		Username:  user.Username,
		Role:      user.Role,
	}, nil
}

func (s *userService) UpdateUser(req dto.UpdateUserReqDTO) (dto.GetSummaryUserResDTO, error) {
	
	user, err := s.repo.FindByID(req.UserID)
	if err != nil {
		return dto.GetSummaryUserResDTO{}, err
	}

	if req.AirportID != nil {
		user.AirportID = req.AirportID
	}
	if req.Username != nil {
		user.Username = *req.Username
	}
	if req.Role != nil {
		user.Role = *req.Role
	}
	if req.Password != nil && *req.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return dto.GetSummaryUserResDTO{}, err
		}
		user.Password = string(hashed)
	}

	if err := s.repo.Update(user); err != nil {
		return dto.GetSummaryUserResDTO{}, err
	}

	return dto.GetSummaryUserResDTO{
		UserID:    user.UserID,
		Username:  user.Username,
		Role:      user.Role,
	}, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

// AUTH
func (s *userService) Authenticate(username string, password string) (dto.GetSummaryUserResDTO, error) {
	
	user, err := s.repo.FindByUsername(username)
	
	if err != nil {
		return dto.GetSummaryUserResDTO{}, err
	}
	
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
   	return dto.GetSummaryUserResDTO{}, errors.New("invalid credentials")
	}


	return dto.GetSummaryUserResDTO{
		UserID:   user.UserID,
		AirportID:   user.AirportID,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}
