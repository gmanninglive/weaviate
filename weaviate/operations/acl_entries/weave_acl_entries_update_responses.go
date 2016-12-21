package acl_entries




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/weaviate/models"
)

/*WeaveACLEntriesUpdateOK Successful response

swagger:response weaveAclEntriesUpdateOK
*/
type WeaveACLEntriesUpdateOK struct {

	// In: body
	Payload *models.ACLEntry `json:"body,omitempty"`
}

// NewWeaveACLEntriesUpdateOK creates WeaveACLEntriesUpdateOK with default headers values
func NewWeaveACLEntriesUpdateOK() *WeaveACLEntriesUpdateOK {
	return &WeaveACLEntriesUpdateOK{}
}

// WithPayload adds the payload to the weave Acl entries update o k response
func (o *WeaveACLEntriesUpdateOK) WithPayload(payload *models.ACLEntry) *WeaveACLEntriesUpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weave Acl entries update o k response
func (o *WeaveACLEntriesUpdateOK) SetPayload(payload *models.ACLEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaveACLEntriesUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}