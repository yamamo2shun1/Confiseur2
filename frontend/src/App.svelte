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
        TableHead,
        TableHeadCell,
        Tabs
    } from 'flowbite-svelte';
    import {
        AngleDownOutline,
        AngleLeftOutline,
        AngleRightOutline,
        AngleUpOutline,
        ChevronRightOutline
    } from "flowbite-svelte-icons";

    let radioGroup = "";
    let keycode = "";

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
        {value: "Minus", name: "-_"},
        {value: "Equal", name: "=+"},
        {value: "O_SBracket", name: "[{"},
        {value: "C_SBracket", name: "]}"},
        {value: "Backslash", name: "\|"},
        {value: "Colon", name: ";:"},
        {value: "Apostrophe", name: "'\""},
        {value: "Backquote", name: "`~"},
        {value: "Comma", name: ",<"},
        {value: "Period", name: ".>"},
        {value: "Slash", name: "/?"},
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
        {value: "PrintScreen", name: "Print Screen"},
        {value: "ScrollLock", name: "Scroll Lock"},
        {value: "Pause", name: "Pause"},
        {value: "Ins", name: "Insert"},
        {value: "Home", name: "Home"},
        {value: "PageUp", name: "Page Up"},
        {value: "Del", name: "Delete"},
        {value: "End", name: "End"},
        {value: "PageDown", name: "Page Down"},
        {value: "Right", name: "Right"},
        {value: "Left", name: "Left"},
        {value: "Down", name: "Down"},
        {value: "Up", name: "Up"},
        {value: "NumLock", name: "Num Lock"},
        {value: "Katakana", name: "カタカナ ひらがな"},
        {value: "Yen", name: "￥|"},
        {value: "Henkan", name: "変換"},
        {value: "Muhenkan", name: "無変換"},
        {value: "M_LBTN", name: "Mouse Left Button"},
        {value: "M_RBTN", name: "Mouse Right Button"},
        {value: "M_WHEEL", name: "Mouse Wheel"},
        {value: "Reset", name: "Reset"},
        {value: "XF_CUT1", name: "XFader Cut Ch.1(Phono/Line In)"},
        {value: "XF_CUT2", name: "XFader Cut Ch.2(USB)"},
        {value: "MGain_Up", name: "Master Gain Up"},
        {value: "MGain_Down", name: "Master Gain Down"},
        {value: "Upper", name: "Upper"},
        {value: "LNPH", name: "Line Phono Switch"},
        {value: "LAYOUT", name: "Layout Switch"},
    ]

    let normalLayout1 = [
        ['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'],
        ['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ';'],
        ['Z', 'X', 'C', 'V', 'B', 'N', 'M', ',', '.', '/'],
        ['', 'GUI', 'ALT', 'STK1', 'CTRL', 'STK2', 'Left', 'Down', 'Up', 'Right']
    ];

    let upperLayout1 = [
        ['1', '2', '3', '4', '5', '6', '7', '8', '9', '0'],
        ['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ';'],
        ['Z', 'X', 'C', 'V', 'B', 'N', 'M', ',', '.', '/'],
        ['', 'GUI', 'ALT', 'STK1', 'CTRL', 'STK2', 'Left', 'Down', 'Up', 'Right']
    ];

    let selectedLayoutIndex = 0;

    let dropdownOpen = false;
    let selectedRow = 0;
    let selectedCol = 0;

    function selectLayout(event, index) {
        selectedLayoutIndex = index;
    }

    function selectKey() {
        let val = parseInt(event.target.value);
        selectedCol = val % 10;
        selectedRow = Math.floor(val / 10);
        //alert(`${event.target.value} val: ${val}, row: ${selectedRow}, col: ${selectedCol}`);
    }

    function renewKeycode() {
        if (selectedLayoutIndex === 0) {
            normalLayout1[selectedRow][selectedCol] = event.target.textContent;
        } else if (selectedLayoutIndex === 1) {
            upperLayout1[selectedRow][selectedCol] = event.target.textContent;
        }

        dropdownOpen = false;
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

        return null;
    }
</script>

<main>
    <Tabs>
        <TabItem open title="Normal" on:click={(event) => selectLayout(event, 0)}>
            <div>
                {#each normalLayout1 as row, rowIndex}
                    {#each row as key, colIndex}
                        {#if key !== null}
                            <RadioButton
                                    value="{10 * rowIndex + colIndex}"
                                    bind:group={radioGroup}
                                    class="w-12 h-12 mx-1 my-1 {isRounded(rowIndex, colIndex) ? 'rounded-full':''}"
                                    on:change={selectKey}
                                    disabled={isDisabled(rowIndex, colIndex)}>
                                {key === 'Left' || key === 'Right' || key === 'Up' || key === 'Down' ? '' : key}
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
            </div>
        </TabItem>
        <TabItem title="Upper" on:click={(event) => selectLayout(event, 1)}>
            <div>
                {#each upperLayout1 as row, rowIndex}
                    {#each row as key, colIndex}
                        {#if key !== null}
                            <RadioButton
                                    value="{10 * rowIndex + colIndex}"
                                    bind:group={radioGroup}
                                    class="w-12 h-12 mx-1 my-1 {isRounded(rowIndex, colIndex) ? 'rounded-full':''}"
                                    on:change={selectKey}
                                    disabled={isDisabled(rowIndex, colIndex)}>
                                {key === 'Left' || key === 'Right' || key === 'Up' || key === 'Down' ? '' : key}
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
            </div>
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
            <Table color="indigo">
                <TableHead>
                    <TableHeadCell colspan="4"><Label class="text-orange-500 dark:text-orange-500">Left</Label>
                    </TableHeadCell>
                    <TableHeadCell colspan="4"><Label class="text-orange-500 dark:text-orange-500">Right</Label>
                    </TableHeadCell>
                </TableHead>
                <TableBody>
                    <TableBodyCell>
                        <Checkbox class="-mr-10 text-orange-500 dark:text-orange-500">GUI</Checkbox>
                    </TableBodyCell>
                    <TableBodyCell>
                        <Checkbox class="-mr-10 text-orange-500 dark:text-orange-500">ALT</Checkbox>
                    </TableBodyCell>
                    <TableBodyCell>
                        <Checkbox class="-mr-10 text-orange-500 dark:text-orange-500">SHIFT</Checkbox>
                    </TableBodyCell>
                    <TableBodyCell>
                        <Checkbox class="-mr-10 text-orange-500 dark:text-orange-500">CTRL</Checkbox>
                    </TableBodyCell>
                    <TableBodyCell>
                        <Checkbox class="-mr-10 text-orange-500 dark:text-orange-500">GUI</Checkbox>
                    </TableBodyCell>
                    <TableBodyCell>
                        <Checkbox class="-mr-10 text-orange-500 dark:text-orange-500">ALT</Checkbox>
                    </TableBodyCell>
                    <TableBodyCell>
                        <Checkbox class="-mr-10 text-orange-500 dark:text-orange-500">SHIFT</Checkbox>
                    </TableBodyCell>
                    <TableBodyCell>
                        <Checkbox class="-mr-15 text-orange-500 dark:text-orange-500">CTRL</Checkbox>
                    </TableBodyCell>
                </TableBody>
            </Table>

        </TabItem>
        <TabItem title="Stick">

        </TabItem>
    </Tabs>
</main>

<style>
    :root {
        font-family: 'Nunito', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
        Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
        background-color: rgba(27, 38, 54, 1);
        text-align: center;
        color: white;
    }

    main {
        text-align: center;
        padding: 1em;
        margin: 0 auto;
        height: 100vh;
    }

    p {
        max-width: 14rem;
        margin: 1rem auto;
        line-height: 1.35;
    }


    @font-face {
        font-family: "Nunito";
        font-style: normal;
        font-weight: 400;
        src: local(""),
        url("assets/fonts/nunito-v16-latin-regular.woff2") format("woff2");
    }
</style>
