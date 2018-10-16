package main

import (
  "os"
  "fmt"
  "time"
  "net/http"
  "io/ioutil"
  "crypto/tls"
  "encoding/json"
)

var DIT_DIR = "./.dit";
var CONFIG_FILE = "./.dit/dit.config.json";

func initDit(){
  _, err := os.Stat("./.dit");

  if os.IsNotExist(err){

    os.MkdirAll(DIT_DIR, os.ModePerm);
    var file, err = os.Create(CONFIG_FILE);

    if err != nil {
      fmt.Println("Error Occured");
      return;
    }
    defer file.Close();


    var creationTime = time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006");

    config := map[string] string {
      "CreationTime": creationTime,
    }
    configStr, _ := json.Marshal(config);
    err = ioutil.WriteFile(CONFIG_FILE, []byte(configStr), os.ModePerm);

    fmt.Println("Initialized Dit Directory!");
  } else {
    fmt.Println("Already a Dit Directory.");
  }
}

func setURL(url string){
  buffer, err := ioutil.ReadFile(".dit/dit.config.json");
  if err != nil {
    fmt.Println("Error");
  }

  var config map[string] interface {};
  json.Unmarshal(buffer, &config);

  config["Url"] = url;
  configStr, _ := json.Marshal(config);

  err = ioutil.WriteFile(CONFIG_FILE, []byte(configStr), os.ModePerm);
  if err != nil {
    fmt.Println("Error Occuredd");
  } else {
    fmt.Println("Configuration Successfully Updated!");
  }
}

func sync(url string){
  resp, err := http.Get("http://intranet.daiict.ac.in/~daiict_nt01/");

  if err != nil {

  }

  defer resp.Body.Close();
  body, err := ioutil.ReadAll(resp.Body);
  fmt.Println(body);
}

func main(){
  // Skip Certificate Verification
  http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true};

  args := os.Args[1:];

  if len(args) == 0 {
    fmt.Println("HELP");
  } else if args[0] == "init" {
    initDit();
  } else if args[0] == "set-url" {
    if(len(args) >= 2){
      url := args[1];
      setURL(url);
    } else {
      fmt.Println("Invalid Arguments: Give URL to set");
    }
  }
}
