ASCII Art Web Dockerize is a Go-based web application that transforms user-input text into stylized ASCII art using fonts like Standard, Shadow, and Thinkertoy. The application features a web interface where users can input text, choose a style, customize font size and color, and instantly view the result. The app is containerized using Docker for consistent deployment across environments.

To run it:

STEP 1: Build the image
docker build -t ascii-art-web .

STEP 2: Run the container:
docker run -p 8080:8080 ascii-art-web

STEP 3: Open your browser and go to:
http://localhost:8080

# Project by Basel and Mujtaba