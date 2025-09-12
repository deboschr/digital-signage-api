package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]dto.SummaryUserDTO, error)
	GetUserByID(id uint) (dto.DetailUserDTO, error)
	CreateUser(req dto.CreateUserReqDTO) (dto.CreateUserResDTO, error)
	UpdateUser(req dto.UpdateUserReqDTO) (dto.UpdateUserResDTO, error)
	DeleteUser(id uint) error
	Authenticate(username, password string) (dto.SummaryUserDTO, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

// GET all → Summary DTO
func (s *userService) GetAllUsers() ([]dto.SummaryUserDTO, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.SummaryUserDTO
	for _, u := range users {
		res = append(res, dto.SummaryUserDTO{
			UserID:   u.UserID,
			Username: u.Username,
			Role:     u.Role,
		})
	}
	return res, nil
}

// GET by ID → Detail DTO
func (s *userService) GetUserByID(id uint) (dto.DetailUserDTO, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return dto.DetailUserDTO{}, err
	}

	var airport *dto.SummaryAirportDTO
	if user.Airport != nil {
		airport = &dto.SummaryAirportDTO{
			AirportID: user.Airport.AirportID,
			Name:      user.Airport.Name,
			Code:      user.Airport.Code,
			Address:   user.Airport.Address,
		}
	}

	return dto.DetailUserDTO{
		UserID:    user.UserID,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Airport:   airport,
	}, nil
}

// POST → Create DTO
func (s *userService) CreateUser(req dto.CreateUserReqDTO) (dto.CreateUserResDTO, error) {
	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.CreateUserResDTO{}, err
	}

	user := models.User{
		AirportID: req.AirportID,
		Username:  req.Username,
		Password:  string(hashed),
		Role:      req.Role,
	}

	if err := s.repo.Create(&user); err != nil {
		return dto.CreateUserResDTO{}, err
	}

	return dto.CreateUserResDTO{
		UserID:    user.UserID,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// PUT/PATCH → Update DTO
func (s *userService) UpdateUser(req dto.UpdateUserReqDTO) (dto.UpdateUserResDTO, error) {
	user, err := s.repo.FindByID(req.UserID)
	if err != nil {
		return dto.UpdateUserResDTO{}, err
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
			return dto.UpdateUserResDTO{}, err
		}
		user.Password = string(hashed)
	}

	if err := s.repo.Update(user); err != nil {
		return dto.UpdateUserResDTO{}, err
	}

	return dto.UpdateUserResDTO{
		UserID:    user.UserID,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// DELETE
func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

// AUTH
func (s *userService) Authenticate(username, password string) (dto.SummaryUserDTO, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return dto.SummaryUserDTO{}, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return dto.SummaryUserDTO{}, err
	}

	return dto.SummaryUserDTO{
		UserID:   user.UserID,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}
