package send

import "context"

type Presenter interface {
	Present(ctx context.Context, response *Response) error
}