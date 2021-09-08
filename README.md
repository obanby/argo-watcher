# Argo watcher
A simple cli to watch:
 - Argo's application's events
 - Repository sync history


## Prerequisites:

- Set the following environment variables
    ```shell
    export ARGO_SERVER_API='<server-endpoint>/api/v1'
    export ARGO_TOKEN='<your-argo-token>'
    ```

    Issue your argo token using:
    ```shell
    curl $ARGOCD_SERVER/api/v1/session -d $'{"username":"admin","password":"password"}'
    ```
    
    For argo in k8s, get password using (assuming you have jq installed):
    ```shell
    kubectl get secret argocd-initial-admin-secret -o json -n argocd | jq -r '.data.password' | base64 -d
    ```

- You must have a running kubernetes with a public api

    > 💡 I am using KIND (kubernetes in docker) to run local k8s cluster. You can utilize the bootstrap script
    > @local-k8s directory to bootstrap your cluster.

- Running Argo server

## Usage

To install the tool, use:

```shell
cd argo-watcher
go install .
```

To watch an application's event, run:
```shell
argo-watcher events <app-name>
```

![image](https://github.com/obanby/argo-watcher/blob/master/assets/events_example.png)

To watch a repo sync history, run: 
```shell
argo-watcher synchistory
```

![image](https://github.com/obanby/argo-watcher/blob/master/assets/repo-synch-history-example.png)

You can also use docker and run as a docker container, using:
```shell
docker build -t argo-watcher .

docker run \
--network host \
-e ARGO_SERVER_API=$ARGO_SERVER_API \
-e ARGO_TOKEN=$ARGO_TOKEN \
-it argo-watcher <command>
```

## Implementation notes

In order to facilitate network communication and mocking, I have created an ArgoClient and a MockClient. Both 
complies with ArgoServerClient interface. This facilitated mocking network requests without ever reaching out 
the network. Also added a consistent way of dealing with response body.

Speaking of response body: Since we want to watch the sync history stream in real time, we have to keep an open 
connection with the server through the streaming endpoint, which lead to using goroutines and channels. 
These channels facilitated communication and cancellation between goroutines.

With respect to the requirement
> "Collect all sync history that happens in realtime in some store"

I have implemented it only for git. For future implementations for additional stores (helm), I would parametrize the
cli command `synchistory`.

For `synchistory` we provided a Print function that extracts data and print it to console. We can change 
the behaviour to handle duplications and custom printing (similar to the `eventPrinter`).

## Improvement points

1- Add more table driven tests rather than having a shared 
global test data

2- Write a custom json.Unmarshall method, to handle situations when 
the server is returning null. Right now it is handled as a hack (`{"items": null}`), 
and in production that won't fly

3- Add cobra for the cli, only if we add two to three more functionalities 

4- Use a library for logging the result rather than that custom implementation. This would allow 
writing to different format based on sample configs

5- Support sync history for helm store

## References:

* [argo api docs](https://argoproj.github.io/argo-cd/developer-guide/api-docs/)
* [argo developer's guide](https://argo-cd.readthedocs.io/en/stable/developer-guide/)
* [kind](https://kind.sigs.k8s.io/)
