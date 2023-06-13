package config

var (
	Port          string = ":3000"
	DBUrl         string = "postgresql://postgres:1111@localhost:5432/banana"
	AuthSvcUrl    string = ":50051"
	PostSvcUrl    string = ":50052"
	CommentSvcUrl string = ":50053"
	JWTSecretKey  string = "nothing ever goes as planned in this cursed world"
)
