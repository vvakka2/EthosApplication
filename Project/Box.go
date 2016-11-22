package main
import (
"ethos/syscall"
ethosLog "ethos/log"
"ethos/efmt"
"ethos/ethos"
"math"
)
func main() {
me := syscall.GetUser()
path := "/user/"+me+"/myDir/"
status := ethosLog.RedirectToLog("box")
if status != syscall.StatusOk {
efmt.Fprintf(syscall.Stderr,"Error opening %v: %v\n",path,status)
syscall.Exit(syscall.StatusOk)
}
data := BoxType {1,1,3,3}
efmt.Println("Values x1:",data.x1,"y1:",data.y1)
efmt.Println("Second Values X2: ",data.x2," Y2: ",data.y2)
distance := math.Sqrt(math.Pow((data.x2-data.x1),2)+math.Pow((data.y2-data.y1),2))
newBox := BoxType {data.x1+distance,data.y1+distance,data.x2+distance,data.y2+distance}
efmt.Println("Distance: ",distance)
efmt.Println("Updated Values x1: ",newBox.x1," y1: ",newBox.y1)
efmt.Println("Updated Values X2: ",newBox.x2," y2: ",newBox.y2) 
fd,status := ethos.OpenDirectoryPath(path)
data.Write(fd)
efmt.Fprint(syscall.Stderr,"Success")
}
