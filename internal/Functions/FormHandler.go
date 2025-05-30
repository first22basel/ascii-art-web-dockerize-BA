package BA

import (
	"html/template"
	"net/http"
)

type PageData struct {
	AsciiArt  string
	UserInput string
	Banner    string
	FontSize  string
	FontColor string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	// Check if MainPage.html is existed
	err := EnsureFile("/internal/frontend/MainPage.html", "https://raw.githubusercontent.com/first22basel/ascii-art-web-dockerize-BA/main/internal/frontend/MainPage.html")
	if err != nil {
		err := EnsureFile("/internal/frontend/500.html", "https://raw.githubusercontent.com/first22basel/ascii-art-web-dockerize-BA/main/internal/frontend/500.html")
		if err != nil {
			http.Error(w, "500 Server Error - Failed to recover 500.html", http.StatusInternalServerError)
			return
		}
		http.ServeFile(w, r, "../internal/frontend/500.html")
		return
	}

	// Check if there is HTTP post request
	if r.Method == http.MethodPost {
		// Parse data from HTTP post request
		r.ParseForm()
		input := r.FormValue("userinput")
		inputStyle := r.FormValue("banner")
		fontSize := r.FormValue("fontsize")
		fontColor := r.FormValue("color")

		// Map the banner file inside an array "fontMap"
		fontMap, err := LoadBanner(inputStyle)
		if err != nil {
			// Check if 500.html is existed
			err := EnsureFile("/internal/frontend/500.html", "https://raw.githubusercontent.com/first22basel/ascii-art-web-dockerize-BA/main/internal/frontend/500.html")
			if err != nil {
				http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
				return
			}
			http.ServeFile(w, r, "../internal/frontend/500.html")
			return
		}

		// Map user's inputs with fontMap
		asciiResult, err := PrintAscii(input, fontMap)
		if err != nil {
			// Check if 500.html is existed
			err := EnsureFile("/internal/frontend/400.html", "https://raw.githubusercontent.com/first22basel/ascii-art-web-dockerize-BA/main/internal/frontend/400.html")
			if err != nil {
				http.Error(w, "400 Server Error - Failed to recover 400.html", http.StatusBadRequest)
				return
			}
			http.ServeFile(w, r, "../internal/frontend/400.html")
			return
		}

		// Load MainPage.html to the website
		MainPageTemp, err := template.ParseFiles("/internal/frontend/MainPage.html")
		if err != nil {
			err := EnsureFile("../internal/frontend/MainPage.html", "https://raw.githubusercontent.com/first22basel/ascii-art-web-dockerize-BA/main/internal/frontend/MainPage.html")
			if err != nil {
				err := EnsureFile("/internal/frontend/500.html", "https://raw.githubusercontent.com/first22basel/ascii-art-web-dockerize-BA/main/internal/frontend/500.html")
				if err != nil {
					http.Error(w, "500 Server Error - Failed to recover 500.html", http.StatusInternalServerError)
					return
				}
			}
		}

		data := PageData{
			AsciiArt:  asciiResult,
			UserInput: input,
			Banner:    inputStyle,
			FontSize:  fontSize,
			FontColor: fontColor,
		}
		MainPageTemp.Execute(w, data)
		return
	}

	// Set default values for size, color, and banner
	data := PageData{
		FontSize:  "16px",
		FontColor: "#00ffcc",
		Banner:    "standard",
	}

	MainPageTemp, err := template.ParseFiles("../internal/frontend/MainPage.html")
	if err != nil {
		err := EnsureFile("/internal/frontend/MainPage.html", "https://raw.githubusercontent.com/first22basel/ascii-art-web-dockerize-BA/main/internal/frontend/MainPage.html")
		if err != nil {
			err := EnsureFile("/internal/frontend/500.html", "https://raw.githubusercontent.com/first22basel/ascii-art-web-dockerize-BA/main/internal/frontend/500.html")
			if err != nil {
				http.Error(w, "500 Server Error - Failed to recover 500.html", http.StatusInternalServerError)
				return
			}
		}
	}
	MainPageTemp.Execute(w, data)
}
