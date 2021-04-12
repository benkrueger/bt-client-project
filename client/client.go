package main

import(
	"fmt"
	"os"
	"flag"
	"../common"
	"net/http"
	"io/ioutil"
	"errors"
)

func main () {
	metadata_file := flag.String("i","","Input metadata file")

	flag.Parse()
	if(*metadata_file != ""){
		f_exists,f_err := exists(*metadata_file);
		if(!f_exists) {
			fmt.Println("Input file does not exist")
			os.Exit(1)
		}
		if(f_err != nil) {
			fmt.Println(f_err)
			os.Exit(1)
		}
	}
	md,err := common.ReadFromMetadataFile(*metadata_file)
	if(err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
	file_infohash,infohash_err := md.GetInfoHash()
	if infohash_err != nil {
		fmt.Println("Infohash error :",err)
		os.Exit(1)
	}
	peerlist_request_err  := RequestPeerList(md.Announce,file_infohash)
	if peerlist_request_err != nil {
		fmt.Println("Peerlist error:",peerlist_request_err)
		os.Exit(1)
	}

}

func RequestPeerList(announce string, infohash string)(error) {
	urlstr := "http://localhost:3000/annouce?info_hash=82d19f99061f6e3f107aa5e7a338762c940436f2" //fmt.Sprintf("http://%s:3000/?info_hash=%s",announce,infohash)
	fmt.Println(urlstr)
	res,err := http.Get(urlstr)
	if err != nil {
		return err
	}
	var peerlistbytes [] byte
	peerlistbytes,err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	err = common.PeerMapUnmarshalJSON(peerlistbytes,infohash)
	if err != nil {
		return err
	}
	if !common.PeersPresentForInfohash(infohash) {
		return errors.New("No peers present for infohash. Maybe try again later.")
	}
	return nil
} 
// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
