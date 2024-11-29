<script>
    import {
        Button,
        Checkbox,
        Dropdown,
        DropdownItem,
        Label,
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

    let radioGroup = "";
    let directionGroup = "";

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

    let normalLayout1 = [
        ['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'],
        ['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ';:'],
        ['Z', 'X', 'C', 'V', 'B', 'N', 'M', ',<', '.>', '/?'],
        ['', 'Mod.', 'Mod.', 'STK1', 'Mod.', 'STK2', 'Left', 'Down', 'Up', 'Right']
    ];

    let normalModifiers1 = [
        [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
        [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
        [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
        [0b00000000, 0b10000000, 0b01000000, 0b00000000, 0b00010000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000]
    ];

    let upperLayout1 = [
        ['1!', '2@', '3#', '4$', '5%', '6^', '7&', '8*', '9(', '0)'],
        ['`~', '\'"', '', '', '', '', '[{', ']}', '-_', ';:'],
        ['', '', '', '', '', '', '=+', ',<', '.>', '/?'],
        ['', 'LGUI', 'LALT', 'STK1', 'LCTRL', 'STK2', '', 'Master Gain Down', 'Master Gain Up', 'Reset']
    ];

    let upperModifiers1 = [
        [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
        [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
        [0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000],
        [0b00000000, 0b10000000, 0b01000000, 0b00000000, 0b00010000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000]
    ];

    let selectedLayoutIndex = 0;

    let dropdownOpen = false;
    let selectedRow = 0;
    let selectedCol = 0;

    let isCheckedLGUI = false;
    let isCheckedLALT = false;
    let isCheckedLSHIFT = false;
    let isCheckedLCTRL = false;
    let isCheckedRGUI = false;
    let isCheckedRALT = false;
    let isCheckedRSHIFT = false;
    let isCheckedRCTRL = false;

    function selectLayout(event, index) {
        selectedLayoutIndex = index;
    }

    function selectKey(event) {
        let val = parseInt(event.target.value);
        selectedCol = val % 10;
        selectedRow = Math.floor(val / 10);
        //alert(`${event.target.value} val: ${val}, row: ${selectedRow}, col: ${selectedCol}`);

        switch (selectedLayoutIndex) {
            case 0:
                isCheckedRCTRL = !!(normalModifiers1[selectedRow][selectedCol] & 0b00000001);
                isCheckedRSHIFT = !!((normalModifiers1[selectedRow][selectedCol] >> 1) & 0b00000001);
                isCheckedRALT = !!((normalModifiers1[selectedRow][selectedCol] >> 2) & 0b00000001);
                isCheckedRGUI = !!((normalModifiers1[selectedRow][selectedCol] >> 3) & 0b00000001);
                isCheckedLCTRL = !!((normalModifiers1[selectedRow][selectedCol] >> 4) & 0b00000001);
                isCheckedLSHIFT = !!((normalModifiers1[selectedRow][selectedCol] >> 5) & 0b00000001);
                isCheckedLALT = !!((normalModifiers1[selectedRow][selectedCol] >> 6) & 0b00000001);
                isCheckedLGUI = !!((normalModifiers1[selectedRow][selectedCol] >> 7) & 0b00000001);
                break;
            case 1:
                alert(upperLayout1[selectedRow][selectedCol]);
                break;
        }
    }

    function renewKeycode(event) {
        if (selectedLayoutIndex === 0) {
            normalLayout1[selectedRow][selectedCol] = event.target.textContent;
        } else if (selectedLayoutIndex === 1) {
            upperLayout1[selectedRow][selectedCol] = event.target.textContent;
        }

        dropdownOpen = false;
    }

    function renewModifiers(event, index) {
        if (selectedLayoutIndex === 0) {
            if (event.target.checked) {
                normalModifiers1[selectedRow][selectedCol] |= 0b00000001 << index;
            } else {
                normalModifiers1[selectedRow][selectedCol] &= ~(0b00000001 << index);
            }
        } else if (selectedLayoutIndex === 1) {
            if (event.target.checked) {
                upperModifiers1[selectedRow][selectedCol] |= 0b00000001 << index;
            } else {
                upperModifiers1[selectedRow][selectedCol] &= ~(0b00000001 << index);
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
</script>

<main class="dark bg-gray-700">
    <Tabs class="border-gray-700">
        <TabItem open title="Normal" on:click={(event) => selectLayout(event, 0)}>
            {#each normalLayout1 as row, rowIndex}
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
        <TabItem title="Upper" on:click={(event) => selectLayout(event, 1)}>
            {#each upperLayout1 as row, rowIndex}
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
    <br>
    <hr>

    <br>
    <Tabs>
        <TabItem open title="Key">
            <Button>Select keycode
                <ChevronRightOutline class="w-6 h-6 ms-2 text-white dark:text-white"/>
            </Button>
            <Dropdown class="overflow-y-auto py-1 h-48" placement="right" bind:open={dropdownOpen}>
                {#each allKeyMaps as keymap}
                    <DropdownItem on:click={renewKeycode}>{@html keymap.name}</DropdownItem>
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
                            <Checkbox class="-mr-10" bind:checked={isCheckedRSHIFT}
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
        <TabItem title="Stick">
            <Table>
                <TableBodyRow>
                    <TableBodyCell>
                        <RadioButton value="0" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                     on:change={selectKey}>
                            <ArrowUpOutline class="-rotate-45"/>
                        </RadioButton>
                        <RadioButton value="1" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                     on:change={selectKey}>
                            <ArrowUpOutline/>
                        </RadioButton>
                        <RadioButton value="2" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                     on:change={selectKey}>
                            <ArrowUpOutline class="rotate-45"/>
                        </RadioButton>
                        <br>
                        <RadioButton value="3" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                     on:change={selectKey}>
                            <ArrowLeftOutline/>
                        </RadioButton>
                        <RadioButton value="4" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                     disabled>
                            <CogOutline/>
                        </RadioButton>
                        <RadioButton value="5" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                     on:change={selectKey}>
                            <ArrowRightOutline/>
                        </RadioButton>
                        <br>
                        <RadioButton value="6" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                     on:change={selectKey}>
                            <ArrowDownOutline class="rotate-45"/>
                        </RadioButton>
                        <RadioButton value="7" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                     on:change={selectKey}>
                            <ArrowDownOutline/>
                        </RadioButton>
                        <RadioButton value="8" bind:group={directionGroup} size="sm" class="w-12 h-12 mx-1 my-1"
                                     on:change={selectKey}>
                            <ArrowDownOutline class="-rotate-45"/>
                        </RadioButton>
                    </TableBodyCell>
                    <TableBodyCell>
                        <Button>Select keycode
                            <ChevronRightOutline class="w-6 h-6 ms-2 text-white dark:text-white"/>
                        </Button>
                        <Dropdown class="overflow-y-auto py-1 h-48" placement="right" bind:open={dropdownOpen}>
                            {#each allKeyMaps as keymap}
                                <DropdownItem on:click={renewKeycode}>{@html keymap.name}</DropdownItem>
                            {/each}
                        </Dropdown>
                        <br>
                        <Table shadow>
                            <TableBody>
                                <TableBodyRow>
                                    <TableBodyCell>
                                        <Label>L:</Label>
                                    </TableBodyCell>
                                    <TableBodyCell>
                                        <Checkbox bind:checked={isCheckedLGUI}
                                                  on:change={(event) => renewModifiers(event, 7)}>GUI
                                        </Checkbox>
                                    </TableBodyCell>
                                    <TableBodyCell>
                                        <Checkbox bind:checked={isCheckedLALT}
                                                  on:change={(event) => renewModifiers(event, 6)}>ALT
                                        </Checkbox>
                                    </TableBodyCell>
                                    <TableBodyCell>
                                        <Checkbox bind:checked={isCheckedRSHIFT}
                                                  on:change={(event) => renewModifiers(event, 5)}>SHIFT
                                        </Checkbox>
                                    </TableBodyCell>
                                    <TableBodyCell>
                                        <Checkbox bind:checked={isCheckedLCTRL}
                                                  on:change={(event) => renewModifiers(event, 4)}>CTRL
                                        </Checkbox>
                                    </TableBodyCell>
                                </TableBodyRow>
                                <TableBodyRow>
                                    <TableBodyCell>
                                        <Label>R:</Label>
                                    </TableBodyCell>
                                    <TableBodyCell>
                                        <Checkbox bind:checked={isCheckedRGUI}
                                                  on:change={(event) => renewModifiers(event, 3)}>GUI
                                        </Checkbox>
                                    </TableBodyCell>
                                    <TableBodyCell>
                                        <Checkbox bind:checked={isCheckedRALT}
                                                  on:change={(event) => renewModifiers(event, 2)}>ALT
                                        </Checkbox>
                                    </TableBodyCell>
                                    <TableBodyCell>
                                        <Checkbox bind:checked={isCheckedRSHIFT}
                                                  on:change={(event) => renewModifiers(event, 1)}>SHIFT
                                        </Checkbox>
                                    </TableBodyCell>
                                    <TableBodyCell>
                                        <Checkbox bind:checked={isCheckedRCTRL}
                                                  on:change={(event) => renewModifiers(event, 0)}>CTRL
                                        </Checkbox>
                                    </TableBodyCell>
                                </TableBodyRow>
                            </TableBody>
                        </Table>
                    </TableBodyCell>
                </TableBodyRow>
            </Table>
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
