# commandhub
Client-side Shell Command Execution over HTTP with Result Retrieval

We can return the output of shell commands that implement JSON formatting. 
At the moment, we do not support shell commands that do not provide JSON-formatted output.

### example
#### Native support for command output in json format

lsblk

```shell
curl -s -X POST -d '{"command":"lsblk -J -d -O -b"}' http://127.0.0.1:8080/command | jq .
```

ceph

```shell
curl -s -X POST -d '{"command":"ceph -s -f json"}' http://127.0.0.1:8080/command | jq . 
```

kubectl

```
curl -s -X POST -d '{"command":"kubectl get pods --output=json"}' http://127.0.0.1:8080/command | jq . 
```

#### Commands in json format are not supported, you can convert them through ChatGPT for example

df

```
curl -s -X POST -d '{"command":"df -Th | awk '\''NR>1 {print $1, $2, $3, $4, $5, $6}'\'' | tail -n +2 | jq -R '\''split(\" \") | {filesystem: .[0], type: .[1], size: .[2], used: .[3], available: .[4], percentage_used: .[5]}'\'' | jq -s '.'"}' http://127.0.0.1:8080/command | jq .
```