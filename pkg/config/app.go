package config

// motive is to write db
import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db * gorm.BD
)

func Connect(){
	d, err := gorm.Open("mysql","hackeramitkumar:Cows@9116260192@/simplerest")
}