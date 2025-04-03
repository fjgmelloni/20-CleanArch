package webserver

import "fmt"

type WebServerStarter struct {
    WebServer *WebServer
}

func NewWebServerStarter(webServer *WebServer) *WebServerStarter {
    return &WebServerStarter{
        WebServer: webServer,
    }
}

// Start method to start the web server
func (s *WebServerStarter) Start() {
    fmt.Println("Starting web server on port", s.WebServer.WebServerPort)
    s.WebServer.Start()
}