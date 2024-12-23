package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sstallion/go-hid"
	"io"
	"log"
	"net/http"
	"strconv"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	swapKeyCodeAndName()
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) LoadKeymapFromKeyboard() {
	// Initialize the hid package.
	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	// Open the device using the VID and PID.
	openConnectedHIDDevices()

	initBlankTOML()

	fmt.Println("--- Current Hardware Layout ScanCode ---")
	fmt.Println("::Layout1::")
	fmt.Println("  Normal ->")
	loadKeymap(0, 0x11)
	fmt.Println("  Upper ->")
	loadKeymap(0, 0x12)
	fmt.Println("  Stick ->")
	loadKeymap(0, 0x13)
	//fmt.Println("  Led ->")
	//loadKeymap(0, 0x14)
	//fmt.Println("  Intensity ->")
	//loadKeymap(0, 0x15)
	//fmt.Println("")
	fmt.Println("::Layout2::")
	fmt.Println("  Normal ->")
	loadKeymap(0, 0x19)
	fmt.Println("  Upper ->")
	loadKeymap(0, 0x1A)
	fmt.Println("  Stick ->")
	loadKeymap(0, 0x1B)
	//fmt.Println("  Led ->")
	//loadKeymap(0, 0x1C)
	//fmt.Println("  Intensity ->")
	//loadKeymap(0, 0x1D)
	//fmt.Println("")

}

func LoadTOML(inputdata string) {
	isInitLayout = true
	
	fmt.Println("-- Remap Layout ScanCode ---")

	_, err = toml.Decode(inputdata, &layouts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("::Layout1::")
	fmt.Println("  Normal ->")
	for i := 0; i < maxRows; i++ {
		fmt.Println(layouts.Layout1.Normal[i])
	}
	for i := 0; i < maxRows; i++ {
		fmt.Print("[ ")
		for j := 0; j < maxColumns; j++ {
			if KEYCODE[layouts.Layout1.Normal[i][j][0]] == 0x00 {
				if err := hid.Exit(); err != nil {
					log.Fatal(err)
				}
			} else {
				if modifiers, err := strconv.ParseUint(layouts.Layout1.Normal[i][j][1], 2, 8); err == nil {
					fmt.Printf("{0x%02X, 0x%02X} ", KEYCODE[layouts.Layout1.Normal[i][j][0]], modifiers)
				}
			}
		}
		fmt.Println("]")
	}
	if len(layouts.Layout1.Upper) != 0 {
		fmt.Println("  Upper ->")
		for i := 0; i < maxRows; i++ {
			fmt.Println(layouts.Layout1.Upper[i])
		}
		for i := 0; i < maxRows; i++ {
			fmt.Print("[ ")
			for j := 0; j < maxColumns; j++ {
				if KEYCODE[layouts.Layout1.Upper[i][j][0]] == 0x00 {
					if err := hid.Exit(); err != nil {
						log.Fatal(err)
					}
				} else {
					if modifiers, err := strconv.ParseUint(layouts.Layout1.Upper[i][j][1], 2, 8); err == nil {
						fmt.Printf("{0x%02X, 0x%02X} ", KEYCODE[layouts.Layout1.Upper[i][j][0]], modifiers)
					}
				}
			}
			fmt.Println("]")
		}
	}
	if len(layouts.Layout1.Stick) != 0 {
		fmt.Println("  Stick ->")
		for i := 0; i < 2; i++ {
			fmt.Println(layouts.Layout1.Stick[i])
		}
		for i := 0; i < 2; i++ {
			fmt.Print("[ ")
			for j := 0; j < 9; j++ {
				if KEYCODE[layouts.Layout1.Stick[i][j][0]] == 0x00 {
					if err := hid.Exit(); err != nil {
						log.Fatal(err)
					}
				} else {
					if modifiers, err := strconv.ParseUint(layouts.Layout1.Stick[i][j][1], 2, 8); err == nil {
						fmt.Printf("{0x%02X, 0x%02X} ", KEYCODE[layouts.Layout1.Stick[i][j][0]], modifiers)
					}
				}
			}
			fmt.Println("]")
		}
	}
	if len(layouts.Layout1.Led) != 0 {
		fmt.Println("  Led ->")
		for i := 0; i < 3; i++ {
			fmt.Print("[ ")
			for j := 0; j < 3; j++ {
				fmt.Printf("0x%02X ", layouts.Layout1.Led[i][j])
			}
			fmt.Println("]")
		}
	}
	if len(layouts.Layout1.Intensity) != 0 {
		fmt.Println("  Intensity ->")
		fmt.Printf("%f", layouts.Layout1.Intensity)
		fmt.Println("]")
	}
	fmt.Println("")
	fmt.Println("::Layout2::")
	fmt.Println("  Normal ->")
	for i := 0; i < maxRows; i++ {
		fmt.Println(layouts.Layout2.Normal[i])
	}
	for i := 0; i < maxRows; i++ {
		fmt.Print("[ ")
		for j := 0; j < maxColumns; j++ {
			if modifiers, err := strconv.ParseUint(layouts.Layout2.Normal[i][j][1], 2, 8); err == nil {
				fmt.Printf("{0x%02X, 0x%02X} ", KEYCODE[layouts.Layout2.Normal[i][j][0]], modifiers)
			}
		}
		fmt.Println("]")
	}
	if len(layouts.Layout2.Upper) != 0 {
		fmt.Println("  Upper ->")
		for i := 0; i < maxRows; i++ {
			fmt.Println(layouts.Layout2.Upper[i])
		}
		for i := 0; i < maxRows; i++ {
			fmt.Print("[ ")
			for j := 0; j < maxColumns; j++ {
				if modifiers, err := strconv.ParseUint(layouts.Layout2.Upper[i][j][1], 2, 8); err == nil {
					fmt.Printf("{0x%02X, 0x%02X} ", KEYCODE[layouts.Layout2.Upper[i][j][0]], modifiers)
				}
			}
			fmt.Println("]")
		}
	}
	if len(layouts.Layout2.Stick) != 0 {
		fmt.Println("  Stick ->")
		for i := 0; i < 2; i++ {
			fmt.Println(layouts.Layout2.Stick[i])
		}
		for i := 0; i < 2; i++ {
			fmt.Print("[ ")
			for j := 0; j < 9; j++ {
				if modifiers, err := strconv.ParseUint(layouts.Layout2.Stick[i][j][1], 2, 8); err == nil {
					fmt.Printf("{0x%02X, 0x%02X} ", KEYCODE[layouts.Layout2.Stick[i][j][0]], modifiers)
				}
			}
			fmt.Println("]")
		}
	}
	if len(layouts.Layout2.Led) != 0 {
		fmt.Println("  Led ->")
		for i := 0; i < 3; i++ {
			fmt.Print("[ ")
			for j := 0; j < 3; j++ {
				fmt.Printf("0x%02X ", layouts.Layout2.Led[i][j])
			}
			fmt.Println("]")
		}
	}
	if len(layouts.Layout2.Intensity) != 0 {
		fmt.Println("  Intensity ->")
		fmt.Printf("%f", layouts.Layout2.Intensity)
		fmt.Println("]")
	}
	fmt.Println("")
}

func (a *App) GetNormalKeyOfLayout1(row int, col int) string {
	fmt.Printf("row=%d, col=%d %s(%s)\n", row, col, layouts.Layout1.Normal[row][col][0], layouts.Layout1.Normal[row][col][1])

	return KEYTOP[layouts.Layout1.Normal[row][col][0]]
}

func (a *App) GetNormalKeyOfLayout2(row int, col int) string {
	//fmt.Printf("row=%d, col=%d %s(%s)\n", row, col, layouts.Layout2.Normal[row][col][0], layouts.Layout2.Normal[row][col][1])

	return KEYTOP[layouts.Layout2.Normal[row][col][0]]
}

func (a *App) GetNormalModifiersOfLayout1(row int, col int) uint64 {
	//fmt.Printf("row=%d, col=%d %s\n", row, col, layouts.Layout1.Normal[row][col][1])

	if modifiers, err := strconv.ParseUint(layouts.Layout1.Normal[row][col][1], 2, 8); err == nil {
		return modifiers
	} else {
		return 0
	}
}

func (a *App) GetNormalModifiersOfLayout2(row int, col int) uint64 {
	//fmt.Printf("row=%d, col=%d %s\n", row, col, layouts.Layout2.Normal[row][col][1])

	if modifiers, err := strconv.ParseUint(layouts.Layout2.Normal[row][col][1], 2, 8); err == nil {
		return modifiers
	} else {
		return 0
	}
}

func (a *App) GetUpperKeyOfLayout1(row int, col int) string {
	//fmt.Printf("row=%d, col=%d %s(%s)\n", row, col, layouts.Layout1.Normal[row][col][0], layouts.Layout1.Normal[row][col][1])

	return KEYTOP[layouts.Layout1.Upper[row][col][0]]
}

func (a *App) GetUpperKeyOfLayout2(row int, col int) string {
	//fmt.Printf("row=%d, col=%d %s(%s)\n", row, col, layouts.Layout2.Normal[row][col][0], layouts.Layout2.Normal[row][col][1])

	return KEYTOP[layouts.Layout2.Upper[row][col][0]]
}

func (a *App) GetUpperModifiersOfLayout1(row int, col int) uint64 {
	//fmt.Printf("row=%d, col=%d %s\n", row, col, layouts.Layout1.Normal[row][col][1])

	if modifiers, err := strconv.ParseUint(layouts.Layout1.Upper[row][col][1], 2, 8); err == nil {
		return modifiers
	} else {
		return 0
	}
}

func (a *App) GetUpperModifiersOfLayout2(row int, col int) uint64 {
	//fmt.Printf("row=%d, col=%d %s\n", row, col, layouts.Layout2.Normal[row][col][1])

	if modifiers, err := strconv.ParseUint(layouts.Layout2.Upper[row][col][1], 2, 8); err == nil {
		return modifiers
	} else {
		return 0
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	var fileContent string

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/toml" {
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	fileContent = string(body)

	// Process the file conten
	result, err := ProcessTOMLFile(fileContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	LoadTOML(fileContent)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(map[string]string{"result": result}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ProcessTOMLFile(content string) (string, error) {
	var parsedData map[string]interface{}

	err := toml.Unmarshal([]byte(content), &parsedData)
	if err != nil {
		return "", fmt.Errorf("TOML data parse failed: %w", err)
	}

	return "TOML file is processed correctly", nil
}

func (a *App) CheckHID() [3]string {
	// Initialize the hid package.
	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	// Open the device using the VID and PID.
	openConnectedHIDDevices()

	return checkHid()
}

func (a *App) Restart() {
	// Initialize the hid package.
	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	// Open the device using the VID and PID.
	openConnectedHIDDevices()

	restart(0)
}

func (a *App) FactoryReset() {
	// Initialize the hid package.
	if err := hid.Init(); err != nil {
		log.Fatal(err)
	}

	// Open the device using the VID and PID.
	openConnectedHIDDevices()

	factoryReset(0)
}
