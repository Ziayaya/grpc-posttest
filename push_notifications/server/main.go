package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/Ziayaya/grpc-posttest/push_notifications/proto" // Sesuaikan dengan path yang benar
	"google.golang.org/grpc"
)

type notifServer struct {
	pb.UnimplementedNotifServiceServer
}

type FruitInfo struct {
	Name       string `json:"name"`
	ID         int    `json:"id"`
	Family     string `json:"family"`
	Order      string `json:"order"`
	Genus      string `json:"genus"`
	Nutritions struct {
		Calories      float32 `json:"calories"`
		Fat           float32 `json:"fat"`
		Sugar         float32 `json:"sugar"`
		Carbohydrates float32 `json:"carbohydrates"`
		Protein       float32 `json:"protein"`
	} `json:"nutritions"`
}

type PoetryInfo struct {
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	Lines     []string `json:"lines"`
	LineCount string   `json:"linecount"`
}

func getFruitInfoFromAPI(fruitName string) (*FruitInfo, error) {
	url := "https://www.fruityvice.com/api/fruit/" + fruitName
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fruitInfo FruitInfo
	if err := json.NewDecoder(resp.Body).Decode(&fruitInfo); err != nil {
		return nil, err
	}

	return &fruitInfo, nil
}

func getPoetryInfoFromAPI(lineCount int32) ([]PoetryInfo, error) {
	url := fmt.Sprintf("https://poetrydb.org/linecount/%d", lineCount)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var poetryInfo []PoetryInfo
	if err := json.NewDecoder(resp.Body).Decode(&poetryInfo); err != nil {
		return nil, err
	}

	return poetryInfo, nil
}

func (s *notifServer) GetPoetryInfo(req *pb.PoetryRequest, stream pb.NotifService_GetPoetryInfoServer) error {
	if req.Linecount <= 0 {
		return fmt.Errorf("line count must be greater than 0")
	}

	poetryInfo, err := getPoetryInfoFromAPI(req.Linecount)
	if err != nil {
		log.Println("Error fetching poetry data from API:", err)
		return err
	}

	// Kirim data puisi ke client secara streaming
	for _, poem := range poetryInfo {
		poetryResponse := &pb.PoetryResponse{
			Title:     poem.Title,
			Author:    poem.Author,
			Lines:     poem.Lines,
			Linecount: poem.LineCount,
		}
		if err := stream.Send(poetryResponse); err != nil {
			log.Println("Error sending poetry data to client:", err)
			return err
		}
		time.Sleep(time.Second) // Simulasi pengambilan data puisi secara periodik
	}
	log.Println("Poem that have", req.Linecount, "lines send to user.")

	return nil
}

func (s *notifServer) GetFruitInfo(req *pb.FruitRequest, stream pb.NotifService_GetFruitInfoServer) error {
	fruitInfo, err := getFruitInfoFromAPI(req.Name)
	if err != nil {
		log.Println("Error fetching data from API:", err)
		return err
	}

	// Kirim data ke client secara streaming
	response := &pb.FruitResponse{
		Name: fruitInfo.Name,
		// Id:     int32(fruitInfo.ID),
		Family: fruitInfo.Family,
		Order:  fruitInfo.Order,
		Genus:  fruitInfo.Genus,
		Nutritions: &pb.NutritionInfo{
			Calories:      fruitInfo.Nutritions.Calories,
			Fat:           fruitInfo.Nutritions.Fat,
			Sugar:         fruitInfo.Nutritions.Sugar,
			Carbohydrates: fruitInfo.Nutritions.Carbohydrates,
			Protein:       fruitInfo.Nutritions.Protein,
		},
	}
	if err := stream.Send(response); err != nil {
		log.Println("Error sending initial data to client:", err)
		return err
	}
	log.Println("Data", req.Name, "send to user.")
	time.Sleep(time.Second)

	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotifServiceServer(grpcServer, &notifServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// func getFruitInfoFromAPI() (*FruitInfo, error) {
// 	resp, err := http.Get("https://www.fruityvice.com/api/fruit/all")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var fruitInfo FruitInfo
// 	if err := json.NewDecoder(resp.Body).Decode(&fruitInfo); err != nil {
// 		return nil, err
// 	}

// 	return &fruitInfo, nil
// }

// Gunakan go routine untuk mengirim data secara streaming
// go func() {
// 	// Kirim data ke client secara streaming
// 	for i := 0; i < 5; i++ { // Contoh, bisa disesuaikan dengan data yang diambil dari API
// 		response := &pb.FruitResponse{
// 			Name:   req.Name,
// 			Id:     int32(fruitInfo.ID),
// 			Family: fruitInfo.Family,
// 			Order:  fruitInfo.Order,
// 			Genus:  fruitInfo.Genus,
// 			Nutritions: &pb.NutritionInfo{
// 				Calories:      fruitInfo.Nutritions.Calories,
// 				Fat:           fruitInfo.Nutritions.Fat,
// 				Sugar:         fruitInfo.Nutritions.Sugar,
// 				Carbohydrates: fruitInfo.Nutritions.Carbohydrates,
// 				Protein:       fruitInfo.Nutritions.Protein,
// 			},
// 		}
// 		if err := stream.Send(response); err != nil {
// 			log.Println("Error sending data to client:", err)
// 			return
// 		}
// 		time.Sleep(time.Second) // Simulasi pengambilan data secara periodik
// 	}
// }()
