package main

import(
	"fmt"
	"flag"
	"net/http"
	"../common"
	"log"
	"os"

)
var t *common.Metadata
//var peers map[string]*common.Peer

func main() {
	file := flag.String("f","","File to track")
	port := flag.String("p","8080","TCP port to listen to peer requests")
	export := flag.Bool("e",false,"Print metadata file to stdout")
	flag.Parse()
	var err error
	
	t,err = common.CreateMetadata(2048,"localhost",*file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var infohash string
	infohash,err = t.GetInfoHash()
	common.InitializePeerMap(infohash,2556)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if *export {
		err = OutputTrackerFile()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}else{
			os.Exit(0)
		}
	}
	fmt.Printf("Announce url: http://localhost:%s/annouce?info_hash=%s\n",*port,infohash)
	fmt.Printf("Infohash %s\n",infohash)
	http.HandleFunc("/annouce",HandleTorrentRequest)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",*port),nil))

}
func OutputTrackerFile() (error){
	jsbytes,err := t.OutputJSON()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n",string(jsbytes))
	return nil
}



func HandleTorrentRequest(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	fmt.Println(r.Method)
	if r.Method == "GET" && r.URL.Query()["info_hash"] != nil{
		jsbytes,_ := common.PeerMapMarshalJSON()
		w.Write(jsbytes)
	}else{
		w.Write([]byte("{\"error\":\"invalid request\"}"))
	}
	
}