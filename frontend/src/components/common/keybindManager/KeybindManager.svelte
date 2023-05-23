<script lang="ts">
    import { Binary } from "./types";
    import type {
        IKeybind,
        IKeybindManagerContext,
        KeybindStore,
    } from "./types";
    import { onDestroy, onMount, setContext } from "svelte";

    let keybindStore: KeybindStore = {};

    setContext("keybind-manager", {
        subscribe: handleSubscribeKeybind,
        unsubscribe: handleUnsubscribeKeybind,
    } as IKeybindManagerContext);

    function handleKeyPress(e: KeyboardEvent) {
        const alt = getBinaryValue(e.altKey);
        const ctrl = getBinaryValue(e.ctrlKey);
        const shift = getBinaryValue(e.shiftKey);

        if (testBranch(alt, ctrl, shift, e.key)) {
            const keybinds = keybindStore[alt][ctrl][shift][e.key];
            keybinds.forEach((keybind) => {
                keybind.handler(e);
            });
        }
    }

    function handleSubscribeKeybind(keybind: IKeybind) {
        const alt = getBinaryValue(keybind.alt);
        const ctrl = getBinaryValue(keybind.ctrl);
        const shift = getBinaryValue(keybind.shift);

        growBranch(alt, ctrl, shift, keybind.key);

        keybindStore[alt][ctrl][shift][keybind.key].push(keybind);
    }

    function handleUnsubscribeKeybind(keybind: IKeybind) {
        const alt = getBinaryValue(keybind.alt);
        const ctrl = getBinaryValue(keybind.ctrl);
        const shift = getBinaryValue(keybind.shift);

        if (testBranch(alt, ctrl, shift, keybind.key)) {
            let newKeybinds = [...keybindStore[alt][ctrl][shift][keybind.key]];
            newKeybinds = newKeybinds.filter((k) => k.name !== keybind.name);
            keybindStore[alt][ctrl][shift][keybind.key] = newKeybinds;
        }
    }

    function getBinaryValue(value: boolean) {
        return value ? Binary.On : Binary.Off;
    }

    function growBranch(alt: Binary, ctrl: Binary, shift: Binary, key: string) {
        if (!keybindStore[alt]) {
            keybindStore[alt] = {};
        }
        if (!keybindStore[alt][ctrl]) {
            keybindStore[alt][ctrl] = {};
        }
        if (!keybindStore[alt][ctrl][shift]) {
            keybindStore[alt][ctrl][shift] = {};
        }
        if (!Array.isArray(keybindStore[alt][ctrl][shift][key])) {
            keybindStore[alt][ctrl][shift][key] = [];
        }
    }

    function testBranch(alt: Binary, ctrl: Binary, shift: Binary, key: string) {
        if (!keybindStore[alt]) {
            return false;
        }
        if (!keybindStore[alt][ctrl]) {
            return false;
        }
        if (!keybindStore[alt][ctrl][shift]) {
            return false;
        }
        if (!Array.isArray(keybindStore[alt][ctrl][shift][key])) {
            return false;
        }

        return true;
    }

    onMount(() => {
        document.addEventListener("keypress", handleKeyPress);
    });

    onDestroy(() => {
        document.removeEventListener("keypress", handleKeyPress);
    });
</script>

<slot />
