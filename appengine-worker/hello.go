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
    "compress/gzip"
    "time"
    "strings"
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
         Ts          string
 }
    rss := new(Posts)
    bufferRss := bytes.NewBuffer(XMLdata)
    //log.Errorf(c, "bufferRss: %v", bufferRss.String())
    decodedRss := xml.NewDecoder(bufferRss)
    err = decodedRss.Decode(rss)
    log.Errorf(c, "rss: %v", rss)

    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
      log.Errorf(c, "tmpl error: %v", err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    // non usare l'array sotto rss.Posts ma passare rss

    t := time.Now()
  	loc, _ := time.LoadLocation("Europe/Rome")
    t.In(loc)
    _, summerOffset := t.Zone()
    t = t.Add(time.Duration(summerOffset)*time.Second)
    rss.Ts = t.Format(time.RFC3339)


    var doc bytes.Buffer
    tmpl.Execute(&doc, rss)
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

    wc := gsclient.Bucket("www.aqquadro.it").Object("index.html").NewWriter(ctx)
    // old writer wc := storage.NewWriter(ctx, "www.aqquadro.it", "delicious.html")
    wc.ContentType = "text/html"
    //wc.ContentEncoding = "UTF-8"
    wc.ContentEncoding = "gzip"
    wc.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader},{"user-alessandro.aglietti@gmail.com", storage.RoleOwner}}

    // gzipppp
    var gzippedbytes bytes.Buffer
    gz := gzip.NewWriter(&gzippedbytes)
    if _, err := gz.Write([]byte(docString)); err != nil {
        log.Errorf(c, "gz.Write error: %v", err)
    }
    if err := gz.Flush(); err != nil {
        log.Errorf(c, "gz.Flush error: %v", err)
    }
    if err := gz.Close(); err != nil {
        log.Errorf(c, "wc.Close error: %v", err)
    }

    if _, err := wc.Write(gzippedbytes.Bytes()); err != nil {
        log.Errorf(c, "wc.Write error: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := wc.Close(); err != nil {
        log.Errorf(c, "wc.Close error: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // update sitemappa
    // YYYY-MM-DDThh:mm:ssTZD

    sitemaptmpl, err := template.ParseFiles("sitemap.xml")
    if err != nil {
      log.Errorf(c, "tmpl error: %v", err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    var sitemaBuffer bytes.Buffer
    sitemaptmpl.Execute(&sitemaBuffer, rss.Ts)
    var siteMapString = sitemaBuffer.String();
    siteMapString = strings.Replace(siteMapString, "&lt;", "<", -1)

    log.Errorf(c, "sitemap: %v", siteMapString)



    wc = gsclient.Bucket("www.aqquadro.it").Object("sitemap.xml").NewWriter(ctx)
    // old writer wc := storage.NewWriter(ctx, "www.aqquadro.it", "delicious.html")
    wc.ContentType = "text/xml"
    wc.ContentEncoding = "UTF-8"
    wc.ACL = []storage.ACLRule{{storage.AllUsers, storage.RoleReader},{"user-alessandro.aglietti@gmail.com", storage.RoleOwner}}

    byteArray := []byte(siteMapString)
    _, err = wc.Write(byteArray)
    if err != nil {
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
