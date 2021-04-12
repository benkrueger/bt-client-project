package common
import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	//"strings"
	"io"
	"io/ioutil"
	"os"
)
type InfoField struct {
	Name string `json:"name"`
	Length int64 `json:"length"`
	Piece_length int64 `json:"piece_length"`
	Pieces []string `json:pieces`
}
type Metadata struct {
	Announce string `json:announce`
	Info *InfoField `json:info`
}


//this function creates an info field struct with piece length plen from path p
func CreateInfoField(plen int64,path string)(*InfoField,error){
	i := InfoField{}

	i.Piece_length = plen
	f,err := os.Open(path)
	fileinfo,err := f.Stat()
	if err != nil {
		return nil,err
	}
	i.Name = fileinfo.Name()
	i.Length = fileinfo.Size()
	num_pieces := int(i.Length/i.Piece_length)
	
	for n := 0;n < num_pieces;n++ {
		p_buffer := make([]byte, i.Piece_length)
		io.ReadAtLeast(f,p_buffer,int(i.Piece_length))
		i.Pieces = append(i.Pieces,fmt.Sprintf("%x",sha1.Sum(p_buffer)))

	}
	f.Close()
	return &i,err
}
//This function returns Metadata file object for the file at path. 
func CreateMetadata(plen int64,url string, path string) (*Metadata,error){
	var err error
	t := Metadata{}
	t.Announce = url
	t.Info,err  = CreateInfoField(plen,path)
	if err != nil {
		return nil,err
	}
	return &t,err
}
func ReadFromMetadataFile(path string)(*Metadata,error) {
	file, err := ioutil.ReadFile(path)
	var data Metadata
	if err != nil {
		return nil,err
	}
	err = json.Unmarshal([]byte(file),&data)
	return &data,err
}
func (t *Metadata)OutputJSON()([]byte,error) {
	b, err := json.Marshal(t)
	return b,err
}

func (t *Metadata)GetInfoHash()(string,error){
	infobytes,err := json.Marshal(t.Info)
	if err != nil {
		return "",err
	}
	return fmt.Sprintf("%x",sha1.Sum(infobytes)),err
}
func (t *Metadata)PrintMetadata() {
	fmt.Printf("Announce %s\n",t.Announce)
	fmt.Printf("Name %s\n",t.Info.Name)
	fmt.Printf("Length %d\n",t.Info.Length)
	fmt.Printf("Piece_length %d\n",t.Info.Piece_length)
	fmt.Printf("Pieces:\n")
	for i,v :=range t.Info.Pieces {
		fmt.Printf("%d:%s \n\t",i,v)
	}
}