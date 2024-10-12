package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func setupRouter(r *gin.Engine) {
	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("static", "./static")
	//r.GET("/telcsis/", bookIndexHandler)
	r.GET("/telcsifarmer", telcsiNewGetHandler)
	r.GET("/telcsifarmer/download/:excelname", telcsiDownloadGetHandler)
	r.GET("/telcsifarmer/working", telcsiWorkingGetHandler)
	r.POST("/telcsifarmer", telcsikNewPostHandler)
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/telcsifarmer")
	})
}

func telcsiNewGetHandler(c *gin.Context) {
	checkExcels()
	c.HTML(http.StatusOK, "telcsifarmer/new.html", gin.H{})
}

func telcsiWorkingGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "telcsifarmer/working.html", gin.H{})
}

func telcsiDownloadGetHandler(c *gin.Context) {
	//str, ex := c.Get("excelname")
	// fmt.Printf("str: %v and ex: %v", str, ex)
	// if !ex {
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	return
	// }
	str := c.Param("excelname")
	//fmt.Printf("str: %v \n", str)
	excelName := "scrapings/" + str
	f, err := excelize.OpenFile(excelName)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var b bytes.Buffer
	if err := f.Write(&b); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+excelName)
	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())

}

func telcsikNewPostHandler(c *gin.Context) {
	telcsi := &TelcsiFarmolas{}
	if err := c.Bind(telcsi); err != nil {
		// Note: if there's a bind error, Gin will call
		// c.AbortWithError. We just need to return here.
		fmt.Printf(" \n\n\nHiba:  %v \n\n\n", err)
		return
	}
	// FIXME: There's a better way to do this validation!
	if telcsi.MinPrice == 0 || telcsi.MaxPrice == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// c.Redirect(http.StatusFound, "/telcsifarmer/working")
	//IDE KELL KÓD, LOGIKA
	fmt.Printf("Telcsik min_price: %d and Telcsik max_price: %d\n", telcsi.MinPrice, telcsi.MaxPrice)
	fmt.Printf("Telcsik Fontik: %s, Telcsik maradék: %s\n", telcsi.ImportantPhones, telcsi.NeutralPhones)
	excelName := telcsiworker(telcsi)

	c.Redirect(http.StatusFound, "/telcsifarmer/download/"+excelName)
}

func checkExcels() {
	const FOLDER = "./scrapings"
	files, _ := os.ReadDir(FOLDER)
	//fmt.Println(len(files))
	for _, f := range files {
		filename := f.Name()
		//fmt.Println(FOLDER + "/" + filename)
		fileInfo, err := os.Stat(FOLDER + "/" + filename)
		if err != nil {
			return
		}
		modificationTime := fileInfo.ModTime()
		//fmt.Printf("ModTime of %s is %v and thus difference is %d\n", filename, modificationTime, time.Since(modificationTime))
		if time.Since(modificationTime) > 2*time.Minute {
			e := os.Remove(FOLDER + "/" + filename)
			if e != nil {
				log.Fatal(e)
			}
			fmt.Printf("Deleted file %s because it was %d old\n", filename, time.Since(modificationTime))
		}

	}
}

func main() {
	r := gin.Default()
	setupRouter(r)
	err := r.Run(":6969")
	if err != nil {
		log.Fatalf("gin Run error: %s", err)
	}

}
