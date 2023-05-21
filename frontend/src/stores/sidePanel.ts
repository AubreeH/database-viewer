import { writable } from "svelte/store";
import type { IDisplayComponent } from "../types/componentTypes";
import type { SvelteComponentTyped } from "svelte";

interface ISidePanelItem<T extends SvelteComponentTyped> extends IDisplayComponent<T> {
    name: string
}

type SidePanelStoreType = ISidePanelItem<any>[]

const SidePanelStore = writable<SidePanelStoreType>([])

export function openNewSidePanel<T extends SvelteComponentTyped>(component: ISidePanelItem<T>) {
    SidePanelStore.update(value => [...value, component])
}

export function goToPreviewSidePanel() {
    SidePanelStore.update(value => value.slice(0, -1))
}

export function resetSidePanel() {
    SidePanelStore.set([])
}

export default SidePanelStore