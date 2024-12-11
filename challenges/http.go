package main
import (
	"fmt"
	"net/http"

)


func main(){
	
	url := "https://linkedin.com"
	ctype,err := contentType(url)

	if err != nil {
		fmt.Println("Error", err)
	}else{
		fmt.Println(ctype)
	}
}

func contentType(url string) (string , error){
	resp,err := http.Get(url)
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-Type")
	if ctype == ""{
		return "", fmt.Errorf("cant find content-Type header")
	}
	
	return ctype, err


}