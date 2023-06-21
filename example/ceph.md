
```shell
root@ubuntu2004-1:~# curl -s -X POST -d '{"command":"ceph -s -f json"}' http://127.0.0.1:8080/command | jq . 
{
  "result": {
    "fsid": "0d5adb2f-77a2-4b29-87d6-7b7b1c726c71",
    "health": {
      "status": "HEALTH_WARN",
      "checks": {
        "AUTH_INSECURE_GLOBAL_ID_RECLAIM_ALLOWED": {
          "severity": "HEALTH_WARN",
          "summary": {
            "message": "mon is allowing insecure global_id reclaim",
            "count": 1
          },
          "muted": false
        },
        "FS_DEGRADED": {
          "severity": "HEALTH_WARN",
          "summary": {
            "message": "1 filesystem is degraded",
            "count": 1
          },
          "muted": false
        },
        "POOL_NO_REDUNDANCY": {
          "severity": "HEALTH_WARN",
          "summary": {
            "message": "7 pool(s) have no replicas configured",
            "count": 7
          },
          "muted": false
        }
      },
      "mutes": []
    },
    "election_epoch": 19,
    "quorum": [
      0
    ],
    "quorum_names": [
      "ubuntu2004-1"
    ],
    "quorum_age": 3,
    "monmap": {
      "epoch": 1,
      "min_mon_release_name": "pacific",
      "num_mons": 1
    },
    "osdmap": {
      "epoch": 88,
      "num_osds": 1,
      "num_up_osds": 1,
      "osd_up_since": 1686829366,
      "num_in_osds": 1,
      "osd_in_since": 1685098907,
      "num_remapped_pgs": 0
    },
    "pgmap": {
      "pgs_by_state": [
        {
          "state_name": "active+clean",
          "count": 193
        }
      ],
      "num_pgs": 193,
      "num_pools": 7,
      "num_objects": 243,
      "data_bytes": 14901,
      "bytes_used": 23982080,
      "bytes_avail": 32184078336,
      "bytes_total": 32208060416
    },
    "fsmap": {
      "epoch": 70,
      "id": 1,
      "up": 1,
      "in": 1,
      "max": 1,
      "by_rank": [
        {
          "filesystem_id": 1,
          "rank": 0,
          "name": "ubuntu2004-1",
          "status": "up:replay",
          "gid": 84104
        }
      ],
      "up:standby": 0
    },
    "mgrmap": {
      "available": true,
      "num_standbys": 0,
      "modules": [
        "iostat",
        "nfs",
        "restful"
      ],
      "services": {}
    },
    "servicemap": {
      "epoch": 162,
      "modified": "2023-06-15T11:44:11.919071+0000",
      "services": {
        "rgw": {
          "daemons": {
            "summary": "",
            "74111": {
              "start_epoch": 160,
              "start_stamp": "2023-06-15T11:43:09.374625+0000",
              "gid": 74111,
              "addr": "192.168.40.232:0/1318697134",
              "metadata": {
                "arch": "x86_64",
                "ceph_release": "pacific",
                "ceph_version": "ceph version 16.2.13 (5378749ba6be3a0868b51803968ee9cde4833a3e) pacific (stable)",
                "ceph_version_short": "16.2.13",
                "cpu": "Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz",
                "distro": "ubuntu",
                "distro_description": "Ubuntu 20.04.6 LTS",
                "distro_version": "20.04",
                "frontend_config#0": "beast endpoint=0.0.0.0:10100",
                "frontend_type#0": "beast",
                "hostname": "ubuntu2004-1",
                "id": "ubuntu2004-1.rgw0",
                "kernel_description": "#167-Ubuntu SMP Mon May 15 17:35:05 UTC 2023",
                "kernel_version": "5.4.0-150-generic",
                "mem_swap_kb": "0",
                "mem_total_kb": "6052932",
                "num_handles": "1",
                "os": "Linux",
                "pid": "248039",
                "realm_id": "",
                "realm_name": "",
                "zone_id": "82de5b1d-d35c-4500-9589-934496a0dd8b",
                "zone_name": "default",
                "zonegroup_id": "4de41f15-3f7d-4058-a4d9-3032ee6d2de8",
                "zonegroup_name": "default"
              },
              "task_status": {}
            }
          }
        }
      }
    },
    "progress_events": {}
  },
  "exitcode": 0,
  "error": ""
}
```