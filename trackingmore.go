package main
import(
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	"strings"
	//"encoding/json"  
 )
																																																										 
func main(){
    //1 List all carriers
    //var url string ="http://api.trackingmore.com/v2/carriers/"
    //var postData string =""
    //httpDo(url,postData,"GET")

    //2 detect a carriers by tracking number
    //var url string ="http://api.trackingmore.com/v2/carriers/detect"
    //var postData string ="{\"tracking_number\":\"RP325552475CN\"}"
    //httpDo(url,postData,"POST")

    //3 create a tracking number
    var url string ="http://api.trackingmore.com/v2/trackings/post"
    var postData string ="{\"tracking_number\": \"RP325552475CN\",\"carrier_code\":\"xru\",\"title\":\"4PX page\",\"customer_name\":\"trackingmore user\",\"customer_email\":\"service@trackingmore.com\",\"order_id\":\"#123\",\"order_create_time\":\"2018/09/08 16:51\",\"destination_code\":\"US\",\"tracking_ship_date\":\"20180908\",\"tracking_postal_code\":\"12201\",\"lang\":\"en\",\"logistics_channel\":\"API TEST\"}"
    httpDo(url,postData,"POST")

    //4 List all trackings
    //var url string ="http://api.trackingmore.com/v2/trackings/get?page=1&limit=100&created_at_min=&created_at_max=1541314361&update_time_min=&update_time_max=1541314361&order_created_time_min=&order_created_time_max=&numbers=RP325552475CN&orders=&lang=en"
    //var postData string =""
    //httpDo(url,postData,"GET")

    //5 Get tracking results of a single tracking
    //var url string ="http://api.trackingmore.com/v2/trackings/xru/RP325552475CN/en"
    //var postData string =""
    //httpDo(url,postData,"GET")

    //6 create muti tracking number
    //var url string ="http://api.trackingmore.com/v2/trackings/batch"
    //var postData string = "[{\"tracking_number\": \"RP325552475CN\",\"carrier_code\":\"xru\",\"title\":\"4PX page\",\"customer_name\":\"trackingmore user\",\"customer_email\":\"service@trackingmore.com\",\"order_id\":\"#123\",\"order_create_time\":\"2018/09/08 16:51\",\"destination_code\":\"US\",\"tracking_ship_date\":\"20180908\",\"tracking_postal_code\":\"12201\",\"lang\":\"en\",\"logistics_channel\":\"API TEST\"},{\"tracking_number\": \"LZ448865302CN\",\"carrier_code\":\"china-ems\",\"title\":\"4PX page\",\"customer_name\":\"trackingmore user\",\"customer_email\":\"service@trackingmore.com\",\"order_id\":\"#123\",\"order_create_time\":\"2018/09/08 16:51\",\"destination_code\":\"US\",\"tracking_ship_date\":\"20180908\",\"tracking_postal_code\":\"12201\",\"lang\":\"en\",\"logistics_channel\":\"API TEST\"}]"
    //httpDo(url,postData,"POST")

    //7 Update Tracking item
    //var url string ="http://api.trackingmore.com/v2/trackings/xru/RP325552475CN"
    //var postData string = "{\"title\": \"4PX page\",\"customer_name\":\"trackingmore user\",\"customer_email\":\"service@trackingmore.com\",\"order_id\":\"#123\",\"logistics_channel\":\"API TEST\",\"customer_phone\":\"+86 13142052920\",\"destination_code\":\"US\",\"status\":\"7\"}"
    //httpDo(url,postData,"PUT")

    //8 Delete a tracking item
    //var url string ="http://api.trackingmore.com/v2/trackings/xru/RP325552475CN"
    //var postData string = ""
    //httpDo(url,postData,"DELETE")

    //9 Get realtime tracking results of a single tracking
    //var url string ="http://api.trackingmore.com/v2/trackings/realtime"
    //var postData string = "{\"tracking_number\": \"RP325552475CN\",\"carrier_code\":\"xru\",\"destination_code\":\"US\",\"tracking_ship_date\": \"20180908\",\"tracking_postal_code\":\"12201\",\"specialNumberDestination\":\"UK\",\"order\":\"#123\",\"order_create_time\":\"2018/09/08 16:51\",\"lang\":\"en\"}"
    //httpDo(url,postData,"POST")

    //10 Modify courier code
    //var url string ="http://api.trackingmore.com/v2/trackings/update"
    //var postData string = "{\"tracking_number\": \"RP325552475CN\",\"carrier_code\":\"xru\",\"update_carrier_code\":\"china-ems\"}"
    //httpDo(url,postData,"POST")

    //11 Get user info
    //var url string ="http://api.trackingmore.com/v2/trackings/getuserinfo"
    //var postData string = ""
    //httpDo(url,postData,"GET")

    //12 Get status number
    //var url string ="http://api.trackingmore.com/v2/trackings/getstatusnumber"
    //var postData string = ""
    //httpDo(url,postData,"GET")

    //13 Set number not update
    //var url string ="http://api.trackingmore.com/v2/trackings/notupdate"
    //var postData string = "[{\"tracking_number\":\"RP325552475CN\",\"carrier_code\":\"xru\"},{\"tracking_number\":\"LZ448865302CN\",\"carrier_code\":\"china-ems\"}]"
    //httpDo(url,postData,"POST")

    //14 Get remote iterm results
    //var url string = "http://api.trackingmore.com/v2/trackings/remote"
    //var postData string = "[{\"country\":\"CN\",\"postcode\":\"400422\"},{\"country\":\"CN\",\"postcode\":\"412000\"}]"
    //httpDo(url,postData,"POST")

    //15 Get cost time iterm results
    //var url string = "http://api.trackingmore.com/v2/trackings/costtime"
    //var postData string = "[{\"carrier_code\":\"xru\",\"destination\":\"US\",\"original\":\"CN\"},{\"carrier_code\":\"china-ems\",\"destination\":\"US\",\"original\":\"CN\"}]"
    //httpDo(url,postData,"POST")

    //16 Delete multiple tracking item
    //var url string = "http://api.trackingmore.com/v2/trackings/delete"
    //var postData string = "[{\"tracking_number\":\"RP325552475CN\",\"carrier_code\":\"xru\"},{\"tracking_number\":\"LZ448865302CN\",\"carrier_code\":\"china-ems\"}]"
    //httpDo(url,postData,"POST")

    //17 Update multiple Tracking item
    //var url string = "http://api.trackingmore.com/v2/trackings/updatemore"
    //var postData string = "[{\"tracking_number\":\"RP325552475CN\",\"carrier_code\":\"xru\",\"title\": \"4PX page\",\"customer_name\":\"trackingmore user\",\"customer_email\":\"service@trackingmore.com\",\"order_id\":\"#123\",\"logistics_channel\":\"API TEST\",\"destination_code\":\"US\",\"status\":\"7\"},{\"tracking_number1\":\"RP325552475CN\",\"carrier_code\":\"xru\",\"title\": \"4PX page\",\"customer_name\":\"trackingmore user\",\"customer_email\":\"service@trackingmore.com\",\"order_id\":\"#123\",\"logistics_channel\":\"API TEST\",\"destination_code\":\"US\",\"status\":\"7\"}]"
    //httpDo(url,postData,"POST")

}


func httpDo(url,postData,method string) {
    client := &http.Client{}
    req, err := http.NewRequest(method,url,strings.NewReader(postData))
    if err != nil {
        // handle error
    }
    req.Header.Set("Content-Type", "application/json'")
    req.Header.Set("Trackingmore-Api-Key", "YOUR API KEY")
    resp, err := client.Do(req)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
    fmt.Println(string(body))
}

 
