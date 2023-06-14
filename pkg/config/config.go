package config

var (
	Port          = ":3000"
	DBUrl         = "postgres://postgres:1111@localhost:5432/banana?sslmode=disable"
	AuthSvcUrl    = ":50051"
	PostSvcUrl    = ":50052"
	CommentSvcUrl = ":50053"
	JWTSecretKey  = "nothing ever goes as planned in this cursed world"
	AMPQUrl       = "amqp://guest:guest@localhost:5672/"
)
