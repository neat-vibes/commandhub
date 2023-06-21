
```shell
root@ubuntu2004-1:~# df -Th
Filesystem     Type      Size  Used Avail Use% Mounted on
udev           devtmpfs  2.9G     0  2.9G   0% /dev
tmpfs          tmpfs     592M  4.7M  587M   1% /run
/dev/sda2      ext4      147G   23G  117G  16% /
tmpfs          tmpfs     2.9G     0  2.9G   0% /dev/shm
tmpfs          tmpfs     5.0M     0  5.0M   0% /run/lock
tmpfs          tmpfs     2.9G     0  2.9G   0% /sys/fs/cgroup
/dev/loop0     squashfs   64M   64M     0 100% /snap/core20/1879
/dev/loop1     squashfs   54M   54M     0 100% /snap/snapd/19122
/dev/loop4     squashfs   92M   92M     0 100% /snap/lxd/24061
/dev/loop2     squashfs   64M   64M     0 100% /snap/core20/1891
/dev/loop3     squashfs   54M   54M     0 100% /snap/snapd/19361
tmpfs          tmpfs     2.9G   28K  2.9G   1% /var/lib/ceph/osd/ceph-0
shm            tmpfs      64M     0   64M   0% /run/containerd/io.containerd.grpc.v1.cri/sandboxes/7a416ee5582d937edc4c26b625fe6cb1712ba4e7a019cfaefdc057f5dfc04176/shm
shm            tmpfs      64M     0   64M   0% /run/containerd/io.containerd.grpc.v1.cri/sandboxes/0ab3c3b0bd5af9d7df02fc21dc8aaac4026dd88ec139d04f1cfff844ca4c7bda/shm
shm            tmpfs      64M     0   64M   0% /run/containerd/io.containerd.grpc.v1.cri/sandboxes/8c1f309732575823cff14b715aa08352bbff36d4b031ecd89b4edceae9a8b544/shm
shm            tmpfs      64M     0   64M   0% /run/containerd/io.containerd.grpc.v1.cri/sandboxes/54b32c41416561e4e2a7997ff7522603883def7ac5210c385698ccf95bb0f9a0/shm
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/7a416ee5582d937edc4c26b625fe6cb1712ba4e7a019cfaefdc057f5dfc04176/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/8c1f309732575823cff14b715aa08352bbff36d4b031ecd89b4edceae9a8b544/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/0ab3c3b0bd5af9d7df02fc21dc8aaac4026dd88ec139d04f1cfff844ca4c7bda/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/54b32c41416561e4e2a7997ff7522603883def7ac5210c385698ccf95bb0f9a0/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/31a4aec478219e4ce68d2a35cdf5e577a7ee3730c9a40982b56d790f8d5c3562/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/00bfc1544944448bda004a93c0596a7c76456722cd71e91fc6064650172d4a5b/rootfs
tmpfs          tmpfs     5.7G   12K  5.7G   1% /var/lib/kubelet/pods/dea86cdb-63ca-4f41-9ff0-9f36a78ac6f3/volumes/kubernetes.io~projected/kube-api-access-zdbr2
tmpfs          tmpfs     5.7G   12K  5.7G   1% /var/lib/kubelet/pods/9519eb4e-f14f-49a2-b55e-ed22ceeac51c/volumes/kubernetes.io~projected/kube-api-access-5h7ps
tmpfs          tmpfs     5.7G   12K  5.7G   1% /var/lib/kubelet/pods/a191c336-e6c7-4afb-a2ef-4900e65c396e/volumes/kubernetes.io~projected/kube-api-access-6c7h7
tmpfs          tmpfs     170M   12K  170M   1% /var/lib/kubelet/pods/c1d73ab1-e63c-4ecc-9571-3aa148db945e/volumes/kubernetes.io~projected/kube-api-access-4h45g
tmpfs          tmpfs     5.7G   12K  5.7G   1% /var/lib/kubelet/pods/b53443f6-6369-4388-8aae-67c1ee213a72/volumes/kubernetes.io~projected/kube-api-access-qnf2f
tmpfs          tmpfs     170M   12K  170M   1% /var/lib/kubelet/pods/a47ad61c-92ed-428f-a4d5-9a7a880a5929/volumes/kubernetes.io~projected/kube-api-access-lbbwn
shm            tmpfs      64M     0   64M   0% /run/containerd/io.containerd.grpc.v1.cri/sandboxes/1f4497a155abcd4d080e201738459c75db9e68cf82ede863b747b2734a101526/shm
shm            tmpfs      64M     0   64M   0% /run/containerd/io.containerd.grpc.v1.cri/sandboxes/bf562887511eb491f106d9cda38b45bc0772ad06a8dc6441296ae361a3fe38ef/shm
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/bf562887511eb491f106d9cda38b45bc0772ad06a8dc6441296ae361a3fe38ef/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/1f4497a155abcd4d080e201738459c75db9e68cf82ede863b747b2734a101526/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/c18888798efe2c4caa47f84d0878bbd65dc8f5cd6a03de89827fbcaeee800fdb/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/3d048fa33d58e152d2c7300d491f3825537cb15438e043db892672c5881e9f65/rootfs
shm            tmpfs      64M     0   64M   0% /run/containerd/io.containerd.grpc.v1.cri/sandboxes/799e640a75a345246049d1a6ebd8bc118b1424ff953d876d279a585e38229596/shm
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/799e640a75a345246049d1a6ebd8bc118b1424ff953d876d279a585e38229596/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/30526d5dbae10e4345d768235ef07dd92457b598ef4d6af3c7bb5572ea4c9343/rootfs
shm            tmpfs      64M     0   64M   0% /run/containerd/io.containerd.grpc.v1.cri/sandboxes/812bd159d7e2d1c7d020fda63dec11442dbe6343157342bdc1b5880f29be1955/shm
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/812bd159d7e2d1c7d020fda63dec11442dbe6343157342bdc1b5880f29be1955/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/71f647f2937498c805ce051e4d90a44c05ec765a8a3ec1952edc1c10f2fd8bba/rootfs
shm            tmpfs      64M     0   64M   0% /run/containerd/io.containerd.grpc.v1.cri/sandboxes/3e55080346cccf39046373a30d6b5eecf95a2e87a44022f67e4eaa14c831d0cf/shm
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/3e55080346cccf39046373a30d6b5eecf95a2e87a44022f67e4eaa14c831d0cf/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/8b65aae9b7a3314a1fa007d7add3f9447410d1caa02a18f049f83534e55b8343/rootfs
tmpfs          tmpfs     592M     0  592M   0% /run/user/0
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/81bb1536bcbbf5a899a3779733596deaf74881939237af189b22525cd2ade465/rootfs
overlay        overlay   147G   23G  117G  16% /run/containerd/io.containerd.runtime.v2.task/k8s.io/5e7cc21f4d661cc8bc414d636565e4b59bcad84715f13c2fdfef35bcffbe19a6/rootfs



root@ubuntu2004-1:~# curl -s -X POST -d '{"command":"df -Th | awk '\''NR>1 {print $1, $2, $3, $4, $5, $6}'\'' | tail -n +2 | jq -R '\''split(\" \") | {filesystem: .[0], type: .[1], size: .[2], used: .[3], available: .[4], percentage_used: .[5]}'\'' | jq -s '.'"}' http://127.0.0.1:8080/command | jq .
{
  "result": [
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "592M",
      "used": "4.6M",
      "available": "587M",
      "percentage_used": "1%"
    },
    {
      "filesystem": "/dev/sda2",
      "type": "ext4",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "2.9G",
      "used": "0",
      "available": "2.9G",
      "percentage_used": "0%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "5.0M",
      "used": "0",
      "available": "5.0M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "2.9G",
      "used": "0",
      "available": "2.9G",
      "percentage_used": "0%"
    },
    {
      "filesystem": "/dev/loop0",
      "type": "squashfs",
      "size": "64M",
      "used": "64M",
      "available": "0",
      "percentage_used": "100%"
    },
    {
      "filesystem": "/dev/loop1",
      "type": "squashfs",
      "size": "54M",
      "used": "54M",
      "available": "0",
      "percentage_used": "100%"
    },
    {
      "filesystem": "/dev/loop4",
      "type": "squashfs",
      "size": "92M",
      "used": "92M",
      "available": "0",
      "percentage_used": "100%"
    },
    {
      "filesystem": "/dev/loop2",
      "type": "squashfs",
      "size": "64M",
      "used": "64M",
      "available": "0",
      "percentage_used": "100%"
    },
    {
      "filesystem": "/dev/loop3",
      "type": "squashfs",
      "size": "54M",
      "used": "54M",
      "available": "0",
      "percentage_used": "100%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "2.9G",
      "used": "28K",
      "available": "2.9G",
      "percentage_used": "1%"
    },
    {
      "filesystem": "shm",
      "type": "tmpfs",
      "size": "64M",
      "used": "0",
      "available": "64M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "shm",
      "type": "tmpfs",
      "size": "64M",
      "used": "0",
      "available": "64M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "shm",
      "type": "tmpfs",
      "size": "64M",
      "used": "0",
      "available": "64M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "shm",
      "type": "tmpfs",
      "size": "64M",
      "used": "0",
      "available": "64M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "5.7G",
      "used": "12K",
      "available": "5.7G",
      "percentage_used": "1%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "5.7G",
      "used": "12K",
      "available": "5.7G",
      "percentage_used": "1%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "5.7G",
      "used": "12K",
      "available": "5.7G",
      "percentage_used": "1%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "170M",
      "used": "12K",
      "available": "170M",
      "percentage_used": "1%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "5.7G",
      "used": "12K",
      "available": "5.7G",
      "percentage_used": "1%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "170M",
      "used": "12K",
      "available": "170M",
      "percentage_used": "1%"
    },
    {
      "filesystem": "shm",
      "type": "tmpfs",
      "size": "64M",
      "used": "0",
      "available": "64M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "shm",
      "type": "tmpfs",
      "size": "64M",
      "used": "0",
      "available": "64M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "shm",
      "type": "tmpfs",
      "size": "64M",
      "used": "0",
      "available": "64M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "shm",
      "type": "tmpfs",
      "size": "64M",
      "used": "0",
      "available": "64M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "shm",
      "type": "tmpfs",
      "size": "64M",
      "used": "0",
      "available": "64M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "tmpfs",
      "type": "tmpfs",
      "size": "592M",
      "used": "0",
      "available": "592M",
      "percentage_used": "0%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    },
    {
      "filesystem": "overlay",
      "type": "overlay",
      "size": "147G",
      "used": "23G",
      "available": "117G",
      "percentage_used": "16%"
    }
  ],
  "exitcode": 0,
  "error": ""
}
```