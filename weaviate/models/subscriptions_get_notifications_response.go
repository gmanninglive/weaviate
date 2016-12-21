package models




import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// SubscriptionsGetNotificationsResponse subscriptions get notifications response
// swagger:model SubscriptionsGetNotificationsResponse
type SubscriptionsGetNotificationsResponse struct {

	// Identifies what kind of resource this is. Value: the fixed string "weave#subscriptionsGetNotificationsResponse".
	Kind *string `json:"kind,omitempty"`

	// Past client notifications.
	Notifications []*ClientNotification `json:"notifications"`
}

// Validate validates this subscriptions get notifications response
func (m *SubscriptionsGetNotificationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotifications(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionsGetNotificationsResponse) validateNotifications(formats strfmt.Registry) error {

	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	for i := 0; i < len(m.Notifications); i++ {

		if swag.IsZero(m.Notifications[i]) { // not required
			continue
		}

		if m.Notifications[i] != nil {

			if err := m.Notifications[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}