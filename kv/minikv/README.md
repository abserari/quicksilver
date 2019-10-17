## minikeyvalue


厌倦了分布式文件系统的复杂性？

minikeyvalue是一个约1000行的分布式键值存储，支持复制，多台计算机以及每台计算机多个驱动器。针对1MB到1GB之间的值进行了优化。受SeaweedFS启发，但很简单。应该扩展到数十亿个文件和PB级数据。用于comma.ai的生产。

minikeyvalue的简单性的一个关键部分是使用stock nginx作为卷服务器。

即使这段代码很烂，磁盘上的格式也非常简单！我们依靠文件系统进行Blob存储，并依靠LevelDB进行索引。可以通过重建来重建索引。可以通过重新平衡来添加或删除卷。

API
- GET /key
    - 302 redirect to nginx volume server.
- PUT /key
    - Blocks. 201 = written, anything else = probably not written.
- DELETE /key
    - Blocks. 204 = deleted, anything else = probably not deleted.

Start Master Server (default port 3000)
```
        ./mkv -volumes localhost:3001,localhost:3002 -db /tmp/indexdb/ server
```
Start Volume Servers (default port 3001)
```
# this is just nginx under the hood
PORT=3001 ./volume /tmp/volume1/
PORT=3002 ./volume /tmp/volume2/
```
## Usage
```
# put "bigswag" in key "wehave"
curl -v -L -X PUT -d bigswag localhost:3000/wehave

# get key "wehave" (should be "bigswag")
curl -v -L localhost:3000/wehave

# delete key "wehave"
curl -v -L -X DELETE localhost:3000/wehave

# unlink key "wehave", this is a virtual delete
curl -v -L -X UNLINK localhost:3000/wehave

# list keys starting with "we"
curl -v -L localhost:3000/we?list

# list unlinked keys ripe for DELETE
curl -v -L localhost:3000/?unlinked

# put file in key "file.txt"
curl -v -L -X PUT -T /path/to/local/file.txt localhost:3000/file.txt

# get file in key "file.txt"
curl -v -L -o /path/to/local/file.txt localhost:3000/file.txt
```
## ./mkv Usage
```
Usage: ./mkv <server, rebuild, rebalance>

  -db string
        Path to leveldb
  -fallback string
        Fallback server for missing keys
  -port int
        Port for the server to listen on (default 3000)
  -protect
        Force UNLINK before DELETE
  -replicas int
        Amount of replicas to make of the data (default 3)
  -subvolumes int
        Amount of subvolumes, disks per machine (default 10)
  -volumes string
        Volumes to use for storage, comma separated
```
## Rebalancing (to change the amount of volume servers)
```
# must shut down master first, since LevelDB can only be accessed by one process
./mkv -volumes localhost:3001,localhost:3002 -db /tmp/indexdb/ rebalance
```
## Rebuilding (to regenerate the LevelDB)
```
./mkv -volumes localhost:3001,localhost:3002 -db /tmp/indexdbalt/ rebuild
```
## Performance
```
# Fetching non-existent key: 116338 req/sec
wrk -t2 -c100 -d10s http://localhost:3000/key

# go run thrasher.go lib.go
starting thrasher
10000 write/read/delete in 2.620922675s
thats 3815.40/sec
```