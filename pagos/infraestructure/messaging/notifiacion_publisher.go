package messaging

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Notificacion struct {
	PedidoID uint   `json:"pedido_id"` 
	Mensaje  string `json:"mensaje"`   
}

type NotificacionPublisher struct {
	conn    *amqp.Connection 
	channel *amqp.Channel    
	queue   amqp.Queue       
}

func NewNotificacionPublisher() (*NotificacionPublisher, error) {
	conn, err := amqp.Dial("amqp://dvelazquez:laconia@54.163.6.194:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"notificaciones", 
		true,             
		false,            
		false,            
		false,            
		nil,              
	)
	if err != nil {
		return nil, err
	}

	return &NotificacionPublisher{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (p *NotificacionPublisher) Publish(notificacion Notificacion) error {
	body, err := json.Marshal(notificacion)
	if err != nil {
		return err
	}

	err = p.channel.Publish(
		"",            
		p.queue.Name,  
		false,         
		false,         
		amqp.Publishing{
			ContentType: "application/json", 
			Body:        body,               
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Notificación enviada: %+v", notificacion)
	return nil
}

func (p *NotificacionPublisher) Close() {
	p.channel.Close()
	p.conn.Close()
}