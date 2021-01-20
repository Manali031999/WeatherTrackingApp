package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"html/template"
)

type AutoGenerated struct {
	Valid bool
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}


func disp(w http.ResponseWriter, r *http.Request) {
	if r.Method=="GET"{
		var city string
	key:="http://api.weatherapi.com/v1/current.json?key=%2087e35af73e7f4383976105618211901&q="
	city=r.FormValue("city")
	url:=key+city
	resp,err:= http.Get(url)
	if err!=nil{
		fmt.Println(err)
	}
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println(err)
	}
	var auto AutoGenerated
	err=json.Unmarshal(body,&auto)
	if err!=nil{
		fmt.Println("Error in unMarshaling",err)
	}
	file, _ := json.MarshalIndent(auto, "", " ")
	_ = ioutil.WriteFile("cities.json", file, 0644)
	usersfiledata, err := ioutil.ReadFile("cities.json")
	if err != nil {
		fmt.Println(err)
	}
	var alluserdetails AutoGenerated
	err = json.Unmarshal([]byte(usersfiledata), &alluserdetails)
	if err != nil {
		fmt.Println("Error JSON Unmarshling for user file")
		fmt.Println(err)
	}
	if alluserdetails.Location.Lat==0{
		alluserdetails.Valid=false
	}else{
		alluserdetails.Valid=true
	}
	t, _ := template.ParseFiles("disp.html")
	t.Execute(w, alluserdetails)
	}	
}

func open(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("input.html")
    t.Execute(w, nil)
}

func main(){
	http.HandleFunc("/",open)
	http.HandleFunc("/disp.html",disp)
    http.ListenAndServe(":8081", nil)
}
