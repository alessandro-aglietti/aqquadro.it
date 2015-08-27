package hello

import (
    "encoding/json"
    "html/template"
    "net/http"
    "google.golang.org/appengine"
    "google.golang.org/appengine/log"
    "google.golang.org/appengine/urlfetch"
    "bytes"
    "golang.org/x/oauth2/google"
    "io/ioutil"
    "google.golang.org/cloud"
    "google.golang.org/cloud/storage"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)

  client := urlfetch.Client(c)

  resp, err := client.Get("http://feeds.delicious.com/v2/json/aqquadro?count=10")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  deliciousBytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  type DeliciousDataModel struct {
  	Title string `json:"d"`
  	Comment string `json:"n"`
  	Url string `json:"u"`
  }

  var deliciousData []DeliciousDataModel

  err = json.Unmarshal(deliciousBytes, &deliciousData)
  if err != nil {
    log.Errorf(c, "json error: %v", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  tmpl, err := template.ParseFiles("delicious.html")
  if err != nil {
    log.Errorf(c, "tmpl error: %v", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  var doc bytes.Buffer
  tmpl.Execute(&doc, deliciousData)
  var docString = doc.String();

  jwt, err := ioutil.ReadFile("aqquadro-hrd-e9c7a4d18539.json")
  if err != nil {
    log.Errorf(c, "jwt error: %v", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  conf, err := google.JWTConfigFromJSON(jwt, storage.ScopeFullControl)
  if err != nil {
    log.Errorf(c, "conf error: %v", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  storageclient := conf.Client(c)

  ctx := cloud.NewContext(appengine.AppID(c), storageclient)

  wc := storage.NewWriter(ctx, "www.aqquadro.it", "delicious.html")
  wc.ContentType = "text/html"
  wc.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader}}
  if _, err := wc.Write([]byte(docString)); err != nil {
      log.Errorf(c, "wc.Write error: %v", err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
  if err := wc.Close(); err != nil {
      log.Errorf(c, "wc.Close error: %v", err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
}
