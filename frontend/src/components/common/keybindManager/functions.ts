import { getContext } from "svelte";
import type { IKeybindManagerContext } from "./types";

export function getKeybindManagerContext(): IKeybindManagerContext {
    return getContext("keybind-manager")
}