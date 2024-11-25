<script>
    import {
        Button,
        Checkbox,
        Dropdown,
        DropdownDivider,
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

    let keys = [
        ['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'],
        ['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ';'],
        ['Z', 'X', 'C', 'V', 'B', 'N', 'M', ',', '.', '/'],
        ['', 'GUI', 'ALT', '', 'CTRL', '', 'LEFT', 'DOWN', 'UP', 'RIGHT']
    ];

    let dropdownOpen = false;
    let selectedRow = 0;
    let selectedCol = 0;

    function selectKey() {
        let val = parseInt(event.target.value);
        selectedCol = val % 10;
        selectedRow = Math.round(val / 10);
        //alert(`row: ${selectedRow}, col: ${selectedCol}`);
    }

    function renewKeycode() {
        keys[selectedRow][selectedCol] = event.target.textContent;
        //alert(`select: ${keys[selectedRow][selectedCol]}`);

        dropdownOpen = false;
    }
</script>

<main>
    <Tabs>
        <TabItem open title="Normal">
            <div>
                <RadioButton value="00" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[0][0]}</RadioButton>
                <RadioButton value="01" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[0][1]}</RadioButton>
                <RadioButton value="02" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[0][2]}</RadioButton>
                <RadioButton value="03" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[0][3]}</RadioButton>
                <RadioButton value="04" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[0][4]}</RadioButton>
                <RadioButton value="05" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[0][5]}</RadioButton>
                <RadioButton value="06" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[0][6]}</RadioButton>
                <RadioButton value="07" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[0][7]}</RadioButton>
                <RadioButton value="08" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[0][8]}</RadioButton>
                <RadioButton value="09" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[0][9]}</RadioButton>
                <br>
                <RadioButton value="10" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[1][0]}</RadioButton>
                <RadioButton value="11" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[1][1]}</RadioButton>
                <RadioButton value="12" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[1][2]}</RadioButton>
                <RadioButton value="13" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[1][3]}</RadioButton>
                <RadioButton value="14" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[1][4]}</RadioButton>
                <RadioButton value="15" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[1][5]}</RadioButton>
                <RadioButton value="16" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[1][6]}</RadioButton>
                <RadioButton value="17" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[1][7]}</RadioButton>
                <RadioButton value="18" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[1][8]}</RadioButton>
                <RadioButton value="19" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[1][9]}</RadioButton>
                <br>
                <RadioButton value="20" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[2][0]}</RadioButton>
                <RadioButton value="21" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[2][1]}</RadioButton>
                <RadioButton value="22" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[2][2]}</RadioButton>
                <RadioButton value="23" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[2][3]}</RadioButton>
                <RadioButton value="24" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[2][4]}</RadioButton>
                <RadioButton value="25" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[2][5]}</RadioButton>
                <RadioButton value="26" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[2][6]}</RadioButton>
                <RadioButton value="27" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[2][7]}</RadioButton>
                <RadioButton value="28" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[2][8]}</RadioButton>
                <RadioButton value="29" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[2][9]}</RadioButton>
                <br>
                <RadioButton value="30" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1" disabled>&nbsp;
                </RadioButton>
                <RadioButton value="31" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[3][1]}</RadioButton>
                <RadioButton value="32" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[3][2]}</RadioButton>
                <RadioButton value="33" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1 rounded-full">STK1
                </RadioButton>
                <RadioButton value="34" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1"
                             on:change={selectKey}>{keys[3][4]}</RadioButton>
                <RadioButton value="35" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1 rounded-full">STK2
                </RadioButton>
                <RadioButton value="36" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1" on:change={selectKey}>&nbsp;<AngleLeftOutline
                        class="w-5 h-6 -ms-1 text-white dark:text-white"/>
                </RadioButton>
                <RadioButton value="37" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1" on:change={selectKey}>&nbsp;<AngleDownOutline
                        class="w-5 h-6 -ms-1 text-white dark:text-white"/>
                </RadioButton>
                <RadioButton value="38" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1" on:change={selectKey}>&nbsp;<AngleUpOutline
                        class="w-5 h-6 -ms-1 text-white dark:text-white"/>
                </RadioButton>
                <RadioButton value="39" bind:group={radioGroup} class="w-12 h-12 mx-1 my-1" on:change={selectKey}>&nbsp;<AngleRightOutline
                        class="w-5 h-6 -ms-1 text-white dark:text-white"/>
                </RadioButton>
            </div>
        </TabItem>
        <TabItem title="Upper">
            <p class="text-sm text-gray-500 dark:text-gray-400">
                <button type="button"
                        class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">
                    Default
                </button>
                <button type="button"
                        class="py-2.5 px-5 me-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700">
                    Alternative
                </button>
                <button type="button"
                        class="text-white bg-gray-800 hover:bg-gray-900 focus:outline-none focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-700 dark:border-gray-700">
                    Dark
                </button>
                <button type="button"
                        class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
                    Light
                </button>
                <button type="button"
                        class="focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">
                    Green
                </button>
                <button type="button"
                        class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900">
                    Red
                </button>
                <button type="button"
                        class="focus:outline-none text-white bg-yellow-400 hover:bg-yellow-500 focus:ring-4 focus:ring-yellow-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:focus:ring-yellow-900">
                    Yellow
                </button>
                <button type="button"
                        class="focus:outline-none text-white bg-purple-700 hover:bg-purple-800 focus:ring-4 focus:ring-purple-300 font-medium rounded-lg text-sm px-5 py-2.5 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900">
                    Purple
                </button>
            </p>
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
                <DropdownItem on:click={renewKeycode}>&nbsp;</DropdownItem>
                <DropdownItem on:click={renewKeycode}>A</DropdownItem>
                <DropdownItem on:click={renewKeycode}>B</DropdownItem>
                <DropdownItem on:click={renewKeycode}>C</DropdownItem>
                <DropdownItem on:click={renewKeycode}>D</DropdownItem>
                <DropdownItem on:click={renewKeycode}>E</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F</DropdownItem>
                <DropdownItem on:click={renewKeycode}>G</DropdownItem>
                <DropdownItem on:click={renewKeycode}>H</DropdownItem>
                <DropdownItem on:click={renewKeycode}>I</DropdownItem>
                <DropdownItem on:click={renewKeycode}>J</DropdownItem>
                <DropdownItem on:click={renewKeycode}>K</DropdownItem>
                <DropdownItem on:click={renewKeycode}>L</DropdownItem>
                <DropdownItem on:click={renewKeycode}>M</DropdownItem>
                <DropdownItem on:click={renewKeycode}>N</DropdownItem>
                <DropdownItem on:click={renewKeycode}>O</DropdownItem>
                <DropdownItem on:click={renewKeycode}>P</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Q</DropdownItem>
                <DropdownItem on:click={renewKeycode}>R</DropdownItem>
                <DropdownItem on:click={renewKeycode}>S</DropdownItem>
                <DropdownItem on:click={renewKeycode}>T</DropdownItem>
                <DropdownItem on:click={renewKeycode}>U</DropdownItem>
                <DropdownItem on:click={renewKeycode}>V</DropdownItem>
                <DropdownItem on:click={renewKeycode}>W</DropdownItem>
                <DropdownItem on:click={renewKeycode}>X</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Y</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Z</DropdownItem>
                <DropdownDivider/>
                <DropdownItem on:click={renewKeycode}>0</DropdownItem>
                <DropdownItem on:click={renewKeycode}>1</DropdownItem>
                <DropdownItem on:click={renewKeycode}>2</DropdownItem>
                <DropdownItem on:click={renewKeycode}>3</DropdownItem>
                <DropdownItem on:click={renewKeycode}>4</DropdownItem>
                <DropdownItem on:click={renewKeycode}>5</DropdownItem>
                <DropdownItem on:click={renewKeycode}>6</DropdownItem>
                <DropdownItem on:click={renewKeycode}>7</DropdownItem>
                <DropdownItem on:click={renewKeycode}>8</DropdownItem>
                <DropdownItem on:click={renewKeycode}>9</DropdownItem>
                <DropdownDivider/>
                <DropdownItem on:click={renewKeycode}>Enter</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Escape</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Backspace</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Tab</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Space</DropdownItem>
                <DropdownDivider/>
                <DropdownItem on:click={renewKeycode}>-_</DropdownItem>
                <DropdownItem on:click={renewKeycode}>=+</DropdownItem>
                <DropdownItem on:click={renewKeycode}>&#091;&#123;</DropdownItem>
                <DropdownItem on:click={renewKeycode}>&#093;&#125;</DropdownItem>
                <DropdownItem on:click={renewKeycode}>&#092;&#124;</DropdownItem>
                <DropdownItem on:click={renewKeycode}>;:</DropdownItem>
                <DropdownItem on:click={renewKeycode}>&#039;&quot;</DropdownItem>
                <DropdownItem on:click={renewKeycode}>`~</DropdownItem>
                <DropdownItem on:click={renewKeycode}>,&lt;</DropdownItem>
                <DropdownItem on:click={renewKeycode}>.&gt;</DropdownItem>
                <DropdownItem on:click={renewKeycode}>/?</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Caps Lock</DropdownItem>
                <DropdownDivider/>
                <DropdownItem on:click={renewKeycode}>F1</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F2</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F3</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F4</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F5</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F6</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F7</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F8</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F9</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F10</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F11</DropdownItem>
                <DropdownItem on:click={renewKeycode}>F12</DropdownItem>
                <DropdownDivider/>
                <DropdownItem on:click={renewKeycode}>Print Screen</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Scroll Lock</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Pause</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Insert</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Home</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Page Up</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Delete</DropdownItem>
                <DropdownItem on:click={renewKeycode}>End</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Page Down</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Right</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Left</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Down</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Up</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Num Lock</DropdownItem>
                <DropdownItem on:click={renewKeycode}>カタカナ ひらがな</DropdownItem>
                <DropdownItem on:click={renewKeycode}>&yen;</DropdownItem>
                <DropdownItem on:click={renewKeycode}>変換</DropdownItem>
                <DropdownItem on:click={renewKeycode}>無変換</DropdownItem>
                <DropdownDivider/>
                <DropdownItem on:click={renewKeycode}>Mouse Left Button</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Mouse Right Button</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Mouse Wheel</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Reset</DropdownItem>
                <DropdownItem on:click={renewKeycode}>XFader Cut Ch.1(Phono/Line In)</DropdownItem>
                <DropdownItem on:click={renewKeycode}>XFader Cut Ch.2(USB)</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Master Gain Up</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Master Gain Down</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Upper</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Line Phono Switch</DropdownItem>
                <DropdownItem on:click={renewKeycode}>Layout Switch</DropdownItem>
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
