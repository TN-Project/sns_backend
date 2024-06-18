package create_group

import(
    "log/slog"
    "github.com/gin-gonic/gin"
    "sns_backend/pkg/db"
)
func ExistUserCheck(username string) (bool, error) {
    db := db.Connect()
    defer db.Close()

    sql := `select count(*) from "user" where username = $1`
    var count int
    err := db.QueryRow(sql, username).Scan(&count)
    if err != nil {
        slog.Error("Error checking user existence: " + err.Error())
        return false, err
    }
    if count == 0 {
        return false, nil
    }   
    return true, nil
}
func ExistUserPost(c *gin.Context){

    c.JSON(200,gin.H{

    })

}