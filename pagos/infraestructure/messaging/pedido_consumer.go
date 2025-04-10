package messaging

import (
	"encoding/json"
	"fmt"
	"log"

	"api2/pagos/aplication/usecase"
	"api2/pagos/domain/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PedidoConsumer struct {
	PagoUseCase *usecase.CreatePagoUseCase
	conn        *amqp.Connection
	channel     *amqp.Channel
	queue       amqp.Queue
}

func NewPedidoConsumer(pagoUseCase *usecase.CreatePagoUseCase) (*PedidoConsumer, error) {
	conn, err := amqp.Dial("amqp://dvelazquez:laconia@54.163.6.194:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"pedidos",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &PedidoConsumer{
		PagoUseCase: pagoUseCase,
		conn:        conn,
		channel:     ch,
		queue:       q,
	}, nil
}

func (c *PedidoConsumer) StartConsuming() {
	msgs, err := c.channel.Consume(
		c.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error al consumir mensajes: %v", err)
	}

	go func() {
		for d := range msgs {
			var pedidoData map[string]interface{}
			err := json.Unmarshal(d.Body, &pedidoData)
			if err != nil {
				log.Printf("Error al deserializar el pedido: %v", err)
				continue
			}

			log.Printf("Pedido recibido: %+v", pedidoData)

			pedidoID, ok := pedidoData["id"].(float64)
			if !ok {
				log.Println("Error: El campo 'id' no es válido")
				continue
			}

			total, ok := pedidoData["total"].(float64)
			if !ok {
				log.Println("Error: El campo 'total' no es válido")
				continue
			}

			pago := models.Pago{
				PedidoID: uint(pedidoID),
				Monto:    total,
				Metodo:   "Tarjeta",
				Estado:   "Procesado",
			}

			err = c.PagoUseCase.Execute(&pago)
			if err != nil {
				log.Printf("Error al procesar pago: %v", err)
				continue
			}

			log.Printf("Pago procesado con éxito: %+v", pago)

			notificacionPublisher, err := NewNotificacionPublisher()
			if err != nil {
				log.Printf("Error al conectar con RabbitMQ para notificaciones: %v", err)
				continue
			}
			defer notificacionPublisher.Close()

			notificacion := Notificacion{
				PedidoID: uint(pedidoID),
				Mensaje:  "Pago completado para el pedido ID: " + fmt.Sprint(uint(pedidoID)),
			}

			err = notificacionPublisher.Publish(notificacion)
			if err != nil {
				log.Printf("Error al enviar notificación: %v", err)
			} else {
				log.Printf("Notificación enviada: %+v", notificacion)
			}

			pedidoData["estado"] = "completado"

			historialPublisher, err := NewHistorialPublisher()
			if err != nil {
				log.Printf("Error al conectar con RabbitMQ para historial: %v", err)
				continue
			}
			defer historialPublisher.Close()

			historialBody, err := json.Marshal(pedidoData)
			if err != nil {
				log.Printf("Error al serializar pedido para historial: %v", err)
				continue
			}

			err = historialPublisher.Publish(historialBody)
			if err != nil {
				log.Printf("Error al publicar en historial: %v", err)
			} else {
				log.Printf("✅ Pedido archivado en historial: %+v", pedidoData)
			}
		}
	}()

	log.Println("Esperando pedidos...")
	select {}
}

func (c *PedidoConsumer) Close() {
	c.channel.Close()
	c.conn.Close()
}

type HistorialPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewHistorialPublisher() (*HistorialPublisher, error) {
	conn, err := amqp.Dial("amqp://dvelazquez:laconia@54.163.6.194:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"historial",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &HistorialPublisher{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (p *HistorialPublisher) Publish(body []byte) error {
	return p.channel.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (p *HistorialPublisher) Close() {
	p.channel.Close()
	p.conn.Close()
}
