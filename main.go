package main

import (
	"fmt"
	"food_delivery_be/common"
	"food_delivery_be/component"
	"food_delivery_be/component/uploadprovider"
	"food_delivery_be/memcache"
	"food_delivery_be/middleware"
	"food_delivery_be/modules/restaurant/restauranttranspot/ginrestaurant"
	"food_delivery_be/modules/restaurantlike/transport/ginrestaurantlike"
	"food_delivery_be/modules/upload/uploadtransport/ginupload"
	"food_delivery_be/modules/user/userstorage"
	"food_delivery_be/modules/user/usertransport/ginuser"
	"food_delivery_be/pubsub/pblocal"
	"food_delivery_be/skio"
	"food_delivery_be/subscriber"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	jg "go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetLevel(logrus.ErrorLevel) // must use env for devops input 0 -> 6
	l.Info("Test")
	l.WithField("body", "body ne").Error("Error roi")


	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DatabaseConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect database - ", err)
	}
	db = db.Debug()
	fmt.Println("Connected to database")
	if err := runServices(db, s3Provider, secretKey); err != nil {
		log.Fatal(err)
	}
}

func runServices(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {
	appCtx := component.NewAppContent(db, upProvider, secretKey, pblocal.NewPubSub())

	userStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
	userCaching := memcache.NewUserCaching(memcache.NewCaching(), userStore)

	r := gin.Default()

	//subscriber.Setup(appCtx)

	//startSocketIOServer(r, appCtx)

	rtEngine := skio.NewEngine()
	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Fatalln(err)
	}

	if err := subscriber.NewEngine(appCtx, rtEngine).Start(); err != nil {
		log.Fatalln(err)
	}

	r.Use(middleware.Recover(appCtx))

	r.StaticFile("/demo", "./demo.html")

	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appCtx))

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx, userCaching), ginuser.GetProfile(appCtx))

	restaurants := v1.Group( "/restaurants", middleware.RequireAuth(appCtx, userCaching))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.GET("/:id/liked-users",  ginrestaurantlike.ListUserLikeRestaurant(appCtx))

		restaurants.POST("/:id/like", ginrestaurantlike.UserLikeRestaurant(appCtx))
		restaurants.DELETE("/:id/unlike",  ginrestaurantlike.UserUnLikeRestaurant(appCtx))
	}

	v1.GET("encode-uid", func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}

		var d reqData
		c.ShouldBind(&d)

		c.JSON(200, gin.H{
			"id": common.NewUID(uint32(d.RealId), d.DbType, 1),
		})
	})

	e, err := jg.NewExporter(jg.Options{
		AgentEndpoint: "localhost:6831",
		Process: jg.Process{ServiceName: "Food-Delivery"},
	})
	if err != nil {
		log.Println(err)
	}

	trace.RegisterExporter(e)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(1)})

	return http.ListenAndServe(
		":8080",
		&ochttp.Handler{
			Handler: r,
		})
	//return r.Run()
}

//
//func startSocketIOServer(engine *gin.Engine, appCtx component.AppContext) {
//	server := socketio.NewServer(&engineio.Options{
//		Transports: []transport.Transport{websocket.Default},
//	})
//
//	server.OnConnect("/", func(s socketio.Conn) error {
//		fmt.Println("Connected: ", s.ID(), " Ip: ", s.RemoteAddr())
//
//		return nil
//	})
//
//	server.OnError("/", func(s socketio.Conn, e error) {
//		fmt.Println("meet error ", e)
//	})
//
//	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
//		fmt.Println("closed: ", reason)
//	})
//
//	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
//		db := appCtx.GetMainDBConnection()
//		store := userstorage.NewSQLStore(db)
//
//		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
//
//		payload, err := tokenProvider.Validate(token)
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//
//		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//
//		if user.Status == 0 {
//			s.Emit("authentication_failed", "You has been banned/deleted")
//			s.Close()
//			return
//		}
//
//		user.Mask(false)
//
//		s.Emit("your_profile", user)
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
//		fmt.Println(msg)
//	})
//
//	type Person struct {
//		Name string `json:"name"`
//		Age  int    `json:"age"`
//	}
//
//	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) {
//		fmt.Println("Server receive notice: ", p.Name, p.Age)
//
//		p.Age = 33
//		s.Emit("notice", p)
//	})
//
//	go server.Serve()
//
//	engine.GET("/socket.io/*any", gin.WrapH(server))
//	engine.POST("/socket.io/*any", gin.WrapH(server))
//}
