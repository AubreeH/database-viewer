import type { SvelteComponentTyped } from "svelte";
import { writable } from "svelte/store";
import type { IDisplayComponent } from "../types/componentTypes";

export interface IModalType<T extends SvelteComponentTyped> extends IDisplayComponent<T> {
    title: string
    open?: boolean
}

export type ModalStoreType = IModalType<any>

const ModalStore = writable<ModalStoreType>()

export function openModal<T extends SvelteComponentTyped>(modal: IModalType<T>) {
    ModalStore.set({...modal, open: true})
}

export function closeModal() {
    ModalStore.update((val) => ({...val, open: false}))
}

export default ModalStore