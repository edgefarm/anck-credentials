# edgefarm.network

EdgeFarm networking layer

## credsmanager

The credsmanager opens a port (default: 6000) and listens for incoming grpc client connections (see [pkg/apis/v1alpha1/config.proto](pkg/apis/v1alpha1/config.proto)).
It handles requests for nats credentails and returns them to the client.
