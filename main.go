package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func handleFiber(path string, collection []fs.DirEntry) *fiber.App {

	engine := html.New("./views", ".tpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {

		filelist := []fiber.Map{}

		for _, entry := range collection {
			f := fiber.Map{
				"title":      entry.Name(),
				"url":        c.Request().URI().String() + "play/" + entry.Name(),
				"coverImage": "",
			}
			filelist = append(filelist, f)
		}

		return c.JSON(filelist)
	})

	app.Get("/playlist", func(c *fiber.Ctx) error {
		//return c.Render("index", collection)
		baseUrl := string(c.Request().URI().Scheme()) + "://" + string(c.Request().URI().Host())
		str := "#EXTM3U\n\n\n"
		for _, entry := range collection {

			str += "#EXTINF:0 " + entry.Name() + "\n"
			str += (baseUrl + "/play/" + entry.Name()) + "\n\n\n"

		}
		return c.SendString(str)
	})

	app.Get("/list", func(c *fiber.Ctx) error {

		filelist := []fiber.Map{}

		for _, entry := range collection {

			info, err1 := entry.Info()
			if err1 == nil {
				fmt.Println(err1)
			}

			f := fiber.Map{
				"title": entry.Name(),
				"url":   path,
				"size":  (info.Size() / 1024) / 1024,
			}
			filelist = append(filelist, f)
		}

		return c.Render("index", fiber.Map{"list": filelist})
	})

	app.Get("/play/:filename", func(ctx *fiber.Ctx) error {

		filename := ctx.Params("filename")

		filePath := path + "/" + filename
		return ctx.SendFile(filePath, true)
	})

	return app

}

func parseDirectory(location string) []fs.DirEntry {

	files, err := os.ReadDir(location)

	extenstion := map[string]bool{".mp4": true, ".mov": true, ".wmv": true, ".avi": true, ".mkv": true, ".webm": true, ".mpeg": true, ".mpg": true, ".flv": true}

	var collection []fs.DirEntry
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range files {

		ext := path.Ext(e.Name())

		if ok := extenstion[ext]; ok && !e.IsDir() {

			collection = append(collection, e)
		}

	}
	return collection
}

func main() {
	//app := gin.Default();
	var path string
	flag.StringVar(&path, "path", "./", "path where to run")
	var host string
	flag.StringVar(&host, "host", "127.0.0.1", "host or ip address run the application")
	var port string
	flag.StringVar(&port, "port", "8989", "port number to run application")
	flag.Parse()
	result := parseDirectory(path)

	app := handleFiber(path, result)

	err := app.Listen(host + ":" + port)
	if err != nil {
		log.Fatal("unable to start the application", err)
	}
}
