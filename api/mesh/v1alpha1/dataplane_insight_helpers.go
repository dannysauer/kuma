package v1alpha1

import (
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/kumahq/kuma/api/helpers"
)

func NewSubscriptionStatus() *DiscoverySubscriptionStatus {
	return &DiscoverySubscriptionStatus{
		Total: &DiscoveryServiceStats{},
		Cds:   &DiscoveryServiceStats{},
		Eds:   &DiscoveryServiceStats{},
		Lds:   &DiscoveryServiceStats{},
		Rds:   &DiscoveryServiceStats{},
	}
}

func NewVersion() *Version {
	return &Version{
		KumaDp: &KumaDpVersion{
			Version:   "",
			GitTag:    "",
			GitCommit: "",
			BuildDate: "",
		},
		Envoy: &EnvoyVersion{
			Version: "",
			Build:   "",
		},
	}
}

func (x *DataplaneInsight) IsOnline() bool {
	for _, s := range x.GetSubscriptions() {
		if s.GetConnectTime() != nil && s.GetDisconnectTime() == nil {
			return true
		}
	}
	return false
}

func (x *DataplaneInsight) GetSubscription(id string) (int, *DiscoverySubscription) {
	for i, s := range x.GetSubscriptions() {
		if s.Id == id {
			return i, s
		}
	}
	return -1, nil
}

func (x *DataplaneInsight) UpdateCert(generation time.Time, expiration time.Time) error {
	if x.MTLS == nil {
		x.MTLS = &DataplaneInsight_MTLS{}
	}
	ts := timestamppb.New(expiration)
	if err := ts.CheckValid(); err != nil {
		return err
	}
	x.MTLS.CertificateExpirationTime = ts
	x.MTLS.CertificateRegenerations++
	ts = timestamppb.New(generation)
	if err := ts.CheckValid(); err != nil {
		return err
	}
	x.MTLS.LastCertificateRegeneration = ts
	return nil
}

func (x *DataplaneInsight) UpdateSubscription(s helpers.Subscription) {
	if x == nil {
		return
	}
	discoverySubscription, ok := s.(*DiscoverySubscription)
	if !ok {
		return
	}
	i, old := x.GetSubscription(discoverySubscription.Id)
	if old != nil {
		x.Subscriptions[i] = discoverySubscription
	} else {
		x.finalizeSubscriptions()
		x.Subscriptions = append(x.Subscriptions, discoverySubscription)
	}
}

// If Kuma CP was killed ungracefully then we can get a subscription without a DisconnectTime.
// Because of the way we process subscriptions the lack of DisconnectTime on old subscription
// will cause wrong status.
func (x *DataplaneInsight) finalizeSubscriptions() {
	now := timestamppb.Now()
	for _, subscription := range x.GetSubscriptions() {
		if subscription.DisconnectTime == nil {
			subscription.DisconnectTime = now
		}
	}
}

// todo(lobkovilya): delete GetLatestSubscription, use GetLastSubscription instead
func (x *DataplaneInsight) GetLatestSubscription() (*DiscoverySubscription, *time.Time) {
	if len(x.GetSubscriptions()) == 0 {
		return nil, nil
	}
	var idx int = 0
	var latest *time.Time
	for i, s := range x.GetSubscriptions() {
		if err := s.ConnectTime.CheckValid(); err != nil {
			continue
		}
		t := s.ConnectTime.AsTime()
		if latest == nil || latest.Before(t) {
			idx = i
			latest = &t
		}
	}
	return x.Subscriptions[idx], latest
}

func (x *DataplaneInsight) GetLastSubscription() helpers.Subscription {
	if len(x.GetSubscriptions()) == 0 {
		return nil
	}
	return x.GetSubscriptions()[len(x.GetSubscriptions())-1]
}

func (x *DiscoverySubscription) SetDisconnectTime(t time.Time) {
	x.DisconnectTime = timestamppb.New(t)
}

func (x *DataplaneInsight) Sum(v func(*DiscoverySubscription) uint64) uint64 {
	var result uint64 = 0
	for _, s := range x.GetSubscriptions() {
		result += v(s)
	}
	return result
}

func (s *DiscoverySubscriptionStatus) StatsOf(typeUrl string) *DiscoveryServiceStats {
	if s == nil {
		return &DiscoveryServiceStats{}
	}
	// we rely on type URL suffix to get rid of the dependency on concrete V2 / V3 implementation
	switch {
	case strings.HasSuffix(typeUrl, "Cluster"):
		if s.Cds == nil {
			s.Cds = &DiscoveryServiceStats{}
		}
		return s.Cds
	case strings.HasSuffix(typeUrl, "ClusterLoadAssignment"):
		if s.Eds == nil {
			s.Eds = &DiscoveryServiceStats{}
		}
		return s.Eds
	case strings.HasSuffix(typeUrl, "Listener"):
		if s.Lds == nil {
			s.Lds = &DiscoveryServiceStats{}
		}
		return s.Lds
	case strings.HasSuffix(typeUrl, "RouteConfiguration"):
		if s.Rds == nil {
			s.Rds = &DiscoveryServiceStats{}
		}
		return s.Rds
	default:
		return &DiscoveryServiceStats{}
	}
}
