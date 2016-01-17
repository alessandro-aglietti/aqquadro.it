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

  jwt, err := ioutil.ReadFile("aqquadro-hrd-a301b0a436c9.json")
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

  gsclient, err := storage.NewClient(c)

  wc := gsclient.Bucket("www.aqquadro.it").Object("delicious.html").NewWriter(ctx)

  //wc := storage.NewWriter(ctx, "www.aqquadro.it", "delicious.html")
  wc.ContentType = "text/html"
  wc.ContentEncoding = "UTF-8"
  wc.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader},{"user-alessandro.aglietti@gmail.com", storage.RoleOwner}}
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
