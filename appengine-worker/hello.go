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

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)

  client := urlfetch.Client(c)

//https://golang.org/pkg/net/http/#pkg-index
//https://github.com/domainersuitedev/delicious-api/blob/master/api/posts.md#v1postsget
  req, err := http.NewRequest("GET", "https://api.del.icio.us/v1/posts/recent?red=api&count=10", nil)

  deliciousAuthJsonFile, err := ioutil.ReadFile("./delicious.feed.sucks")
  if err != nil {
    log.Errorf(c, "ioutil.ReadFile: %v", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  type deliciousJsonobject struct {
    Username string `json:"user"`
    Password string `json:"pwd"`
  }
  var deliciousJsonType deliciousJsonobject
  json.Unmarshal(deliciousAuthJsonFile, &deliciousJsonType)
  log.Errorf(c, "deliciousJsonType: %v", deliciousJsonType)

  req.SetBasicAuth(deliciousJsonType.Username, deliciousJsonType.Password)



  respFeed, err := client.Do(req)
  //log.Errorf(c, "respFeed: %v", respFeed)

  if ( respFeed.StatusCode == 200 ) {
    XMLdata, err := ioutil.ReadAll(respFeed.Body)
    if err != nil {
      log.Errorf(c, "ioutil.ReadAll: %v", err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    type Post struct {
            Title       string `xml:"description,attr"`
            Link        string `xml:"href,attr"`
            Description string `xml:"extended,attr"`
    }

type Posts struct {
         Posts []Post `xml:"post"`
 }



    rss := new(Posts)
    bufferRss := bytes.NewBuffer(XMLdata)
    //log.Errorf(c, "bufferRss: %v", bufferRss.String())
    decodedRss := xml.NewDecoder(bufferRss)
    err = decodedRss.Decode(rss)
    log.Errorf(c, "rss: %v", rss)

    tmpl, err := template.ParseFiles("delicious.html")
    if err != nil {
      log.Errorf(c, "tmpl error: %v", err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    var doc bytes.Buffer
    tmpl.Execute(&doc, rss.Posts)
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
}
