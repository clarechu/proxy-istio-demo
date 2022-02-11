package pkg

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

type KafkaHandler struct {
	conn *kafka.Conn
}

func NewKafka(address string) *KafkaHandler {
	// to produce messages
	topic := "my-topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		panic(err)
	}
	return &KafkaHandler{
		conn: conn,
	}
}

func (d *KafkaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	log.Printf("get message data--> %v", message)
	err := d.Produce(message)
	if err != nil {
		log.Printf("%+v", err)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (d *KafkaHandler) Produce(message string) (err error) {

	//d.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = d.conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)
	if err != nil {
		log.Printf("failed to write messages: %s", err)
		return err
	}
	return err
}
