
package main
// 主要是对go-kit 的一次尝试
import (
    "encoding/json"
    "log"
    "net/http"
 
    "golang.org/x/net/context"
 
    httptransport "github.com/go-kit/kit/transport/http"
)
 
func main() {
    ctx := context.Background()
    svc := stringService{}
 
    uppercaseHandler := httptransport.NewServer(
        ctx,
        makeUppercaseEndpoint(svc),
        decodeUppercaseRequest,
        encodeResponse,
    )
 
    countHandler := httptransport.NewServer(
        ctx,
        makeCountEndpoint(svc),
        decodeCountRequest,
        encodeResponse,
    )
 
    http.Handle("/uppercase", uppercaseHandler)
    http.Handle("/count", countHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request uppercaseRequest
    if err := json.NewDecoder(r.Body).Decode(&amp;request); err != nil {
        return nil, err
    }
    return request, nil
}
 
func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request countRequest
    if err := json.NewDecoder(r.Body).Decode(&amp;request); err != nil {
        return nil, err
    }
    return request, nil
}
 
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}