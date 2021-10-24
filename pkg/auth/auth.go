package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	AccountIDKey        = "account_id_key"
	AccountKey          = "account_key"
	VerificationDataKey = "verification_data_key"
)

const (
	UserRole    = 1
	AdminRole   = 2
	CompanyRole = 3
	CommonRole  = 0
)

type AuthHandler struct {
	Logger *zap.Logger
	Config *config.Config
}

func NewAuthHandler(logger *zap.Logger, cnf *config.Config) *AuthHandler {
	return &AuthHandler{
		Logger: logger,
		Config: cnf,
	}
}

// GenericResponse is the format of our response
type GenericResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Errors []string `json:"errors"`
}

// Below data types are used for encoding and decoding b/t go types and json
type TokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type AuthResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Username     string `json:"username"`
}

type UsernameUpdate struct {
	Username string `json:"username"`
}

type CodeVerificationReq struct {
	Code string `json:"code"`
	Type string `json:"type"`
}

type PasswordResetReq struct {
	Password   string `json:"password"`
	PasswordRe string `json:"password_re"`
	Code       string `json:"code"`
}

var ErrUserAlreadyExists = fmt.Sprintf("User already exists with the given email")
var ErrUserNotFound = fmt.Sprintf("No user account exists with given email. Please sign in first")
var UserCreationFailed = fmt.Sprintf("Unable to create user.Please try again later")

var PgDuplicateKeyMsg = "duplicate key value violates unique constraint"
var PgNoRowsMsg = "no rows in result set"

type Authentication interface {
	Authenticate(reqAccount *models.Account, account *models.Account) bool
	GenerateAccessToken(account *models.Account) (string, error)
	GenerateRefreshToken(account *models.Account) (string, error)
	GenerateCustomKey(accountID string, password string) string
	ValidateAccessToken(token string) (string, error)
	ValidateRefreshToken(token string) (string, string, error)
}

// RefreshTokenCustomClaims specifies the claims for refresh token
type RefreshTokenCustomClaims struct {
	AccountID    string `json:"account_id"`
	CustomKey    string `json:"custom_key"`
	KeyType      string `json:"key_type"`
	AccountInfor string `json:"account_infor"`
	jwt.StandardClaims
}

// AccessTokenCustomClaims specifies the claims for access token
type AccessTokenCustomClaims struct {
	AccountID    string `json:"account_id"`
	Email        string `json:"email"`
	KeyType      string `json:"key_type"`
	AccountInfor string `json:"account_infor"`
	jwt.StandardClaims
}

func (auth *AuthHandler) Authenticate(reqAccount *models.Account, account *models.Account) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(reqAccount.Password)); err != nil {
		auth.Logger.Error("Password not same", zap.Error(err))
		return false
	}
	return true
}

// GenerateRefreshToken generate a new refresh token for the given user
func (auth *AuthHandler) GenerateRefreshToken(account *models.Account) (string, error) {

	cusKey := auth.GenerateCustomKey(string(account.Id), account.TokenHash)
	tokenType := "refresh"
	jsonAccount, err := json.Marshal(account)
	if err != nil {
		auth.Logger.Error("Can not convert account to json", zap.Error(err))
		return "", err
	}
	claims := RefreshTokenCustomClaims{
		string(account.Id),
		cusKey,
		tokenType,
		string(jsonAccount),
		jwt.StandardClaims{
			Issuer: "job4e.auth.service",
		},
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(auth.Config.RefreshTokenPrivateKey))
	if err != nil {
		auth.Logger.Error("unable to parse private key", zap.Error(err))
		return "", errors.New("could not generate refresh token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

// GenerateAccessToken generates a new access token for the given user
func (auth *AuthHandler) GenerateAccessToken(account *models.Account) (string, error) {
	userID := string(account.Id)
	tokenType := "access"
	jsonAccount, err := json.Marshal(account)
	if err != nil {
		auth.Logger.Error("Can not convert account to json", zap.Error(err))
		return "", err
	}
	claims := AccessTokenCustomClaims{
		userID,
		account.Email,
		tokenType,
		string(jsonAccount),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(auth.Config.JwtExpiration)).Unix(),
			Issuer:    "job4e.auth.service",
		},
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(auth.Config.AccessTokenPrivateKey))
	if err != nil {
		auth.Logger.Error("unable to parse private key", zap.Error(err))
		return "", errors.New("could not generate access token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

// GenerateCustomKey creates a new key for our jwt payload
// the key is a hashed combination of the userID and user tokenhash
func (auth *AuthHandler) GenerateCustomKey(userID string, tokenHash string) string {

	// data := userID + tokenHash
	h := hmac.New(sha256.New, []byte(tokenHash))
	h.Write([]byte(userID))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

// ValidateAccessToken parses and validates the given access token
// returns the userId present in the token payload
func (auth *AuthHandler) ValidateAccessToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			auth.Logger.Error("Unexpected signing method in auth token")
			return nil, errors.New("Unexpected signing method in auth token")
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(auth.Config.AccessTokenPublicKey))
		if err != nil {
			auth.Logger.Error("unable to parse public key", zap.Error(err))
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		auth.Logger.Error("unable to parse claims", zap.Error(err))
		return "", err
	}

	claims, ok := token.Claims.(*AccessTokenCustomClaims)
	if !ok || !token.Valid || claims.AccountID == "" || claims.KeyType != "access" {
		return "", errors.New("invalid token: authentication failed")
	}
	return claims.AccountInfor, nil
}

// ValidateRefreshToken parses and validates the given refresh token
// returns the userId and customkey present in the token payload
func (auth *AuthHandler) ValidateRefreshToken(tokenString string) (string, string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			auth.Logger.Error("Unexpected signing method in auth token")
			return nil, errors.New("Unexpected signing method in auth token")
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(auth.Config.RefreshTokenPublicKey))
		if err != nil {
			auth.Logger.Error("unable to parse public key", zap.Error(err))
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		auth.Logger.Error("unable to parse claims", zap.Error(err))
		return "", "", err
	}

	claims, ok := token.Claims.(*RefreshTokenCustomClaims)

	if !ok || !token.Valid || claims.AccountID == "" || claims.KeyType != "refresh" {
		auth.Logger.Error("could not extract claims from token")
		return "", "", errors.New("invalid token: authentication failed")
	}
	return claims.AccountInfor, claims.CustomKey, nil
}
