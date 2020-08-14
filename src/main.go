package main

import "fmt"
import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"


func main(){
	fmt.Println("hello")
	cm := kafka.ConfigMap{}
	cm.SetKey("bootstrap.servers", "localhost:9092")
	cm = nil
}
