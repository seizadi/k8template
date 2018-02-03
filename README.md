# Kubernetes Templates

Notes on requirements

- There are use cases where a more complex template solution might augment
what we are trying to do here. One use case is where we use a standard
solution like NGINX Ingress Controller but just want to customize it.

Here is a case where you might have some internal yaml that was copied from
original NGINX site and modified. You end up with version out of date with
the latest stabel release:

```yaml
  name: default-http-backend
  labels:
    app: default-http-backend
  namespace: ingress-nginx
spec:
...
        image: gcr.io/google_containers/defaultbackend:1.4
```      
versus
```yaml
      - name: default-http-backend
        image: gcr.io/google_containers/defaultbackend:1.0
```
That was for the default backend and the Ingress controller
```yaml
      containers:
        - name: nginx-ingress-controller
          image: quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.10.2
```
versus
```yaml
      containers:
      - image: gcr.io/google_containers/nginx-ingress-controller:0.9.0-beta.14
        name: ingress-nginx
```
It would be tempting to just change just the image versions but as the
code changes much about the parameters like configMap or other
annotations baked into the older manifest need to change.

If we are not going to use Helm then it would make sense to fork the
original deploy project in the above example for 
[NGIX it is here](https://github.com/kubernetes/ingress-nginx/tree/master/deploy).

Then apply the changes as templetes on top of the vendor solution
this way we know exactly what changed.
