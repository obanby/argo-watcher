# Argo watcher
A simple cli to watch an argo application's events, and a repository sync history.

## Usage 

In order to run the cli successfully you must have these prerequisites first:

1- You must set the following environment variables
```
 export ARGOSERVER_API='<server-endpoint>/api/v1'
 export ARGO_TOKEN='<your-argo-token>'
```

to issue your argo token you should use
```shell
    curl $ARGOCD_SERVER/api/v1/session -d $'{"username":"admin","password":"password"}'
```

To get the password if argo is running on k8s (assuming you have jq installed)

```shell
  kubectl get secret argocd-initial-admin-secret -o json -n argocd | jq -r '.data.password' | base64 -d
```

2- You must have a running kubernetes with a public api

3- Running Argo server

> 💡 I am using KIND (kubernetes in docker) to run local k8s cluster. You can utilize the bootstrap script @
> local-k8s directory to bootstrap your cluster.

To watch an application's event, run: `argo-watcher events <app-name>`

![image](https://github.com/obanby/argo-watcher/blob/master/assets/events_example.png)

To watch a repo sync history run: `argo-watcher synchistory`

![image](https://github.com/obanby/argo-watcher/blob/master/assets/repo-synch-history-example.png)

## Implementation notes

In order to facilitate network communication and mocking, I have created an ArgoClient and a MockClient. Both 
complies with ArgoServerClient interface. This facilitated mocking network requests without ever reaching out 
the network. Also added a consistent way of dealing with response body.

Speaking of response body: Since we want to watch the sync stream in real time, we have to keep an open connection
with the server through the streaming endpoint, which lead to using goroutines and channels. These channels facilitated
communication and cancellation between gotoutines.

## Improvement points

1- Code clean up is usually endless. However, I would like to add more table driven tests rather than having a shared 
global test data

2- I would probably write a custom json.Unmarshall method, since we want to be able to cleanly handle situations when 
the server is returning the value as null. Right now it is handled as a hack, and in production that won't fly

3- I would probably add cobra for the cli, only if we add two to three more functionalities 

4- I would also probably use a library for logging the result rather than that custom implementation. This would allow 
writing to different format based on sample configs

## References:

* [argo api docs](https://argoproj.github.io/argo-cd/developer-guide/api-docs/)
* [argo developer's guide](https://argo-cd.readthedocs.io/en/stable/developer-guide/)
* [kind](https://kind.sigs.k8s.io/)
