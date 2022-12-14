package calc

import (
	context "context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	grpc "google.golang.org/grpc"
)

type server struct {}

func CalculateServer() {
	fmt.Println("Start gRPC greet server")
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterCalcServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func ( *server) Calc(ctx context.Context, req *CalcReq) (*CalcRes, error) {
	fmt.Println("Request num1: ", req.Calculate.Num1)
	fmt.Println("Request num2: ", req.Calculate.Num2)
	num1 := req.GetCalculate().Num1
	num2 := req.GetCalculate().Num2
	result := num1 + num2
	res := &CalcRes{
		Result: result,
	}
	return res, nil
}

func ( *server) CalcManyTimes(req *CalcManyTimesReq, stream CalcService_CalcManyTimesServer) error {
	fmt.Println("Request num1: ", req.Calculate.Num1)
	fmt.Println("Request num2: ", req.Calculate.Num2)
	num1 := req.GetCalculate().Num1
	num2 := req.GetCalculate().Num2
	result := num1 + num2
	for i := 0; i < 16; i++ {
		res := &CalcRes{
			Result: result,
		}
		result += result
		stream.SendMsg(res)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func ( *server) LongCalc(stream CalcService_LongCalcServer) error {
	fmt.Println("Starting to do a calculate client streaming")
	result := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&LongCalcRes{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Failed to receive response: %v", err)
		}
		num1 := req.GetCalculate().GetNum1()
		num2 := req.GetCalculate().GetNum2()
		result += num1 + num2
	}
}

func ( *server) ManyCalc(stream CalcService_ManyCalcServer) error {
	fmt.Println("Starting to do a many calculate bidirestional stream")

	var sum int32
	var result int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Failed to receive request: %v", err)
		}
		num1 := req.GetCalculate().GetNum1()
		num2 := req.GetCalculate().GetNum2()
		sum = num1 + num2
		result += sum
		sendErr := stream.Send(&ManyCalcRes{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Failed to send response: %v", sendErr)
		}
		log.Printf("result: %v", result)
	}
}

