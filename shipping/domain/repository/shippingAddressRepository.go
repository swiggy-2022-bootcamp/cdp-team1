package repository

//ShippingAddressRepository ..
type ShippingAddressRepository interface {
	DBHealthCheck() bool
}
