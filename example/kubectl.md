
```
root@ubuntu2004-1:~# curl -s -X POST -d '{"command":"kubectl get pods --output=json"}' http://127.0.0.1:8080/command | jq . 
{
  "result": {
    "apiVersion": "v1",
    "items": [
      {
        "apiVersion": "v1",
        "kind": "Pod",
        "metadata": {
          "creationTimestamp": "2023-05-31T15:00:10Z",
          "generateName": "nfs-client-provisioner-5cdb9888c4-",
          "labels": {
            "app": "nfs-client-provisioner",
            "pod-template-hash": "5cdb9888c4"
          },
          "name": "nfs-client-provisioner-5cdb9888c4-kksrd",
          "namespace": "default",
          "ownerReferences": [
            {
              "apiVersion": "apps/v1",
              "blockOwnerDeletion": true,
              "controller": true,
              "kind": "ReplicaSet",
              "name": "nfs-client-provisioner-5cdb9888c4",
              "uid": "335a8e93-9635-489f-87e4-5003fb939fc1"
            }
          ],
          "resourceVersion": "530434",
          "uid": "a191c336-e6c7-4afb-a2ef-4900e65c396e"
        },
        "spec": {
          "containers": [
            {
              "env": [
                {
                  "name": "PROVISIONER_NAME",
                  "value": "k8s-sigs.io/nfs-subdir-external-provisioner"
                },
                {
                  "name": "NFS_SERVER",
                  "value": "192.168.40.232"
                },
                {
                  "name": "NFS_PATH",
                  "value": "/tmp/nfs"
                }
              ],
              "image": "eipwork/nfs-subdir-external-provisioner:v4.0.2",
              "imagePullPolicy": "IfNotPresent",
              "name": "nfs-client-provisioner",
              "resources": {},
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "volumeMounts": [
                {
                  "mountPath": "/persistentvolumes",
                  "name": "nfs-client-root"
                },
                {
                  "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                  "name": "kube-api-access-6c7h7",
                  "readOnly": true
                }
              ]
            }
          ],
          "dnsPolicy": "ClusterFirst",
          "enableServiceLinks": true,
          "nodeName": "ubuntu2004-1",
          "preemptionPolicy": "PreemptLowerPriority",
          "priority": 0,
          "restartPolicy": "Always",
          "schedulerName": "default-scheduler",
          "securityContext": {},
          "serviceAccount": "nfs-client-provisioner",
          "serviceAccountName": "nfs-client-provisioner",
          "terminationGracePeriodSeconds": 30,
          "tolerations": [
            {
              "effect": "NoExecute",
              "key": "node.kubernetes.io/not-ready",
              "operator": "Exists",
              "tolerationSeconds": 300
            },
            {
              "effect": "NoExecute",
              "key": "node.kubernetes.io/unreachable",
              "operator": "Exists",
              "tolerationSeconds": 300
            }
          ],
          "volumes": [
            {
              "name": "nfs-client-root",
              "nfs": {
                "path": "/tmp/nfs",
                "server": "192.168.40.232"
              }
            },
            {
              "name": "kube-api-access-6c7h7",
              "projected": {
                "defaultMode": 420,
                "sources": [
                  {
                    "serviceAccountToken": {
                      "expirationSeconds": 3607,
                      "path": "token"
                    }
                  },
                  {
                    "configMap": {
                      "items": [
                        {
                          "key": "ca.crt",
                          "path": "ca.crt"
                        }
                      ],
                      "name": "kube-root-ca.crt"
                    }
                  },
                  {
                    "downwardAPI": {
                      "items": [
                        {
                          "fieldRef": {
                            "apiVersion": "v1",
                            "fieldPath": "metadata.namespace"
                          },
                          "path": "namespace"
                        }
                      ]
                    }
                  }
                ]
              }
            }
          ]
        },
        "status": {
          "conditions": [
            {
              "lastProbeTime": null,
              "lastTransitionTime": "2023-06-03T06:05:40Z",
              "status": "True",
              "type": "Initialized"
            },
            {
              "lastProbeTime": null,
              "lastTransitionTime": "2023-06-03T06:05:40Z",
              "message": "containers with unready status: [nfs-client-provisioner]",
              "reason": "ContainersNotReady",
              "status": "False",
              "type": "Ready"
            },
            {
              "lastProbeTime": null,
              "lastTransitionTime": "2023-06-03T06:05:40Z",
              "message": "containers with unready status: [nfs-client-provisioner]",
              "reason": "ContainersNotReady",
              "status": "False",
              "type": "ContainersReady"
            },
            {
              "lastProbeTime": null,
              "lastTransitionTime": "2023-06-03T06:05:40Z",
              "status": "True",
              "type": "PodScheduled"
            }
          ],
          "containerStatuses": [
            {
              "image": "eipwork/nfs-subdir-external-provisioner:v4.0.2",
              "imageID": "",
              "lastState": {},
              "name": "nfs-client-provisioner",
              "ready": false,
              "restartCount": 0,
              "started": false,
              "state": {
                "waiting": {
                  "reason": "ContainerCreating"
                }
              }
            }
          ],
          "hostIP": "192.168.40.232",
          "phase": "Pending",
          "qosClass": "BestEffort",
          "startTime": "2023-06-03T06:05:40Z"
        }
      }
    ],
    "kind": "List",
    "metadata": {
      "resourceVersion": ""
    }
  },
  "exitcode": 0,
  "error": ""
}
```