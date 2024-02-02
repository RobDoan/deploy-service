#!/usr/bin/env bash


function install_linkerd() {
    echo "Installing Linkerd"
    linkerd version
    linkerd check --pre
    linkerd install --crds | kubectl apply -f -
    linkerd install --set proxyInit.runAsRoot=true | kubectl apply -f -
    linkerd viz install | kubectl apply -f -
    linkerd check
    linkerd dashboard &
}

function uninstall_linkerd(){
  echo "Uninstalling Linkerd"
  linkerd install --crds --ignore-cluster | kubectl delete -f -
  linkerd install --ignore-cluster | kubectl delete -f -
}

function main() {
  echo "Installing Linkerd $1"
  if [ "$1" == "uninstall" ]; then
    uninstall_linkerd
    return
  fi
  install_linkerd
}

main $1