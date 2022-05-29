//เริ่มต้นการเขียนโรแกรมภาษา go
package main //ประกาศ
import "fmt" //เรียกใช้งาน library

func main() {
	fmt.Print("Hello World")
}

//การใช้ Go command
//ใช้คำสั่ง go mod init github.com/yadaaaa/ชื่ออะไรก็ได้ ลงใน terminal เพื่อเก็บพาท
//พิมพ์ pk แล้วกด tab เพื่อขึ้นโครงสร้างของโค้ด
//พิมพ์ fp แล้วกด tab เพื่อขึ้น fmt.Print("Hello World")
//พิมพ์ go run main.go เพื่อรันโปรแกรมไฟล์นี้
//ctrl+ship+p แล้วหาคำว่า Insert Snippet เพื่อพิมพ์เร็วขึ้น!
//ใช้คำสั่ง go build .\main.go เพื่อใช้เป็น library .exe รันในเครื่องอื่นได้
//ใช้สำคั่ง go get github.com/liudng/dogo เพื่อเรียก package จากด้านนอก -> dogo รันอัตโนมัติทุกครั้งที่มีการเปลี่ยนแปลงโค้ด และสร้างไฟล์ go.sum ขึ้นมา อย่าลืมสร้าง dogo.json
//...
