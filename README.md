# akka-cluster-manager
A wrapper around the akka-cluster Command Line Management: http://doc.akka.io/docs/akka/2.4.1/java/cluster-usage.html#cluster-command-line-java

Following is the usage of akka-cluster Command Line Management usage:
```
Usage: bin/akka-cluster <node-hostname> <jmx-port> <command> ...
Supported commands are:
           join <node-url> - Sends request a JOIN node with the specified URL
          leave <node-url> - Sends a request for node with URL to LEAVE the cluster
           down <node-url> - Sends a request for marking node with URL as DOWN
             member-status - Asks the member node for its current status
                   members - Asks the cluster for addresses of current members
               unreachable - Asks the cluster for addresses of unreachable members
            cluster-status - Asks the cluster for its current status (member ring,
                             unavailable nodes, meta data etc.)
                    leader - Asks the cluster who the current leader is
              is-singleton - Checks if the cluster is a singleton cluster (single
                             node cluster)
              is-available - Checks if the member node is available
Where the <node-url> should be on the format of
  'akka.<protocol>://<actor-system-name>@<hostname>:<port>'

Examples: bin/akka-cluster localhost 9999 is-available
          bin/akka-cluster localhost 9999 join akka.tcp://MySystem@darkstar:2552
          bin/akka-cluster localhost 9999 cluster-status
```

You don't need remember the JMX port or the exact node name to send a command to it.
Usage: akka-cluster <command> --env <environment-name>

Examples: akka-cluster leave --env stage --node darkstar
          akka-cluster join  --env prod  --node darkstarg
