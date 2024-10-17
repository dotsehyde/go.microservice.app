package jwt

import (
	"crypto"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	path string
}

func NewJWT() *JWT {
	return &JWT{
		path: "secrets/rsa/token_rsa",
	}
}

func (j *JWT) GenerateToken(ttl time.Duration, session any) (int64, string, error) {
	now := time.Now().UTC()
	expire := now.Add(ttl).UTC().Unix()
	// Create the Claims
	claims := jwt.MapClaims{
		"session": session,
		"exp":     expire,
		"iat":     now.UTC().Unix(),
		"nbf":     now.UTC().Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	sec, err := j.privateKey()
	if err != nil {
		return 0, "", fmt.Errorf("privateKey: %v", err)
	}
	tokenString, err := token.SignedString(sec)
	if err != nil {
		return 0, "", fmt.Errorf("token.SignedString: %v", err)
	}
	return expire, tokenString, nil
}
func (j *JWT) ValidateToken(token string) (jwt.MapClaims, error) {
	tok, err := jwt.Parse(
		token, func(jwtToken *jwt.Token) (any, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
			}
			pub, err := j.publicKey()
			if err != nil {
				return "", fmt.Errorf("publicKey: %v", err)
			}
			return pub, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if !tok.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not cast claims")
	}
	return claims, nil
}

func (j *JWT) GenerateKey() *JWT {
	sec, _ := j.privateKey()
	pub, _ := j.publicKey()
	if sec != nil && pub != nil {
		return j
	}
	prikeyArgs := []string{"genrsa", "-out", j.path, "4096"}
	pubkeyArgs := []string{"rsa", "-in", j.path, "-pubout", "-out", fmt.Sprintf("%s.pub", j.path)}

	cmd := exec.Command("openssl", prikeyArgs...) // pipe the commands output to the applications
	// standard output
	cmd.Stdout = os.Stdout
	// Run still runs the command and waits for completion
	// but the output is instantly piped to Stdout
	if err := cmd.Run(); err != nil {
		log.Panicln("could not run command private key: ", err)
	}
	cmd = exec.Command("openssl", pubkeyArgs...) // pipe the commands output to the applications
	// standard output
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Panicln("could not run command public key: ", err)
	}
	return j
}

func (j *JWT) Public() crypto.PublicKey {
	key, _ := j.publicKey()
	return key
}
func (j *JWT) privateKey() (crypto.PrivateKey, error) {
	keyByte, err := os.ReadFile(j.path)
	if err != nil {
		log.Println("could not read private key: ", err)
		return nil, err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyByte)
	if err != nil {
		log.Println("could not parse private key: ", err)
		return nil, err
	}
	return key, nil
}
func (j *JWT) publicKey() (crypto.PublicKey, error) {
	keyByte, err := os.ReadFile(fmt.Sprintf("%s.pub", j.path))
	if err != nil {
		log.Println("could not read public key: ", err)
		return nil, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyByte)
	if err != nil {
		log.Println("could not parse public key: ", err)
		return nil, err
	}
	return key, nil
}
