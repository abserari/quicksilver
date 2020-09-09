package mr

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"sort"
	"strconv"
	"time"
)

//
// Map functions return a slice of KeyValue.
//
type KeyValue struct {
	Key   string
	Value string
}

// for sorting by key.
type ByKey []KeyValue

// for sorting by key.
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

//
// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
//
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

func genWorkerID() (uuid string) {
	// generate 32 bits timestamp
	unix32bits := uint32(time.Now().UTC().Unix())

	buff := make([]byte, 12)

	numRead, err := rand.Read(buff)

	if numRead != len(buff) || err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x-%x\n", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])
}

//
// main/mrworker.go calls this function.
//
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {

	// Your worker implementation here.

	// uncomment to send the Example RPC to the master.
	// CallExample()
	workerId := genWorkerID()
	retry := 3
	// for to get tasks
	for {
		args := WorkArgs{Workerid: workerId}
		reply := WorkReply{}
		working := call("Master.Work", &args, &reply)
		// log.Println(working, reply.Isfinished)
		if reply.Isfinished || !working {
			log.Println("finished")
			return
		}
		log.Println("task info:", reply)
		//working switch map or reduce
		switch reply.MapReduce {
		case "map":
			MapWork(reply, mapf)
			retry = 3
		case "reduce":
			ReduceWork(reply, reducef)
			retry = 3
		default:
			log.Println("error reply: would retry times:", retry)
			if retry < 0 {
				return
			}
			retry--
		}

		commitArgs := CommitArgs{Workerid: workerId, Taskid: reply.Taskid, MapReduce: reply.MapReduce}
		commitReply := CommitReply{}
		_ = call("Master.Commit", &commitArgs, &commitReply)

		time.Sleep(500 * time.Millisecond)
	}
}

// get tasks and write to mr-tmp-taskid-reduceid
func MapWork(task WorkReply, mapf func(string, string) []KeyValue) {
	// check task info
	file, err := os.Open(task.Filename)
	if err != nil {
		log.Fatalf("cannot open %v", task.Filename)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", task.Filename)
	}
	kva := mapf(task.Filename, string(content))

	sort.Sort(ByKey(kva))

	// create file buckets
	tmpName := "mr-tmp-" + strconv.Itoa(task.Taskid)
	var fileBucket = make(map[int]*json.Encoder)
	for i := 0; i < task.BucketNumber; i++ {
		ofile, _ := os.Create(tmpName + "-" + strconv.Itoa(i))
		fileBucket[i] = json.NewEncoder(ofile)
		defer ofile.Close()
	}

	for _, kv := range kva {
		key := kv.Key
		reduce_idx := ihash(key) % task.BucketNumber
		err := fileBucket[reduce_idx].Encode(&kv)
		if err != nil {
			log.Fatal("Unable to write to file")
		}
	}
}

// get reduce task and reduce all reduce id = task.Taskid files.
func ReduceWork(task WorkReply, reducef func(string, []string) string) {
	//check task info
	intermediate := []KeyValue{}

	// read mr-tmp n files to add inermediate then write
	for mapTaskNumber := 0; mapTaskNumber < task.BucketNumber; mapTaskNumber++ {
		filename := "mr-tmp-" + strconv.Itoa(mapTaskNumber) + "-" + strconv.Itoa(task.Taskid)
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal("Unable to read from: ", filename)
		}
		defer f.Close()

		decoder := json.NewDecoder(f)
		var kv KeyValue
		for decoder.More() {
			err := decoder.Decode(&kv)
			if err != nil {
				log.Fatal("Json decode failed, ", err)
			}
			intermediate = append(intermediate, kv)
		}
	}

	sort.Sort(ByKey(intermediate))

	// write to "mr-out-Y" Y is reduce task id
	i := 0
	ofile, err := os.Create("mr-out-" + strconv.Itoa(task.Taskid+1))
	if err != nil {
		log.Fatal("Unable to create file: ", ofile)
	}
	defer ofile.Close()

	log.Println("complete to ", task.Taskid, "start to write in to ", ofile)

	for i < len(intermediate) {
		j := i + 1
		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}
		values := []string{}
		for k := i; k < j; k++ {
			values = append(values, intermediate[k].Value)
		}
		output := reducef(intermediate[i].Key, values)

		// this is the correct format for each line of Reduce output.
		fmt.Fprintf(ofile, "%v %v\n", intermediate[i].Key, output)

		i = j
	}
	ofile.Close()

}

//
// example function to show how to make an RPC call to the master.
//
// the RPC argument and reply types are defined in rpc.go.
//
func CallExample() {

	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the reply.
	call("Master.Example", &args, &reply)

	// reply.Y should be 100.
	fmt.Printf("reply.Y %v\n", reply.Y)
}

//
// send an RPC request to the master, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := masterSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
