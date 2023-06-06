package auth

import (
	"crud/entity"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func generateSalt() string {
	// Generate random salt using cryptographic randomness
	salt := make([]byte, 16)
	rand.Read(salt)

	return hex.EncodeToString(salt)
}

func GenerateHash(password string) (string, string) {
	// Generate salt
	salt := generateSalt()

	// Combine password and salt
	passwordWithSalt := []byte(password + salt)

	// Hash the password + salt combination
	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)

	return string(hashedPassword), salt
}

func VerifyLogin(password string, admin *entity.Actor, salt string) error {
	// Combine password and salt
	passwordWithSalt := []byte(admin.Password + salt)

	// Hash the password + salt combination
	err := bcrypt.CompareHashAndPassword([]byte(password), passwordWithSalt)

	return err
}

func GenerateTokenJwt(admin *entity.Actor) string {
	// Inisialisasi klaim-klaim yang ingin Anda sertakan dalam token
	claims := jwt.MapClaims{
		"id":          admin.ID,
		"role_id":     admin.RoleID,
		"username":    admin.Username,
		"is_verified": admin.IsVerified,
		"is_active":   admin.IsActive,
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(time.Hour * 1).Unix(),
	}

	// Tandatangani token dengan kunci rahasia
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("ZDtQ+A8H/XtD/dWUMw2fgUQeU9u/w+01Z4Pkq4flwkI="))
	if err != nil {
		// Penanganan kesalahan
		fmt.Println(err)
	}

	// Gunakan signedToken seperti yang Anda butuhkan
	return signedToken
}

func VerifyToken(accessToken string) (jwt.MapClaims, error) {
	// Token yang diterima
	receivedToken := accessToken

	// Verifikasi token dengan kunci rahasia
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("ZDtQ+A8H/XtD/dWUMw2fgUQeU9u/w+01Z4Pkq4flwkI="), nil
	})
	if err != nil {
		// Penanganan kesalahan
		fmt.Println("error:", err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Token valid, akses klaim-klaim yang ada
		//fmt.Println(claims["id"], claims["role_id"], claims["username"], claims["is_verified"], claims["is_active"])
		return claims, nil
	} else {
		// Token tidak valid
		//fmt.Println("token not valid")
		return nil, nil
	}

}
