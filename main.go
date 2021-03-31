package main

import(
	"fmt"
	"os"
	"flag"
	"./torrentfile"
	"./client"
)

func main () {
	torrent_file := flag.String("i","","Input torrent file")
	//output_dir := flag.String("o","","Download file path")
	verbose := flag.Bool("v",false,"Verbose torrent download")
	flag.Parse()
	if(*torrent_file != ""){
		f_exists,f_err := exists(*torrent_file);
		if(!f_exists) {
			fmt.Println("Input file does not exist")
			os.Exit(1)
		}
		if(f_err != nil) {
			fmt.Println(f_err)
			os.Exit(1)
		}
	}
	torrent_info,torrent,err := torrentfile.OpenFile(*torrent_file)
	if(*verbose){
		fmt.Println(torrent_info)
	}
	client,err := client.Init(torrent)
	if(err == nil){
		fmt.Println("Error in starting torrent client.")
	}	
	
	client.BeginDownload(output_dir,*verbose)
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
