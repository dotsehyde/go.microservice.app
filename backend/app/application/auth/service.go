package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/gateways"
	"github.com/BenMeredithConsult/locagri-apps/app/adapters/presenters"
	"github.com/BenMeredithConsult/locagri-apps/app/application"
	"github.com/BenMeredithConsult/locagri-apps/app/application/producer"
	"github.com/BenMeredithConsult/locagri-apps/app/domain"
	"github.com/BenMeredithConsult/locagri-apps/app/domain/requestdto"
	"github.com/BenMeredithConsult/locagri-apps/utils/jwt"
)

const (
	ACCESS_TOKEN_EXPIRY  = time.Second * 60 * 60 * 24       // 1 day
	REFRESH_TOKEN_EXPIRY = time.Second * 60 * 60 * 24 * 366 // 1year
)

type service struct {
	repo     gateways.AuthRepo
	jwt      *jwt.JWT
	cache    gateways.CacheService
	producer gateways.EventProducer
}

func NewAuthService(repo gateways.AuthRepo, jwt *jwt.JWT, cacheSrv gateways.CacheService, producer gateways.EventProducer) gateways.AuthService {
	return &service{
		repo:     repo,
		jwt:      jwt,
		cache:    cacheSrv,
		producer: producer,
	}
}

func (s *service) VerifyOTP(phone string, otp string) (*presenters.UserAuthPresenter, error) {
	var correctCode string
	err := s.cache.Get(phone, &correctCode)
	if err != nil {
		return nil, fmt.Errorf("OTP expired.")
	}
	if correctCode != otp {
		return nil, fmt.Errorf("Invalid OTP.")
	}
	user, err := s.repo.SelectByPhone(phone)
	session := &presenters.AuthSession{
		ID:    user.ID,
		Phone: user.Phone,
		Name:  user.FirstName + " " + user.LastName,
		Photo: user.ProfilePhoto,
	}
	expiry, token, err := s.jwt.GenerateToken(ACCESS_TOKEN_EXPIRY, session)
	if err != nil {
		return nil, err
	}
	rExpiry, rToken, err := s.jwt.GenerateToken(REFRESH_TOKEN_EXPIRY, strconv.Itoa(session.ID))
	if err != nil {
		return nil, err
	}
	return &presenters.UserAuthPresenter{
		User:               presenters.NewAuthUser(user),
		Token:              token,
		TokenExpiry:        expiry,
		RefreshToken:       rToken,
		RefreshTokenExpiry: rExpiry,
	}, nil
}

func (s *service) SendOTP(phone string) (string, error) {
	if ok, _ := s.repo.ExistByPhone(phone); !ok {
		return "", fmt.Errorf("User not found")
	}
	otpCode := application.OTP(6)
	s.producer.Queue(producer.TopicSMSNotification, domain.SMSPayload{
		Recipients: []string{phone},
		Message:    fmt.Sprintf("Your OTP is %s", otpCode),
	})
	go func() {
		s.cache.Set(phone, "000000", time.Minute*10)
	}()
	return "OTP sent successfully", nil
}

func (s *service) ResetAccount(phone string) (string, error) {
	user, err := s.repo.SelectByPhone(phone)
	if err != nil {
		return "", err
	}
	return s.SendOTP(user.Phone)
}

func (s *service) Login(phone string) (string, error) {
	user, err := s.repo.SelectByPhone(phone)
	if err != nil {
		return "", err
	}
	return s.SendOTP(user.Phone)
}

func (s *service) Register(req *requestdto.CreateUserRequest) (string, error) {
	user, err := s.repo.Insert(req)
	if err != nil {
		return "", err
	}
	return s.SendOTP(user.Phone)
}

func (s *service) RefreshToken(refreshToken string) (*presenters.UserAuthPresenter, error) {
	claims, err := s.jwt.ValidateToken(refreshToken)
	if err != nil {
		return nil, err
	}
	sessionId, _ := strconv.Atoi(claims["session"].(string))
	user, err := s.repo.SelectById(sessionId)
	session := &presenters.AuthSession{
		ID:    user.ID,
		Phone: user.Phone,
		Name:  user.FirstName + " " + user.LastName,
		Photo: user.ProfilePhoto,
	}
	expiry, token, err := s.jwt.GenerateToken(ACCESS_TOKEN_EXPIRY, session)
	if err != nil {
		return nil, err
	}
	rExpiry, rToken, err := s.jwt.GenerateToken(REFRESH_TOKEN_EXPIRY, strconv.Itoa(session.ID))
	if err != nil {
		return nil, err
	}
	return &presenters.UserAuthPresenter{
		User:               presenters.NewAuthUser(user),
		Token:              token,
		TokenExpiry:        expiry,
		RefreshToken:       rToken,
		RefreshTokenExpiry: rExpiry,
	}, nil
}
