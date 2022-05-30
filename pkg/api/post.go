package api

type PostInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

//func CreatePost(ctx *gin.Context) {
//	var input PostInput
//	var post *mydb.Post
//	mydb.CreatePost(post)
//	//mydb.Database.Db.Model(&post).Association("Users").Append(&user)
//}
//
//func UpdatePost(ctx *gin.Context) {
//
//	category, okCat := getCategoryById(id, w)
//	if !okCat {
//		return
//	}
//
//	err := json.NewDecoder(r.Body).Decode(&category)
//	if err != nil {
//		return
//	}
//
//	mydb.Database.Db.Save(&category)
//
//}
//
//func DeletePost(ctx *gin.Context) {
//
//	category, okCat := getCategoryById(id, w)
//	if !okCat {
//		return
//	}
//
//	mydb.Database.Db.Delete(&category)
//	w.WriteHeader(http.StatusOK)
//}
