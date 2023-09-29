package authentication

import (
	"log"
	"os"
	"randomize721/go-fiber-server/utils/responses"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var authenticatedCredential = new(User)

func Login(c *fiber.Ctx) error {
	authenticatedCredential.Username = os.Getenv("AUTH_USERNAME")
	authenticatedCredential.Password = os.Getenv("AUTH_PASSWORD")

	loginCredential := new(User)

	if err := c.BodyParser(loginCredential); err != nil {
		response := responses.Default{
			ResponseCode:    400,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(400).JSON(response)
	}

	if loginCredential.Username != authenticatedCredential.Username {
		message := "Autentikasi gagal, username tidak dikenal."
		response := responses.Default{
			ResponseCode:    401,
			ResponseMessage: message,
		}
		log.Println(message)
		return c.Status(401).JSON(response)
	}

	if loginCredential.Password != authenticatedCredential.Password {
		message := "Autentikasi gagal, password salah."
		response := responses.Default{
			ResponseCode:    401,
			ResponseMessage: message,
		}
		log.Println(message)
		return c.Status(401).JSON(response)
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
	expirationTime := time.Now().Add(5 * time.Minute)

	jwtClaim := &Claim{
		Username: loginCredential.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		response := responses.Default{
			ResponseCode:    500,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(500).JSON(response)
	}

	jwtToken := JwtResponse{
		JwtToken: tokenString,
	}

	response := responses.Data{
		ResponseCode:    200,
		ResponseMessage: "Berhasil login.",
		Data:            jwtToken,
	}
	return c.Status(200).JSON(response)
}

func ValidateJwt(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	authorizationHeader := headers["Authorization"]
	if authorizationHeader == "" {
		message := "Token JWT tidak ditemukan."
		response := responses.Default{
			ResponseCode:    400,
			ResponseMessage: message,
		}
		log.Println(message)
		return c.Status(400).JSON(response)
	}
	tokenString := strings.Split(authorizationHeader, " ")[1]

	claim := &Claim{}
	jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			response := responses.Default{
				ResponseCode:    401,
				ResponseMessage: err.Error(),
			}
			log.Println(err.Error())
			return c.Status(401).JSON(response)
		}
		response := responses.Default{
			ResponseCode:    400,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(400).JSON(response)
	}
	if !token.Valid {
		response := responses.Default{
			ResponseCode:    401,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(401).JSON(response)
	}

	return c.Status(200).JSON(claim)
}
