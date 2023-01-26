# Kubernetes Support

Kubernetes offers an api from the master node(s) which can be used to get some
internal information.

For accessing this api, we need an access token.

## AWS

To get the token via AWS IAM Authenticator, you can run:

    $ aws-iam-authenticator token -i CLUSTERID

This is a bearer token that can be used on our API side to work with the cluster.
