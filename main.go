package main
import (
	// "fmt"
	// "io/ioutil"
	// "net/http"
	// "regexp"
	"engine"

)
func main(){
	engine.Con
}
// func main(){
// 	resp,err :=	http.Get("https://manhua.dmzj.com/")
// 	if err !=nil{
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Println("error ",resp.StatusCode)
// 		return 
// 	}
// // utf8Reader :=	tranform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
// 	all,err := 	ioutil.ReadAll(resp.Body)
// 	if err !=nil{
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n",all)

	
// }

// func printCityList(contents []byte){
// 	re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
// 	matches:=re.FindAll(contents,-1)
// 	for _, m := range matches {
// 		fmt.Printf("%s\n",m)
// 	}
// 	fmt.Printf("Matches found : %d\n",
// len (matches))

// }