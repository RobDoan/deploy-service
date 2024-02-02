kubectl create ns test --dry-run -o yaml \
  | linkerd inject - \
  | kubectl apply -f -

helm install backend-a -n test --set ui.message='A backend' podinfo/podinfo
helm install backend-b -n test --set ui.message='B backend' podinfo/podinfo

kubectl -n test port-forward svc/backend-a-podinfo 9898

curl -sX GET -H localhost:9898/api/info

curl -sX GET -H "x-request-id: alternative" localhost:9898/api/info