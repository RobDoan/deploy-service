kubectl create ns uat --dry-run=client -o yaml | linkerd inject -  | kubectl apply -f -


helm upgrade --install backend -n uat --set ui.message='UAT backend' podinfo/podinfo

helm install frontend -n uat --set backend=http://backend-podinfo:9898/env podinfo/podinfo

kubectl -n uat port-forward svc/frontend-podinfo 9898 &

# get list ofr forwarding ports
lsof -i -P -n | grep LISTEN | grep 9898

kubectl create ns jira-123 --dry-run=client -o yaml \
  | linkerd inject - \
  | kubectl apply -f -

helm upgrade --install backend -n jira-123 --set ui.message='jira-123 backend' --set httproute.enabled=true ./podinfo

curl -sX GET -H "x-request-id: jira-124" localhost:9898/echo
curl -sX POST -H "x-request-id: jira-123" localhost:9898/echo


# install-podinfo.sh use go script
helm upgrade --install frontend -n podinfo-uat --set backend=http://backend-podinfo:9898/env podinfo/podinfo
go run cmd/deploy/*.go -chart podinfo/podinfo  -name backend -service podinfo -uat ./assets/http-router.yml
go run cmd/deploy/*.go -jira jira-123 -chart podinfo/podinfo  -name backend -service podinfo ./assets/http-router.yml
go run cmd/deploy/*.go -jira jira-124 -chart podinfo/podinfo  -name backend -service podinfo ./assets/http-router.yml

kubectl -n podinfo-uat port-forward svc/frontend-podinfo 9898 &