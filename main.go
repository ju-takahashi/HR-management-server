package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 社員情報
type member struct {
	Id       string `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	JoinDate string `form:"joinDate" json:"joinDate"`
}

func main() {
	router := gin.Default()

	members := []member{
		{Id: "1", Name: "テスト太郎", JoinDate: "2021/08/19"},
		{Id: "2", Name: "テスト花子", JoinDate: "2021/08/20"},
		{Id: "3", Name: "山田太郎", JoinDate: "2021/08/21"},
		{Id: "4", Name: "北海道道三郎", JoinDate: "2021/08/22"},
		{Id: "5", Name: "岩手いわ", JoinDate: "2021/08/23"},
		{Id: "6", Name: "沖縄縄子", JoinDate: "2021/08/24"},
		{Id: "7", Name: "新潟潟子", JoinDate: "2021/08/25"},
		{Id: "8", Name: "千葉葉子", JoinDate: "2021/08/26"},
		{Id: "9", Name: "金沢沢子", JoinDate: "2021/08/27"},
		{Id: "10", Name: "テスト二郎", JoinDate: "2021/08/28"},
	}

	id := 11

	// 社員一覧データの取得
	router.GET("/members", func(ctx *gin.Context) {
		ctx.JSON(200, members)
	})

	// 社員一覧データの取得(社員検索)
	router.GET("/members/", func(ctx *gin.Context) {
		// 検索パラメータとして指定された値を取得する。
		// URL例：/members/?name=テスト → "テスト" を取得
		searchName := ctx.Request.URL.Query().Get("name")

		var searchedMembers []member
		for i := 0; i < len(members); i++ {
			// 文字列の部分一致を調べるために、stringsパッケージを使用する。
			// stringsパッケージのContainsメソッドを使用する
			// (特定の文字列が検索対象の文字列に含まれる場合、Trueを返却する。含まれない場合はFalse)
			if strings.Contains(members[i].Name, searchName) {
				searchedMembers = append(searchedMembers, members[i])
			}
		}
		ctx.JSON(200, searchedMembers)
	})

	// 社員詳細データの取得
	router.GET("/members/:id", func(ctx *gin.Context) {
		// パラメータの社員IDを取得する
		selectedId := ctx.Param("id")

		var selectedMember member
		for i := 0; i < len(members); i++ {
			// IDの一致する社員情報を取得する。
			if members[i].Id == selectedId {
				selectedMember = members[i]
				break
			}
		}

		ctx.JSON(200, selectedMember)
	})

	// 社員情報の更新
	router.PUT("/members", func(ctx *gin.Context) {
		// Json形式で社員情報を取得する。
		// 社員情報が取得できなかった場合はエラー(400 Bad Request)
		var editMember member
		if err := ctx.ShouldBindJSON(&editMember); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i := 0; i < len(members); i++ {
			// IDの一致する社員情報を更新する。
			if members[i].Id == editMember.Id {
				members[i].Name = editMember.Name
				members[i].JoinDate = editMember.JoinDate
				break
			}
		}

		ctx.JSON(200, editMember)
	})

	// 社員情報の追加
	router.POST("/members", func(ctx *gin.Context) {
		// Json形式で社員情報を取得する。
		// 社員情報が取得できなかった場合はエラー(400 Bad Request)
		var addMember member
		if err := ctx.ShouldBindJSON(&addMember); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 社員情報の追加の際、名前のみ設定されているため、IDを設定する。
		addMember.Id = strconv.Itoa(id)
		id++
		// 社員情報に追加する。
		newMembers := append(members, addMember)
		members = newMembers

		ctx.JSON(200, addMember)
	})

	// 社員情報の削除
	router.DELETE("/members/:id", func(ctx *gin.Context) {
		selectedId := ctx.Param("id")

		for i := 0; i < len(members); i++ {
			// IDの一致する社員情報を削除する。
			if members[i].Id == selectedId {
				members = remove(members, i)
				break
			}
		}

		ctx.JSON(200, members)
	})

	router.Run(":8080")
}

func remove(s []member, i int) []member {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
