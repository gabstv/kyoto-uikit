package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var taiwindconfig = []byte(`const colors = require('tailwindcss/colors')
const dc = require('tailwindcss/defaultConfig')

module.exports = {
    purge: [
        "../*.go",
        "../*.html",
        "../*.gohtml",
        "../**/*.go",
        "../**/*.html",
        "../**/*.gohtml",
        ".twui/*.html",
    ],
    mode: 'jit',
    darkMode: 'media', // or 'media' or 'class'
    theme: {
        extend: {
            fontFamily: {
                serif: ['Times'].concat(dc.theme.fontFamily.serif)
            },
            colors: {
                gray: colors.trueGray
            }
        },
    },
    variants: {},
    plugins: [],
}`)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: twui <static folder>")
		os.Exit(1)
	}
	staticf := os.Args[1]
	os.MkdirAll(filepath.Join(staticf, ".twui"), 0744)
	if _, err := os.Stat(filepath.Join(staticf, ".gitignore")); os.IsNotExist(err) {
		ignored := []byte(`node_modules/
dist
.twui/
`)
		if err := os.WriteFile(filepath.Join(staticf, ".gitignore"), ignored, 0644); err != nil {
			panic(err)
		}
	}
	if _, err := os.Stat(filepath.Join(staticf, "tailwind.config.js")); os.IsNotExist(err) {

		if err := os.WriteFile(filepath.Join(staticf, "tailwind.config.js"), taiwindconfig, 0644); err != nil {
			panic(err)
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	defer os.Chdir(wd)

	if err := os.Chdir(staticf); err != nil {
		fmt.Println("failed to finish")
	}
	if _, err := os.Stat("package.json"); os.IsNotExist(err) {

		d := []byte(`{
			"name": "kyoto",
			"version": "1.0.0",
			"description": "",
			"main": "index.js",
			"scripts": {
			  "build": "yarn build:css && yarn build:img && yarn build:meta",
			  "build:css": "tailwindcss -c tailwind.config.js -o dist/css/tailwind.min.css",
			  "build:css:w": "tailwindcss -w -c tailwind.config.js -o dist/css/tailwind.min.css",
			  "build:img": "(rm -rf dist/img && mkdir -p dist/img && cp -r img/* dist/img/) || true",
			  "build:meta": "(rm -rf dist/meta && mkdir -p dist/meta && cp -r meta/* dist/meta/) || true"
			},
			"author": "",
			"license": "ISC",
			"dependencies": {
			  "autoprefixer": "^10.3.1",
			  "postcss": "^8.3.6",
			  "tailwindcss": "^2.2.7"
			}
		  }`)

		if err := os.WriteFile("package.json", d, 0644); err != nil {
			fmt.Println(err.Error())
			return
		}

		cmd := exec.Command("yarn", "--yes")
		x, err := cmd.CombinedOutput()
		fmt.Println(x)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("ready")
	}
}
