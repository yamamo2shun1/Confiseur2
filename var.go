package main

var KEY_LIST [103]string = [103]string{
	"A",           // a
	"B",           // b
	"C",           // c
	"D",           // d
	"E",           // e
	"F",           // f
	"G",           // g
	"H",           // h
	"I",           // i
	"J",           // j
	"K",           // k
	"L",           // l
	"M",           // m
	"N",           // n
	"O",           // o
	"P",           // p
	"Q",           // q
	"R",           // r
	"S",           // s
	"T",           // t
	"U",           // u
	"V",           // v
	"W",           // w
	"X",           // x
	"Y",           // y
	"Z",           // z
	"1",           // 1!
	"2",           // 2@
	"3",           // 3#
	"4",           // 4$
	"5",           // 5%
	"6",           // 6^
	"7",           // 7&
	"8",           // 8*
	"9",           // 9(
	"0",           // 0)
	"Enter",       // Enter
	"Esc",         // Escape
	"BS",          // Backspace
	"Tab",         // Tab
	"Space",       // Space
	"Minus",       // -_
	"Equal",       // =+
	"O_SBracket",  // [{
	"C_SBracket",  // ]}
	"Backslash",   // \|
	"Colon",       // ;:
	"Apostrophe",  // '"
	"Backquote",   // `~
	"Comma",       // ,<
	"Period",      // .>
	"Slash",       // /?
	"CapsLock",    // Caps Lock
	"F1",          // F1
	"F2",          // F2
	"F3",          // F3
	"F4",          // F4
	"F5",          // F5
	"F6",          // F6
	"F7",          // F7
	"F8",          // F8
	"F9",          // F9
	"F10",         // F10
	"F11",         // F11
	"F12",         // F12
	"PrintScreen", // Print Screen
	"ScrollLock",  // Scroll Lock
	"Pause",       // Pause
	"Ins",         // Insert
	"Home",        // Home
	"PageUp",      // Page Up
	"Del",         // Delete
	"End",         // End
	"PageDown",    // Page Down
	"Right",       // Right
	"Left",        // Left
	"Down",        // Down
	"Up",          // Up
	"NumLock",     // Num Lock
	"Katakana",    // カタカナ ひらがな
	"Yen",         // ￥|
	"Henkan",      // 変換
	"Muhenkan",    // 無変換
	"L_Control",   // Left Control
	"L_Shift",     // Left Shift
	"L_Alt",       // Left Alt
	"L_Gui",       // Left GUI
	"R_Control",   // Right Control
	"R_Shift",     // Right Shift
	"R_Alt",       // Right Alt
	"R_Gui",       // Right GUI
	"M_LBTN",      // Mouse Left Button
	"M_RBTN",      // Mouse Right Button
	"M_WHEEL",     // Mouse Wheel
	"Reset",       // Reset
	"XF_CUT1",     // XFader Cut Ch.1(Phono/Line In)
	"XF_CUT2",     // XFader Cut Ch.2(USB)
	"MGain_Up",    // Master Gain Up
	"MGain_Down",  // Master Gain Down
	"Upper",       // Upper
	"LNPH",        // Line Phono Switch
	"LAYOUT",      // Layout Switch
	"Null",
}

var KEYCODE = map[string]byte{
	"A":           0x04, // a
	"B":           0x05, // b
	"C":           0x06, // c
	"D":           0x07, // d
	"E":           0x08, // e
	"F":           0x09, // f
	"G":           0x0A, // g
	"H":           0x0B, // h
	"I":           0x0C, // i
	"J":           0x0D, // j
	"K":           0x0E, // k
	"L":           0x0F, // l
	"M":           0x10, // m
	"N":           0x11, // n
	"O":           0x12, // o
	"P":           0x13, // p
	"Q":           0x14, // q
	"R":           0x15, // r
	"S":           0x16, // s
	"T":           0x17, // t
	"U":           0x18, // u
	"V":           0x19, // v
	"W":           0x1A, // w
	"X":           0x1B, // x
	"Y":           0x1C, // y
	"Z":           0x1D, // z
	"1":           0x1E, // 1!
	"2":           0x1F, // 2@
	"3":           0x20, // 3#
	"4":           0x21, // 4$
	"5":           0x22, // 5%
	"6":           0x23, // 6^
	"7":           0x24, // 7&
	"8":           0x25, // 8*
	"9":           0x26, // 9(
	"0":           0x27, // 0)
	"Enter":       0x28, // Enter
	"Esc":         0x29, // Escape
	"BS":          0x2A, // Backspace
	"Tab":         0x2B, // Tab
	"Space":       0x2C, // Space
	"Minus":       0x2D, // -_
	"Equal":       0x2E, // =+
	"O_SBracket":  0x2F, // [{
	"C_SBracket":  0x30, // ]}
	"Backslash":   0x31, // \|
	"Colon":       0x33, // ;:
	"Apostrophe":  0x34, // '"
	"Backquote":   0x35, // `~
	"Comma":       0x36, // ,<
	"Period":      0x37, // .>
	"Slash":       0x38, // /?
	"CapsLock":    0x39, // Caps Lock
	"F1":          0x3A, // F1
	"F2":          0x3B, // F2
	"F3":          0x3C, // F3
	"F4":          0x3D, // F4
	"F5":          0x3E, // F5
	"F6":          0x3F, // F6
	"F7":          0x40, // F7
	"F8":          0x41, // F8
	"F9":          0x42, // F9
	"F10":         0x43, // F10
	"F11":         0x44, // F11
	"F12":         0x45, // F12
	"PrintScreen": 0x46, // Print Screen
	"ScrollLock":  0x47, // Scroll Lock
	"Pause":       0x48, // Pause
	"Ins":         0x49, // Insert
	"Home":        0x4A, // Home
	"PageUp":      0x4B, // Page Up
	"Del":         0x4C, // Delete
	"End":         0x4D, // End
	"PageDown":    0x4E, // Page Down
	"Right":       0x4F, // Right
	"Left":        0x50, // Left
	"Down":        0x51, // Down
	"Up":          0x52, // Up
	"NumLock":     0x53, // Num Lock
	"Katakana":    0x88, // カタカナ ひらがな
	"Yen":         0x89, // ￥|
	"Henkan":      0x8A, // 変換
	"Muhenkan":    0x8B, // 無変換
	"L_Control":   0xE0, // Left Control
	"L_Shift":     0xE1, // Left Shift
	"L_Alt":       0xE2, // Left Alt
	"L_Gui":       0xE3, // Left GUI
	"R_Control":   0xE4, // Right Control
	"R_Shift":     0xE5, // Right Shift
	"R_Alt":       0xE6, // Right Alt
	"R_Gui":       0xE7, // Right GUI
	"M_LBTN":      0xF0, // Mouse Left Button
	"M_RBTN":      0xF1, // Mouse Right Button
	"M_WHEEL":     0xF2, // Mouse Wheel
	"Reset":       0xF7, // Reset
	"XF_CUT1":     0xF8, // XFader Cut Ch.1(Phono/Line In)
	"XF_CUT2":     0xF9, // XFader Cut Ch.2(USB)
	"MGain_Up":    0xFA, // Master Gain Up
	"MGain_Down":  0xFB, // Master Gain Down
	"Upper":       0xFC, // Upper
	"LNPH":        0xFD, // Line Phono Switch
	"LAYOUT":      0xFE, // Layout Switch
	"Null":        0xFF,
}

var KEYTOP = map[string]string{
	"A":           "A",                              // a
	"B":           "B",                              // b
	"C":           "C",                              // c
	"D":           "D",                              // d
	"E":           "E",                              // e
	"F":           "F",                              // f
	"G":           "G",                              // g
	"H":           "H",                              // h
	"I":           "I",                              // i
	"J":           "J",                              // j
	"K":           "K",                              // k
	"L":           "L",                              // l
	"M":           "M",                              // m
	"N":           "N",                              // n
	"O":           "O",                              // o
	"P":           "P",                              // p
	"Q":           "Q",                              // q
	"R":           "R",                              // r
	"S":           "S",                              // s
	"T":           "T",                              // t
	"U":           "U",                              // u
	"V":           "V",                              // v
	"W":           "W",                              // w
	"X":           "X",                              // x
	"Y":           "Y",                              // y
	"Z":           "Z",                              // z
	"1":           "1!",                             // 1!
	"2":           "2@",                             // 2@
	"3":           "3#",                             // 3#
	"4":           "4$",                             // 4$
	"5":           "5%",                             // 5%
	"6":           "6^",                             // 6^
	"7":           "7&",                             // 7&
	"8":           "8*",                             // 8*
	"9":           "9(",                             // 9(
	"0":           "0)",                             // 0)
	"Enter":       "Enter",                          // Enter
	"Esc":         "Escape",                         // Escape
	"BS":          "Backspace",                      // Backspace
	"Tab":         "Tab",                            // Tab
	"Space":       "Space",                          // Space
	"Minus":       "-_",                             // -_
	"Equal":       "=+",                             // =+
	"O_SBracket":  "[{",                             // [{
	"C_SBracket":  "}]",                             // ]}
	"Backslash":   "\\|",                            // \|
	"Colon":       ";:",                             // ;:
	"Apostrophe":  "'\"",                            // '"
	"Backquote":   "`~",                             // `~
	"Comma":       ",<",                             // ,<
	"Period":      ".>",                             // .>
	"Slash":       "/?",                             // /?
	"CapsLock":    "Caps Lock",                      // Caps Lock
	"F1":          "F1",                             // F1
	"F2":          "F2",                             // F2
	"F3":          "F3",                             // F3
	"F4":          "F4",                             // F4
	"F5":          "F5",                             // F5
	"F6":          "F6",                             // F6
	"F7":          "F7",                             // F7
	"F8":          "F8",                             // F8
	"F9":          "F9",                             // F9
	"F10":         "F10",                            // F10
	"F11":         "F11",                            // F11
	"F12":         "F12",                            // F12
	"PrintScreen": "Print Screen",                   // Print Screen
	"ScrollLock":  "Scroll Lock",                    // Scroll Lock
	"Pause":       "Pause",                          // Pause
	"Ins":         "Insert",                         // Insert
	"Home":        "Home",                           // Home
	"PageUp":      "Page Up",                        // Page Up
	"Del":         "Delete",                         // Delete
	"End":         "End",                            // End
	"PageDown":    "Page Down",                      // Page Down
	"Right":       "Right",                          // Right
	"Left":        "Left",                           // Left
	"Down":        "Down",                           // Down
	"Up":          "Up",                             // Up
	"NumLock":     "Num Lock",                       // Num Lock
	"Katakana":    "カタカナ ひらがな",              // カタカナ ひらがな
	"Yen":         "￥|",                             // ￥|
	"Henkan":      "変換",                           // 変換
	"Muhenkan":    "無変換",                         // 無変換
	"M_LBTN":      "Mouse Left Button",              // Mouse Left Button
	"M_RBTN":      "Mouse Right Button",             // Mouse Right Button
	"M_WHEEL":     "Mouse Wheel",                    // Mouse Wheel
	"Reset":       "Reset",                          // Reset
	"XF_CUT1":     "XFader Cut Ch.1(Phono/Line In)", // XFader Cut Ch.1(Phono/Line In)
	"XF_CUT2":     "XFader Cut Ch.2(USB)",           // XFader Cut Ch.2(USB)
	"MGain_Up":    "Master Gain Up",                 // Master Gain Up
	"MGain_Down":  "Master Gain Down",               // Master Gain Down
	"Upper":       "Upper",                          // Upper
	"LNPH":        "Line Phono Switch",              // Line Phono Switch
	"LAYOUT":      "Layout Switch",                  // Layout Switch
	"Null":        "&nbsp;",
}
var KEYNAME = make(map[byte]string)

func swapKeyCodeAndName() {
	for key, value := range KEYCODE {
		KEYNAME[value] = key
	}
}
