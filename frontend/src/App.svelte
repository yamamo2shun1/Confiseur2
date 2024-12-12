<script>
    import {
        Button,
        Checkbox,
        Dropdown,
        DropdownItem,
        Dropzone,
        Label,
        Modal,
        RadioButton,
        TabItem,
        Table,
        TableBody,
        TableBodyCell,
        TableBodyRow,
        TableHead,
        TableHeadCell,
        Tabs
    } from 'flowbite-svelte';
    import {
        AngleDownOutline,
        AngleLeftOutline,
        AngleRightOutline,
        AngleUpOutline,
        ArrowDownOutline,
        ArrowLeftOutline,
        ArrowRightOutline,
        ArrowUpOutline,
        ChevronRightOutline,
        CogOutline,
        VolumeDownOutline,
        VolumeUpOutline
    } from "flowbite-svelte-icons";
    import tomlify from 'tomlify-j0.4';
    import {LogError, LogInfo} from "../wailsjs/runtime/runtime.js";
    import {
        GetNormalKeyOfLayout1,
        GetNormalKeyOfLayout2,
        GetNormalModifiersOfLayout1,
        GetNormalModifiersOfLayout2,
        GetUpperKeyOfLayout1,
        GetUpperKeyOfLayout2,
        GetUpperModifiersOfLayout1,
        GetUpperModifiersOfLayout2
    } from "../wailsjs/go/main/App.js";

    let radioGroup = "";
    let directionGroup = "";

    let selectedSettingTab = 0;

    let allKeyMaps = [
        {value: " ", name: '&nbsp;'},
        {value: "A", name: "A"},
        {value: "B", name: "B"},
        {value: "C", name: "C"},
        {value: "D", name: "D"},
        {value: "E", name: "E"},
        {value: "F", name: "F"},
        {value: "G", name: "G"},
        {value: "H", name: "H"},
        {value: "I", name: "I"},
        {value: "J", name: "J"},
        {value: "K", name: "K"},
        {value: "L", name: "L"},
        {value: "M", name: "M"},
        {value: "N", name: "N"},
        {value: "O", name: "O"},
        {value: "P", name: "P"},
        {value: "Q", name: "Q"},
        {value: "R", name: "R"},
        {value: "S", name: "S"},
        {value: "T", name: "T"},
        {value: "U", name: "U"},
        {value: "V", name: "V"},
        {value: "W", name: "W"},
        {value: "X", name: "X"},
        {value: "Y", name: "Y"},
        {value: "Z", name: "Z"},
        {value: "1", name: "1!"},
        {value: "2", name: "2@"},
        {value: "3", name: "3#"},
        {value: "4", name: "4$"},
        {value: "5", name: "5%"},
        {value: "6", name: "6^"},
        {value: "7", name: "7&"},
        {value: "8", name: "8*"},
        {value: "9", name: "9("},
        {value: "0", name: "0)"},
        {value: "Enter", name: "Enter"},
        {value: "Esc", name: "Escape"},
        {value: "BS", name: "Backspace"},
        {value: "Tab", name: "Tab"},
        {value: "Space", name: "Space"},
        {value: "-_", name: "-_"},
        {value: "=+", name: "=+"},
        {value: "[{", name: "[{"},
        {value: "]}", name: "]}"},
        {value: "\|", name: "\|"},
        {value: ";:", name: ";:"},
        {value: "'\"", name: "'\""},
        {value: "`~", name: "`~"},
        {value: ",<", name: ",<"},
        {value: ".>", name: ".>"},
        {value: "/?", name: "/?"},
        {value: "CpLk", name: "Caps Lock"},
        {value: "F1", name: "F1"},
        {value: "F2", name: "F2"},
        {value: "F3", name: "F3"},
        {value: "F4", name: "F4"},
        {value: "F5", name: "F5"},
        {value: "F6", name: "F6"},
        {value: "F7", name: "F7"},
        {value: "F8", name: "F8"},
        {value: "F9", name: "F9"},
        {value: "F10", name: "F10"},
        {value: "F11", name: "F11"},
        {value: "F12", name: "F12"},
        {value: "PrScn", name: "Print Screen"},
        {value: "ScLck", name: "Scroll Lock"},
        {value: "Pause", name: "Pause"},
        {value: "Ins", name: "Insert"},
        {value: "Home", name: "Home"},
        {value: "PgUp", name: "Page Up"},
        {value: "Del", name: "Delete"},
        {value: "End", name: "End"},
        {value: "PgDwn", name: "Page Down"},
        {value: "Right", name: "Right"},
        {value: "Left", name: "Left"},
        {value: "Down", name: "Down"},
        {value: "Up", name: "Up"},
        {value: "NmLck", name: "Num Lock"},
        {value: "Katakana", name: "カタカナ ひらがな"},
        {value: "￥|", name: "￥|"},
        {value: "Cnv", name: "変換"},
        {value: "NoCnv", name: "無変換"},
        {value: "M_LB", name: "Mouse Left Button"},
        {value: "M_RB", name: "Mouse Right Button"},
        {value: "M_WHL", name: "Mouse Wheel"},
        {value: "Reset", name: "Reset"},
        {value: "XF_CUT1", name: "XFader Cut Ch.1(Phono/Line In)"},
        {value: "XF_CUT2", name: "XFader Cut Ch.2(USB)"},
        {value: "MGain_Up", name: "Master Gain Up"},
        {value: "MGain_Down", name: "Master Gain Down"},
        {value: "Upper", name: "Upper"},
        {value: "LNPH", name: "Line Phono Switch"},
        {value: "L1/L2", name: "Layout Switch"},
    ]

    let keytopMap = new Map(allKeyMaps.map(item => [item.name, item.value]));

    let normalLayout = [
        [
            ['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'],
            ['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ';:'],
            ['Z', 'X', 'C', 'V', 'B', 'N', 'M', ',<', '.>', '/?'],
            ['', 'Mod.', 'Mod.', 'STK1', 'Mod.', 'STK2', 'Left', 'Down', 'Up', 'Right']
        ],
        [
            ['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'],
            ['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ';:'],
            ['Z', 'X', 'C', 'V', 'B', 'N', 'M', ',<', '.>', '/?'],
            ['', 'Mod.', 'Mod.', 'STK1', 'Mod.', 'STK2', 'Left', 'Down', 'Up', 'Right']
        ]
    ];

    let normalModifiers = [
        [
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b10000000, 0b01000000, 0b00000000, 0b00010000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000]
        ],
        [
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b10000000, 0b01000000, 0b00000000, 0b00010000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000]
        ]
    ];

    let upperLayout = [
        [
            ['1!', '2@', '3#', '4$', '5%', '6^', '7&', '8*', '9(', '0)'],
            ['`~', '\'"', '', '', '', '', '[{', ']}', '-_', ';:'],
            ['', '', '', '', '', '', '=+', ',<', '.>', '/?'],
            ['', 'LGUI', 'LALT', 'STK1', 'LCTRL', 'STK2', '', 'Master Gain Down', 'Master Gain Up', 'Reset']
        ],
        [
            ['1!', '2@', '3#', '4$', '5%', '6^', '7&', '8*', '9(', '0)'],
            ['`~', '\'"', '', '', '', '', '[{', ']}', '-_', ';:'],
            ['', '', '', '', '', '', '=+', ',<', '.>', '/?'],
            ['', 'LGUI', 'LALT', 'STK1', 'LCTRL', 'STK2', '', 'Master Gain Down', 'Master Gain Up', 'Reset']
        ],
    ];

    let upperModifiers = [
        [
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b10000000, 0b01000000, 0b00000000, 0b00010000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000]
        ],
        [
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
            [0b00000000, 0b10000000, 0b01000000, 0b00000000, 0b00010000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000]
        ]
    ];

    let stk1Keys = [
        ['Upper', 'Upper', ''],
        ['Backspace', '', 'Tab'],
        ['', 'Enter', ''],
    ];

    let stk1Modifiers = [
        [0b00100000, 0b00000000, 0b00000000],
        [0b00000000, 0b00000000, 0b00000000],
        [0b00000000, 0b00000000, 0b00000000],
    ];

    let stk2Keys = [
        ['', '', 'Upper'],
        ['Delete', '', 'Space'],
        ['', 'Escape', 'Space'],
    ];

    let stk2Modifiers = [
        [0b00000000, 0b00000010, 0b00000010],
        [0b00000000, 0b00000000, 0b00000000],
        [0b00000000, 0b00000000, 0b00000010],
    ];

    let selectedLayoutIndex = 0;
    let selectedNormal = true;

    let dropdownOpen = false;
    let selectedRow = 0;
    let selectedCol = 0;

    let stkDropdownOpen = false;
    let selectedDirectionRow = 0;
    let selectedDirectionCol = 0;
    let currentStkKeycode = "";
    let isDisabledStkDropdown = true;

    let isCheckedLGUI = false;
    let isCheckedLALT = false;
    let isCheckedLSHIFT = false;
    let isCheckedLCTRL = false;
    let isCheckedRGUI = false;
    let isCheckedRALT = false;
    let isCheckedRSHIFT = false;
    let isCheckedRCTRL = false;

    let isStkCheckedLGUI = false;
    let isStkCheckedLALT = false;
    let isStkCheckedLSHIFT = false;
    let isStkCheckedLCTRL = false;
    let isStkCheckedRGUI = false;
    let isStkCheckedRALT = false;
    let isStkCheckedRSHIFT = false;
    let isStkCheckedRCTRL = false;

    let clickOutsideImport = false;

    function selectLayout(event, index) {
        selectedLayoutIndex = index;
    }

    function selectNormal(event, flag) {
        selectedNormal = flag;
    }

    function selectKey(event) {
        let val = parseInt(event.target.value);
        selectedCol = val % 10;
        selectedRow = Math.floor(val / 10);
        //alert(`${event.target.value} val: ${val}, row: ${selectedRow}, col: ${selectedCol}`);

        if (selectedRow === 3 && (selectedCol === 3 || selectedCol === 5)) {
            selectedSettingTab = 1;
            directionGroup = "";

            isStkCheckedLGUI = false;
            isStkCheckedLALT = false;
            isStkCheckedLSHIFT = false;
            isStkCheckedLSHIFT = false;
            isStkCheckedRGUI = false;
            isStkCheckedRALT = false;
            isStkCheckedRSHIFT = false;
            isStkCheckedRSHIFT = false;
        } else {
            selectedSettingTab = 0;
        }

        if (selectedNormal) {
            isDisabledStkDropdown = true;
            currentStkKeycode = "";

            isCheckedRCTRL = !!(normalModifiers[0][selectedRow][selectedCol] & 0b00000001);
            isCheckedRSHIFT = !!((normalModifiers[0][selectedRow][selectedCol] >> 1) & 0b00000001);
            isCheckedRALT = !!((normalModifiers[0][selectedRow][selectedCol] >> 2) & 0b00000001);
            isCheckedRGUI = !!((normalModifiers[0][selectedRow][selectedCol] >> 3) & 0b00000001);
            isCheckedLCTRL = !!((normalModifiers[0][selectedRow][selectedCol] >> 4) & 0b00000001);
            isCheckedLSHIFT = !!((normalModifiers[0][selectedRow][selectedCol] >> 5) & 0b00000001);
            isCheckedLALT = !!((normalModifiers[0][selectedRow][selectedCol] >> 6) & 0b00000001);
            isCheckedLGUI = !!((normalModifiers[0][selectedRow][selectedCol] >> 7) & 0b00000001);
        } else {
            isDisabledStkDropdown = true;
            currentStkKeycode = "";

            isCheckedRCTRL = !!(upperModifiers[0][selectedRow][selectedCol] & 0b00000001);
            isCheckedRSHIFT = !!((upperModifiers[0][selectedRow][selectedCol] >> 1) & 0b00000001);
            isCheckedRALT = !!((upperModifiers[0][selectedRow][selectedCol] >> 2) & 0b00000001);
            isCheckedRGUI = !!((upperModifiers[0][selectedRow][selectedCol] >> 3) & 0b00000001);
            isCheckedLCTRL = !!((upperModifiers[0][selectedRow][selectedCol] >> 4) & 0b00000001);
            isCheckedLSHIFT = !!((upperModifiers[0][selectedRow][selectedCol] >> 5) & 0b00000001);
            isCheckedLALT = !!((upperModifiers[0][selectedRow][selectedCol] >> 6) & 0b00000001);
            isCheckedLGUI = !!((upperModifiers[0][selectedRow][selectedCol] >> 7) & 0b00000001);
        }
    }

    function selectDirection(event) {
        let val = parseInt(event.target.value);
        selectedDirectionCol = val % 3;
        selectedDirectionRow = Math.floor(val / 3);
        //alert(`${event.target.value} val: ${val}, row: ${selectedRow}, col: ${selectedCol}`);

        if (selectedRow === 3 && selectedCol === 3) {
            currentStkKeycode = stk1Keys[selectedDirectionRow][selectedDirectionCol];

            isStkCheckedRCTRL = !!(stk1Modifiers[selectedDirectionRow][selectedDirectionCol] & 0b00000001);
            isStkCheckedRSHIFT = !!((stk1Modifiers[selectedDirectionRow][selectedDirectionCol] >> 1) & 0b00000001);
            isStkCheckedRALT = !!((stk1Modifiers[selectedDirectionRow][selectedDirectionCol] >> 2) & 0b00000001);
            isStkCheckedRGUI = !!((stk1Modifiers[selectedDirectionRow][selectedDirectionCol] >> 3) & 0b00000001);
            isStkCheckedLCTRL = !!((stk1Modifiers[selectedDirectionRow][selectedDirectionCol] >> 4) & 0b00000001);
            isStkCheckedLSHIFT = !!((stk1Modifiers[selectedDirectionRow][selectedDirectionCol] >> 5) & 0b00000001);
            isStkCheckedLALT = !!((stk1Modifiers[selectedDirectionRow][selectedDirectionCol] >> 6) & 0b00000001);
            isStkCheckedLGUI = !!((stk1Modifiers[selectedDirectionRow][selectedDirectionCol] >> 7) & 0b00000001);
        } else if (selectedRow === 3 && selectedCol === 5) {
            currentStkKeycode = stk2Keys[selectedDirectionRow][selectedDirectionCol];

            isStkCheckedRCTRL = !!(stk2Modifiers[selectedDirectionRow][selectedDirectionCol] & 0b00000001);
            isStkCheckedRSHIFT = !!((stk2Modifiers[selectedDirectionRow][selectedDirectionCol] >> 1) & 0b00000001);
            isStkCheckedRALT = !!((stk2Modifiers[selectedDirectionRow][selectedDirectionCol] >> 2) & 0b00000001);
            isStkCheckedRGUI = !!((stk2Modifiers[selectedDirectionRow][selectedDirectionCol] >> 3) & 0b00000001);
            isStkCheckedLCTRL = !!((stk2Modifiers[selectedDirectionRow][selectedDirectionCol] >> 4) & 0b00000001);
            isStkCheckedLSHIFT = !!((stk2Modifiers[selectedDirectionRow][selectedDirectionCol] >> 5) & 0b00000001);
            isStkCheckedLALT = !!((stk2Modifiers[selectedDirectionRow][selectedDirectionCol] >> 6) & 0b00000001);
            isStkCheckedLGUI = !!((stk2Modifiers[selectedDirectionRow][selectedDirectionCol] >> 7) & 0b00000001);

        }

        isDisabledStkDropdown = false;
    }

    function renewKeytop(event) {
        if (selectedNormal) {
            normalLayout[0][selectedRow][selectedCol] = keytopMap.get(event.target.textContent);
        } else {
            upperLayout[0][selectedRow][selectedCol] = keytopMap.get(event.target.textContent);
        }

        dropdownOpen = false;
    }

    function renewModifiers(event, index) {
        if (selectedNormal) {
            if (event.target.checked) {
                normalModifiers[0][selectedRow][selectedCol] |= 0b00000001 << index;
            } else {
                normalModifiers[0][selectedRow][selectedCol] &= ~(0b00000001 << index);
            }
        } else {
            if (event.target.checked) {
                upperModifiers[0][selectedRow][selectedCol] |= 0b00000001 << index;
            } else {
                upperModifiers[0][selectedRow][selectedCol] &= ~(0b00000001 << index);
            }

        }
    }

    function renewStkKeycode(event) {
        if (selectedRow === 3 && selectedCol === 3) {
            stk1Keys[selectedDirectionRow][selectedDirectionCol] = event.target.textContent;
        } else if (selectedRow === 3 && selectedCol === 5) {
            stk2Keys[selectedDirectionRow][selectedDirectionCol] = event.target.textContent;
        }

        currentStkKeycode = event.target.textContent;

        stkDropdownOpen = false;
    }

    function renewStkModifiers(event, index) {
        if (selectedRow === 3 && selectedCol === 3) {
            if (event.target.checked) {
                stk1Modifiers[selectedDirectionRow][selectedDirectionCol] |= 0b00000001 << index;
            } else {
                stk1Modifiers[selectedDirectionRow][selectedDirectionCol] &= ~(0b00000001 << index);
            }
        } else if (selectedRow === 3 && selectedCol === 5) {
            if (event.target.checked) {
                stk2Modifiers[selectedDirectionRow][selectedDirectionCol] |= 0b00000001 << index;
            } else {
                stk2Modifiers[selectedDirectionRow][selectedDirectionCol] &= ~(0b00000001 << index);
            }

        }
    }

    function isDisabled(row, col) {
        return row === 3 && col === 0;
    }

    function isRounded(row, col) {
        return (row === 3 && col === 3) || (row === 3 && col === 5);
    }

    function isNokey(name) {
        if (name === '')
            return true;

        return false;
    }

    function getIcon(name) {
        if (name === 'Left')
            return AngleLeftOutline;

        if (name === 'Down')
            return AngleDownOutline;

        if (name === 'Up')
            return AngleUpOutline;

        if (name === 'Right')
            return AngleRightOutline;

        if (name === 'Master Gain Up')
            return VolumeUpOutline;

        if (name === 'Master Gain Down')
            return VolumeDownOutline;

        return null;
    }

    /*
    function isValidTOML(content) {
        try {
            parse(content);
            return true;
        } catch (error) {
            LogError('Invalid TOML data: ' + error.message);
            return false;
        }
    }
     */

    async function sendFileContent(content) {
        try {
            const response = await fetch('http://localhost:8080/upload', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/toml'
                },
                body: content
            });

            if (!response.ok) {
                throw new Error('Failed to send file content');
            }

            const result = await response.text();
            LogInfo('File content sent successfully: ' + result);

            for (let row = 0; row < 4; row++) {
                for (let col = 0; col < 10; col++) {
                    if ((row === 3 && col === 0) || (row === 3 && col === 3) || (row === 3 && col === 5)) {
                        continue;
                    }

                    GetNormalKeyOfLayout1(row, col).then((key) => {
                        GetNormalModifiersOfLayout1(row, col).then((mods) => {
                            LogInfo(`NORMAL: row: ${row}, col: ${col}, key: ${key}, mods: ${mods}`);
                            if (key === '&nbsp;' && mods > 0) {
                                normalLayout[0][row][col] = 'Mod.';
                            } else if (key === '&nbsp;' && mods === 0) {
                                normalLayout[0][row][col] = '';
                            } else {
                                normalLayout[0][row][col] = key;
                            }
                            normalModifiers[0][row][col] = mods;
                        });
                    });

                    GetUpperKeyOfLayout1(row, col).then((key) => {
                        GetUpperModifiersOfLayout1(row, col).then((mods) => {
                            LogInfo(`UPPER: row: ${row}, col: ${col}, key: ${key}, mods: ${mods}`);
                            if (key === '&nbsp;' && mods > 0) {
                                upperLayout[0][row][col] = 'Mod.';
                            } else if (key === '&nbsp;' && mods === 0) {
                                upperLayout[0][row][col] = '';
                            } else {
                                upperLayout[0][row][col] = key;
                            }
                            upperModifiers[0][row][col] = mods;
                        });
                    });

                    GetNormalKeyOfLayout2(row, col).then((key) => {
                        GetNormalModifiersOfLayout2(row, col).then((mods) => {
                            LogInfo(`NORMAL: row: ${row}, col: ${col}, key: ${key}, mods: ${mods}`);
                            if (key === '&nbsp;' && mods > 0) {
                                normalLayout[0][row][col] = 'Mod.';
                            } else if (key === '&nbsp;' && mods === 0) {
                                normalLayout[1][row][col] = '';
                            } else {
                                normalLayout[1][row][col] = key;
                            }
                            normalModifiers[1][row][col] = mods;
                        });
                    });

                    GetUpperKeyOfLayout2(row, col).then((key) => {
                        GetUpperModifiersOfLayout2(row, col).then((mods) => {
                            LogInfo(`UPPER: row: ${row}, col: ${col}, key: ${key}, mods: ${mods}`);
                            if (key === '&nbsp;' && mods > 0) {
                                upperLayout[1][row][col] = 'Mod.';
                            } else if (key === '&nbsp;' && mods === 0) {
                                upperLayout[1][row][col] = '';
                            } else {
                                upperLayout[1][row][col] = key;
                            }
                            upperModifiers[1][row][col] = mods;
                        });
                    });
                }
            }
        } catch (error) {
            LogError('Error sending file content: ' + error);
        }
    }

    let value = [];
    const dropHandle = (event) => {
        value = [];
        event.preventDefault();
        if (event.dataTransfer.items) {
            [...event.dataTransfer.items].forEach((item, i) => {
                if (item.kind === 'file') {
                    const file = item.getAsFile();
                    value.push(file.name);
                    value = value;

                    LogInfo("[dropHandle] File detected: " + file.name);

                    if (checkFile(file.name)) {
                        file.text().then((fileContent) => {
                            sendFileContent(fileContent);
                        });
                    }
                }
            });
        } else {
            [...event.dataTransfer.files].forEach((file, i) => {
                value = file.name;
            });
        }
    };

    const handleChange = (event) => {
        value = [];
        const files = event.target.files;

        if (files.length > 0) {
            value.push(files[0].name);
            value = value;

            LogInfo("[handleHandle] File detected: " + files[0].name);

            if (checkFile(files[0].name)) {
                files[0].text().then((fileContent) => {
                    sendFileContent(fileContent);
                });
            }
        }
    };

    const checkFile = (file) => {
        const fileExtension = file.split('.').pop();
        return fileExtension === 'toml';
    };

    const showFiles = (files) => {
        if (files.length > 0) {
            if (checkFile(files[0])) {
                return files[0];
            } else {
                return '';
            }
        }
    };

    function saveFile() {
        let data = {
            title: "Example TOML",
            owner: {
                name: "Tom Preston-Werner",
                dob: new Date('1979-05-27T07:32:00Z'),
            },
            database: {
                server: "192.168.1.1",
                ports: [8000, 8001, 8002],
                connection_max: 5000,
                enabled: true
            },
            servers: {
                alpha: {
                    ip: "10.0.0.1",
                    dc: "eqdc10"
                },
                beta: {
                    ip: "10.0.0.2",
                    dc: "eqdc10"
                }
            },
            clients: {
                data: [["gamma", "delta"], [1, 2]]
            }
        };
        const tomlStr = tomlify.toToml(data, {space: 2});

        const blob = new Blob([tomlStr], {type: "application/toml;charset=utf-8"});
        //const blob = new Blob(['hello world.'], {type: "application/toml;charset=utf-8"});
        const link = document.createElement("a");

        link.href = URL.createObjectURL(blob);
        link.download = "my_stk_layout.toml";
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    }
</script>

<main class="dark bg-gray-700">
    <div class="flex space-x-1 float-end">
        <div>
            <Button outline on:click={() => (clickOutsideImport = true)}>Import</Button>
        </div>
        <div>
            <Button outline on:click={saveFile}>Export</Button>
        </div>
    </div>

    <Modal title="Import TOML file" bind:open={clickOutsideImport} outsideclose>
        <Dropzone
                on:drop={dropHandle}
                on:dragover={(event) => {
                    event.preventDefault();
                }}
                on:change={handleChange}
                accept=".toml">
            <svg aria-hidden="true" class="mb-3 w-10 h-10 text-gray-400" fill="none" stroke="currentColor"
                 viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"/>
            </svg>
            {#if value.length === 0 || checkFile(value[0]) === false}
                <p class="mb-2 text-sm text-gray-500 dark:text-gray-400"><span
                        class="font-semibold">Click to upload</span> or drag and drop</p>
                <p class="text-xs text-gray-500 dark:text-gray-400">TOML only.</p>
            {:else}
                <p>{showFiles(value)}</p>
            {/if}
        </Dropzone>
    </Modal>

    <Tabs class="border-gray-700" tabStyle="underline">
        <TabItem open title="Layout1" on:click={(event) => selectLayout(event, 0)}>
            <Tabs class="border-gray-700" tabStyle="underline">
                <TabItem open title="Normal" on:click={(event) => selectNormal(event, true)}>
                    {#each normalLayout[0] as row, rowIndex}
                        {#each row as key, colIndex}
                            {#if key !== null}
                                <RadioButton
                                        value="{10 * rowIndex + colIndex}"
                                        bind:group={radioGroup}
                                        size="sm"
                                        class="w-12 h-12 mx-1 my-1 bg-primary-600 {isRounded(rowIndex, colIndex) ? 'rounded-full':''}"
                                        on:change={selectKey}
                                        disabled={isDisabled(rowIndex, colIndex)}>
                                    {#if key !== 'Left' && key !== 'Right' && key !== 'Up' && key !== 'Down' && key !== 'Master Gain Up' && key !== 'Master Gain Down'}
                                        {@html key}
                                    {/if}
                                    {#if isNokey(key)}
                                        &nbsp;
                                    {/if}
                                    {#if getIcon(key)}
                                        &nbsp;<svelte:component this={getIcon(key)}
                                                                class="w-5 h-6 -ms-1 text-white dark:text-white"/>
                                    {/if}
                                </RadioButton>
                            {/if}
                        {/each}
                        <br>
                    {/each}
                </TabItem>
                <TabItem title="Upper" on:click={(event) => selectNormal(event, false)}>
                    {#each upperLayout[0] as row, rowIndex}
                        {#each row as key, colIndex}
                            {#if key !== null}
                                <RadioButton
                                        value="{10 * rowIndex + colIndex}"
                                        bind:group={radioGroup}
                                        size="sm"
                                        class="w-12 h-12 mx-1 my-1 {isRounded(rowIndex, colIndex) ? 'rounded-full':''}"
                                        on:change={selectKey}
                                        disabled={isDisabled(rowIndex, colIndex)}>
                                    {key === 'Left' || key === 'Right' || key === 'Up' || key === 'Down' || key === 'Master Gain Up' || key === 'Master Gain Down' ? '' : key}
                                    {#if isNokey(key)}
                                        &nbsp;
                                    {/if}
                                    {#if getIcon(key)}
                                        &nbsp;<svelte:component this={getIcon(key)}
                                                                class="w-5 h-6 -ms-1"/>
                                    {/if}
                                </RadioButton>
                            {/if}
                        {/each}
                        <br>
                    {/each}
                </TabItem>
            </Tabs>
        </TabItem>
        <TabItem title="Layout2" on:click={(event) => selectLayout(event, 1)}>
            <Tabs class="border-gray-700" tabStyle="underline">
                <TabItem open title="Normal" on:click={(event) => selectNormal(event, true)}>
                    {#each normalLayout[1] as row, rowIndex}
                        {#each row as key, colIndex}
                            {#if key !== null}
                                <RadioButton
                                        value="{10 * rowIndex + colIndex}"
                                        bind:group={radioGroup}
                                        size="sm"
                                        class="w-12 h-12 mx-1 my-1 bg-primary-600 {isRounded(rowIndex, colIndex) ? 'rounded-full':''}"
                                        on:change={selectKey}
                                        disabled={isDisabled(rowIndex, colIndex)}>
                                    {#if key !== 'Left' && key !== 'Right' && key !== 'Up' && key !== 'Down' && key !== 'Master Gain Up' && key !== 'Master Gain Down'}
                                        {@html key}
                                    {/if}
                                    {#if isNokey(key)}
                                        &nbsp;
                                    {/if}
                                    {#if getIcon(key)}
                                        &nbsp;<svelte:component this={getIcon(key)}
                                                                class="w-5 h-6 -ms-1 text-white dark:text-white"/>
                                    {/if}
                                </RadioButton>
                            {/if}
                        {/each}
                        <br>
                    {/each}
                </TabItem>
                <TabItem title="Upper" on:click={(event) => selectNormal(event, false)}>
                    {#each upperLayout[1] as row, rowIndex}
                        {#each row as key, colIndex}
                            {#if key !== null}
                                <RadioButton
                                        value="{10 * rowIndex + colIndex}"
                                        bind:group={radioGroup}
                                        size="sm"
                                        class="w-12 h-12 mx-1 my-1 {isRounded(rowIndex, colIndex) ? 'rounded-full':''}"
                                        on:change={selectKey}
                                        disabled={isDisabled(rowIndex, colIndex)}>
                                    {key === 'Left' || key === 'Right' || key === 'Up' || key === 'Down' || key === 'Master Gain Up' || key === 'Master Gain Down' ? '' : key}
                                    {#if isNokey(key)}
                                        &nbsp;
                                    {/if}
                                    {#if getIcon(key)}
                                        &nbsp;<svelte:component this={getIcon(key)}
                                                                class="w-5 h-6 -ms-1"/>
                                    {/if}
                                </RadioButton>
                            {/if}
                        {/each}
                        <br>
                    {/each}
                </TabItem>
            </Tabs>

        </TabItem>
    </Tabs>
    <br>
    <hr>

    <br>
    <Tabs tabStyle="underline">
        <TabItem open={selectedSettingTab === 0} title="Key" disabled>
            <Button>Select keycode
                <ChevronRightOutline class="w-6 h-6 ms-2 text-white dark:text-white"/>
            </Button>
            <Dropdown class="overflow-y-auto py-1 h-48" placement="right" bind:open={dropdownOpen}>
                {#each allKeyMaps as keymap}
                    <DropdownItem on:click={renewKeytop}>{@html keymap.name}</DropdownItem>
                {/each}
            </Dropdown>
            <br>
            <br>
            <Table shadow>
                <TableHead>
                    <TableHeadCell colspan={4}><Label class="">Left</Label>
                    </TableHeadCell>
                    <TableHeadCell colspan={4}><Label class="">Right</Label>
                    </TableHeadCell>
                </TableHead>
                <TableBody>
                    <TableBodyRow>
                        <TableBodyCell>
                            <Checkbox class="-mr-10" bind:checked={isCheckedLGUI}
                                      on:change={(event) => renewModifiers(event, 7)}>GUI
                            </Checkbox>
                        </TableBodyCell>
                        <TableBodyCell>
                            <Checkbox class="-mr-10" bind:checked={isCheckedLALT}
                                      on:change={(event) => renewModifiers(event, 6)}>ALT
                            </Checkbox>
                        </TableBodyCell>
                        <TableBodyCell>
                            <Checkbox class="-mr-10" bind:checked={isCheckedLSHIFT}
                                      on:change={(event) => renewModifiers(event, 5)}>SHIFT
                            </Checkbox>
                        </TableBodyCell>
                        <TableBodyCell>
                            <Checkbox class="-mr-10" bind:checked={isCheckedLCTRL}
                                      on:change={(event) => renewModifiers(event, 4)}>CTRL
                            </Checkbox>
                        </TableBodyCell>
                        <TableBodyCell>
                            <Checkbox class="-mr-10" bind:checked={isCheckedRGUI}
                                      on:change={(event) => renewModifiers(event, 3)}>GUI
                            </Checkbox>
                        </TableBodyCell>
                        <TableBodyCell>
                            <Checkbox class="-mr-10" bind:checked={isCheckedRALT}
                                      on:change={(event) => renewModifiers(event, 2)}>ALT
                            </Checkbox>
                        </TableBodyCell>
                        <TableBodyCell>
                            <Checkbox class="-mr-10" bind:checked={isCheckedRSHIFT}
                                      on:change={(event) => renewModifiers(event, 1)}>SHIFT
                            </Checkbox>
                        </TableBodyCell>
                        <TableBodyCell>
                            <Checkbox class="-mr-15" bind:checked={isCheckedRCTRL}
                                      on:change={(event) => renewModifiers(event, 0)}>CTRL
                            </Checkbox>
                        </TableBodyCell>
                    </TableBodyRow>
                </TableBody>
            </Table>

        </TabItem>
        <TabItem open={selectedSettingTab === 1} title="Stick" disabled>
            <div class="flex space-x-10">
                <div>
                    <RadioButton value="0" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                 on:change={selectDirection}>
                        <ArrowUpOutline class="-rotate-45"/>
                    </RadioButton>
                    <RadioButton value="1" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                 on:change={selectDirection}>
                        <ArrowUpOutline/>
                    </RadioButton>
                    <RadioButton value="2" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                 on:change={selectDirection}>
                        <ArrowUpOutline class="rotate-45"/>
                    </RadioButton>
                    <br>
                    <RadioButton value="3" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                 on:change={selectDirection}>
                        <ArrowLeftOutline/>
                    </RadioButton>
                    <RadioButton value="4" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                 disabled>
                        <CogOutline/>
                    </RadioButton>
                    <RadioButton value="5" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                 on:change={selectDirection}>
                        <ArrowRightOutline/>
                    </RadioButton>
                    <br>
                    <RadioButton value="6" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                 on:change={selectDirection}>
                        <ArrowDownOutline class="rotate-45"/>
                    </RadioButton>
                    <RadioButton value="7" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                 on:change={selectDirection}>
                        <ArrowDownOutline/>
                    </RadioButton>
                    <RadioButton value="8" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                 on:change={selectDirection}>
                        <ArrowDownOutline class="-rotate-45"/>
                    </RadioButton>
                </div>
                <div>
                    <div class="flex space-x-10">
                        <div>
                            <Button disabled={isDisabledStkDropdown}>Select keycode
                                <ChevronRightOutline class="w-6 h-6 ms-2 text-white dark:text-white"/>
                            </Button>
                            <Dropdown class="overflow-y-auto py-1 h-48" placement="right" bind:open={stkDropdownOpen}>
                                {#each allKeyMaps as keymap}
                                    <DropdownItem on:click={renewStkKeycode}>{@html keymap.name}</DropdownItem>
                                {/each}
                            </Dropdown>
                        </div>
                        <div>
                            <Label class="mt-3">
                                Current Keycode : {currentStkKeycode}
                            </Label>
                        </div>
                    </div>
                    <br>
                    <Table shadow>
                        <TableBody>
                            <TableBodyRow>
                                <TableBodyCell>
                                    <Label>L:</Label>
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Checkbox bind:checked={isStkCheckedLGUI}
                                              on:change={(event) => renewStkModifiers(event, 7)}>GUI
                                    </Checkbox>
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Checkbox bind:checked={isStkCheckedLALT}
                                              on:change={(event) => renewStkModifiers(event, 6)}>ALT
                                    </Checkbox>
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Checkbox bind:checked={isStkCheckedLSHIFT}
                                              on:change={(event) => renewStkModifiers(event, 5)}>SHIFT
                                    </Checkbox>
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Checkbox bind:checked={isStkCheckedLCTRL}
                                              on:change={(event) => renewStkModifiers(event, 4)}>CTRL
                                    </Checkbox>
                                </TableBodyCell>
                            </TableBodyRow>
                            <TableBodyRow>
                                <TableBodyCell>
                                    <Label>R:</Label>
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Checkbox bind:checked={isStkCheckedRGUI}
                                              on:change={(event) => renewStkModifiers(event, 3)}>GUI
                                    </Checkbox>
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Checkbox bind:checked={isStkCheckedRALT}
                                              on:change={(event) => renewStkModifiers(event, 2)}>ALT
                                    </Checkbox>
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Checkbox bind:checked={isStkCheckedRSHIFT}
                                              on:change={(event) => renewStkModifiers(event, 1)}>SHIFT
                                    </Checkbox>
                                </TableBodyCell>
                                <TableBodyCell>
                                    <Checkbox bind:checked={isStkCheckedRCTRL}
                                              on:change={(event) => renewStkModifiers(event, 0)}>CTRL
                                    </Checkbox>
                                </TableBodyCell>
                            </TableBodyRow>
                        </TableBody>
                    </Table>
                </div>
            </div>
        </TabItem>
    </Tabs>
</main>

<style>
    :root {
        font-family: 'Nunito', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
        Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
        text-align: center;
    }

    main {
        text-align: center;
        padding: 1em;
        margin: 0 auto;
        height: 100vh;
    }

    @font-face {
        font-family: "Nunito";
        font-style: normal;
        font-weight: 400;
        src: local(""),
        url("assets/fonts/nunito-v16-latin-regular.woff2") format("woff2");
    }
</style>
