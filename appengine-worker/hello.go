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
    "encoding/xml"
)

type Rss struct {
         Channel Channel `xml:"channel"`
 }

 type Item struct {
         Title       string `xml:"title"`
         Link        string `xml:"link"`
         Description string `xml:"description"`
 }

 type Channel struct {
         Title       string `xml:"title"`
         Link        string `xml:"link"`
         Description string `xml:"description"`
         Items       []Item `xml:"item"`
 }

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)

  client := urlfetch.Client(c)

  tmpl, err := template.ParseFiles("delicious.html")
  if err != nil {
    log.Errorf(c, "tmpl error: %v", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  req, err := http.NewRequest("GET", "https://api.delicious.com/v1/posts/recent", nil)
  req.Header.Add("Authorization", "")

  respFeed, err := client.Get("http://feeds.delicious.com/v2/rss/aqquadro")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  XMLdata, err := ioutil.ReadAll(respFeed.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  rss := new(Rss)
  bufferRss := bytes.NewBuffer(XMLdata)
  decodedRss := xml.NewDecoder(bufferRss)
  err = decodedRss.Decode(rss)

  var doc bytes.Buffer
  tmpl.Execute(&doc, rss.Channel.Items)
  var docString = doc.String();
  //log.Errorf(c, "docString: %v", docString)

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
  if err != nil {
    log.Errorf(c, "storage.NewClient error: %v", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

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
