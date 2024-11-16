package clients

// import (
// 	"context"
// 	"log"
// 	"time"

// 	user "github.com/Yux77Yux/platform_backend/generated/user"
// 	"google.golang.org/grpc"
// )

func main() {
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("did not connect: %w", err)
	// }
	// defer conn.Close()

	// client := user.NewUserServiceClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// response, err := client.(ctx, &pb.HelloRequest{Name: "World"})
	// if err != nil {
	// 	log.Fatalf("could not greet: %w", err)
	// }

	// log.Printf("info: greeting: %s", response.Message)
}
