package config

var (
	Port = ":3000"
	//DBUrl         = "postgres://maulerrr:esF25OkWzbga@ep-red-frog-609625.us-east-2.aws.neon.tech/banana"
	DBUrl         = "postgres://maulerrr:1111@localhost:5432/banana"
	AuthSvcUrl    = ":50051"
	PostSvcUrl    = ":50052"
	CommentSvcUrl = ":50053"
	JWTSecretKey  = "nothing ever goes as planned in this cursed world"
	AMPQUrl       = "amqp://guest:guest@0.0.0.0:5672/"
)
