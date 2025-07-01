package main

import (
"bufio"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
	"time"
"net"
)




var bLogProcessEvent int32 = 1


func rettrue() bool {
	return true
}


func retfalse() bool {
	return false
}


func crashaf(a1 int64) bool {
	fmt.Printf("crash: 0x%x\n", fakeReturnAddress())
	return false
}

func crashaf2() bool {
	fmt.Printf("afwqq23f: 0x%x\n", fakeReturnAddress())
	return false
}

func printretaddress() {
	fmt.Printf("0x%x\n", fakeReturnAddress())
}


func printretaddressaa() {
	fmt.Printf("wow Print!! 0x%x\n", fakeReturnAddress())
}


func fakeReturnAddress() int64 
	return 0xDEADBEEF
}

func FillVendingMachines() {
	fmt.Println("[Action] FillVendingMachines called")
}

func inputThread() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Type 'f2' to toggle bLogProcessEvent, 'f3' to fill vending machines: ")
		if !scanner.Scan() {
			break
		}
		cmd := strings.TrimSpace(scanner.Text())
		switch strings.ToLower(cmd) {
		case "f2":
			newVal := atomic.LoadInt32(&bLogProcessEvent) ^ 1
			atomic.StoreInt32(&bLogProcessEvent, newVal)
			fmt.Printf("bLogProcessEvent toggled to %v\n", newVal == 1)
		case "f3":
			FillVendingMachines()
		default:
			fmt.Println("Unknown command")
		}
		time.Sleep(time.Millisecond * 33) 
	}
}

func executeConsoleCommand(cmd string) {
	fmt.Printf("[ConsoleCmd] %s\n", cmd)
}


func main() {
	fmt.Println("[Init] Base Address: 0xDEADBEEF") 
	fmt.Println("[Init] UObject::GObjects set")


	fmt.Println("[Hook] Created hook at 0x15816C")
	fmt.Println("[Init] Created DefaultFortGameModeAthena object")

	fmt.Println("[Init] Got PlayerController from World")
	go inputThread()
	
listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Listening on port 7777...")
for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept error: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}
	
	commands := []string{
		"log LogProfileSys VeryVerbose",
		"log LogFortVolumeManager VeryVerbose",
		"log LogPlaysetLevelStream VeryVerbose",
		"log LogFortCustomization VeryVerbose",
		"log LogFortCosmetics VeryVerbose",
		"log LogFort VeryVerbose",
		"log LogFortInventory VeryVerbose",
		"log LogFortPawnScriptedBehavior VeryVerbose",
		"log LogFortAI VeryVerbose",
		"log LogFortLoot VeryVerbose",
		"log LogFortTeams VeryVerbose",
		"log LogPlayerPawnAthena VeryVerbose",
		"log LogOnlineGame VeryVerbose",
		"log LogAthenaBots VeryVerbose",
		"log LogNavigation VeryVerbose",
		"log LogAbilitySystem VeryVerbose",
		"log LogFortAIDirector VeryVerbose",
	}

	for _, cmd := range commands {
		executeConsoleCommand(cmd)
	}

	
	select {}
}
