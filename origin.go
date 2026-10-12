package main

import (
    "crypto/tls"
    "fmt"
    "log"
    "net/url"
    "os"
    "os/signal"
    "time"

    "github.com/gorilla/websocket"
)

func main() {

    target := "wisp server"
    path := "/" =
    useTLS := true =
    

    payload := "\n\n\r\n\n\r\n\n\r\n\n\rwhats up g\n\n\n"
    scheme := "ws"
    if useTLS {
        scheme = "wss"
    }
    
    u := url.URL{
        Scheme: scheme,
        Host:   target,
        Path:   path,
    }
    
    customorigin := fmt.Sprintf(payload)
    
    headers := map[string][]string{
        "Origin":     {customorigin},
        "User-Agent": {fmt.Sprintf("Mozilla/5.0 (X11; Linux x86_64) cucknet/%s", time.Now().Format("2006"))},
        "X-Real-IP":  {"6.7.6.7"}, // doesnt work its for fun
    }
    
    log.Printf("target: %s", u.String())
    log.Printf("payload size: %d bytes", len(payload))
    
    dialer := websocket.Dialer{
        HandshakeTimeout: 10 * time.Second,
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true,
            ServerName:         target,
        },
    }
    
    conn, resp, err := dialer.Dial(u.String(), headers)
    if err != nil {
        log.Fatalf("connection failed: %v", err)
    }
    defer conn.Close()
    
    log.Printf("connected: %s", resp.Status)
    
    packet := []byte{
        0x01, // connect
        0x00, 0x00, 0x00, 0x01, 
        0x01, //tcp
        0x00, 0x50,
    }
    packet = append(packet, []byte("example.com")...)
    
    if err := conn.WriteMessage(websocket.BinaryMessage, packet); err != nil {
        log.Printf(" send failed: %v", err)
    } else {
        log.Println("connect packet sent")
    }
    
    conn.SetReadDeadline(time.Now().Add(3 * time.Second))
    msgType, msg, err := conn.ReadMessage()
    if err == nil {
        log.Printf("responded: %d bytes (type: %d)", len(msg), msgType)
        if len(msg) >= 5 {
            packetType := msg[0]
            log.Printf("packet type: 0x%02x", packetType)
        }
    } else {
        log.Printf("no response: %v", err)
    }
    
    log.Println("\nodin din DONE")
    
    time.Sleep(2 * time.Second)
    
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
}
