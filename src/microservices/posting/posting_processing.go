package posting

import (
	"fmt"
	"github.com/admin-agora/backend/messages"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

type CommentPostProcessing struct {
	db *gorm.DB
}

func (p *CommentPostProcessing) Init() error {
	dosn := "agora_mysql_admin:DigitalMcDonalds$3.21@tcp(agora-mysql-dev.mysql.database.azure.com)/agora?charset=utf8mb4&parseTime=true"
	fmt.Println(dosn)
	gormDb, sqlErr := gorm.Open(mysql.Open(dosn), &gorm.Config{})

	// Assign the SQL connection to
	p.db = gormDb

	if sqlErr != nil {
		return sqlErr
	}

	return nil
}

func (p *CommentPostProcessing) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal("error in byte reading")
	}
	log.Println("parsed the io stream")
	commentPosted := &messages.CommentPosted{}
	unmarshallErr := protojson.Unmarshal(body, commentPosted)
	if unmarshallErr != nil {
		log.Fatal("error in unmarshall", unmarshallErr)
	}
	// json decode to protobuf
	// persist comment
	// breakdown
	_, fmtError := fmt.Fprintf(w, "Hello, %s!", html.EscapeString(commentPosted.Comment))
	if fmtError != nil {
		log.Fatal("error in f print f")
	}
}



