package pkg

type EntityId interface {
	GetId() string
}
type RateLimiter interface {
	HasLimitReached(entityId EntityId) bool
}

type IpAddressEntityId struct {
	ip string
}

func NewIpAddressEntityId(ip string) *IpAddressEntityId {
	return &IpAddressEntityId{ip: ip}
}

func (r *IpAddressEntityId) GetId() string {
	return r.ip
}
