package model

type State string

const (
	DRAFT          State = "draft"
	ON_NEGOTIATION State = "on_negotiation"
	COMPLETED      State = "completed"
	REJECTED       State = "rejected"
	CANCELLED      State = "cancelled"
	ON_REVIEW      State = "on_review"
)

type Order struct {
	Id            string
	SellerId      string
	AssignedPPK   string
	State         State
	Total         float32
	Discounts     map[string]interface{}
	AdditionalFee map[string]interface{}
	OrderedAt     int64
	CreatedAt     int64
	CreatedBy     string
	UpdatedAt     *int64
	UpdatedBy     *string
	Version       int
}

type OrderDetail struct {
	OrderId          string
	OrderItemId      string
	ProductId        string
	ProductVariantId string
	OriginalPrice    float32
	LastPrice        float32
	Qty              float32
	ProductSnapshot  map[string]interface{}
}

type AddOn struct {
	OrderItemId    string
	AddOnVariantId string
	Type           string
	OriginalPrice  float32
	LastPrice      float32
	UpdatedAt      *int64
	UpdatedBy      *string
}

type Shipment struct {
	Id       string
	OrderId  string
	Provider string
	Service  string
	// Items []
	Origin               map[string]interface{}
	Recipient            map[string]interface{}
	State                string
	OriginalFee          float32
	LastFee              float32
	EstimatedArrivalDate int64
	ShippedAt            int64
	AWBNumber            *string
	AWBRetrievedAt       *int64
	Notes                *string
	AdditionalData       map[string]interface{}
	CreatedAt            int64
	CreatedBy            string
	UpdatedAt            *int64
	UpdatedBy            *string
}
