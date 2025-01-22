package rabbitmq

type QueueOptions struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Arguments  map[string]interface{}
}

type QueueOption func(*QueueOptions)

func WithDurable(durable bool) QueueOption {
	return func(o *QueueOptions) {
		o.Durable = durable
	}
}

func WithAutoDelete(autoDelete bool) QueueOption {
	return func(o *QueueOptions) {
		o.AutoDelete = autoDelete
	}
}

func WithExclusive(exclusive bool) QueueOption {
	return func(o *QueueOptions) {
		o.Exclusive = exclusive
	}
}

func WithNoWait(noWait bool) QueueOption {
	return func(o *QueueOptions) {
		o.NoWait = noWait
	}
}

func WithArguments(arguments map[string]interface{}) QueueOption {
	return func(o *QueueOptions) {
		o.Arguments = arguments
	}
}

type PublishOptions struct {
	ContentType string
	Body        []byte
	Mandatory   bool
	Immediate   bool
	RoutingKey  string
}

type PublishOption func(*PublishOptions)

func WithContentType(contentType string) PublishOption {
	return func(o *PublishOptions) {
		o.ContentType = contentType
	}
}

func WithBody(body []byte) PublishOption {
	return func(o *PublishOptions) {
		o.Body = body
	}
}

func WithMandatory(mandatory bool) PublishOption {
	return func(o *PublishOptions) {
		o.Mandatory = mandatory
	}
}

func WithRoutingKey(routingKey string) PublishOption {
	return func(o *PublishOptions) {
		o.RoutingKey = routingKey
	}
}

// ConsumeOptions holds the options for consuming messages
type ConsumeOptions struct {
	AutoAck        bool
	Exclusive      bool
	NoLocal        bool
	NoWait         bool
	QueueArguments map[string]interface{}
}

// ConsumeOption is a function that modifies ConsumeOptions
type ConsumeOption func(*ConsumeOptions)

func WithAutoAck(autoAck bool) ConsumeOption {
	return func(o *ConsumeOptions) {
		o.AutoAck = autoAck
	}
}

func WithNoLocal(noLocal bool) ConsumeOption {
	return func(o *ConsumeOptions) {
		o.NoLocal = noLocal
	}
}

func WithQueueArguments(arguments map[string]interface{}) ConsumeOption {
	return func(o *ConsumeOptions) {
		o.QueueArguments = arguments
	}
}
