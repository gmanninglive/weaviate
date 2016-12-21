package acl_entries




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/weaviate/models"
)

/*WeaveACLEntriesPatchOK Successful response

swagger:response weaveAclEntriesPatchOK
*/
type WeaveACLEntriesPatchOK struct {

	// In: body
	Payload *models.ACLEntry `json:"body,omitempty"`
}

// NewWeaveACLEntriesPatchOK creates WeaveACLEntriesPatchOK with default headers values
func NewWeaveACLEntriesPatchOK() *WeaveACLEntriesPatchOK {
	return &WeaveACLEntriesPatchOK{}
}

// WithPayload adds the payload to the weave Acl entries patch o k response
func (o *WeaveACLEntriesPatchOK) WithPayload(payload *models.ACLEntry) *WeaveACLEntriesPatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weave Acl entries patch o k response
func (o *WeaveACLEntriesPatchOK) SetPayload(payload *models.ACLEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaveACLEntriesPatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}