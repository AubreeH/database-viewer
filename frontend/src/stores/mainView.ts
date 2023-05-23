import type { SvelteComponentTyped } from "svelte";
import { writable } from "svelte/store";
import type { IDisplayComponent } from "../types/componentTypes";

export interface IMainViewStore {
    tabs: IDisplayComponent<any>[],
    activeTab: number
}

export const MainViewStore = writable<IMainViewStore>({
    tabs: [],
    activeTab: -1,
})

export function openTabInBackground<T extends SvelteComponentTyped>(component: IDisplayComponent<T>) {
    MainViewStore.update(value => ({
        ...value,
        tabs: [...value.tabs, component],
    }))
}

export function openTabFocused<T extends SvelteComponentTyped>(component: IDisplayComponent<T>) {
    MainViewStore.update(value => {
        const newTabs = [...value.tabs]
        let activeTab = value.activeTab + 1
        newTabs.splice(activeTab, 0, component)

        return {
            ...value,
            tabs: newTabs,
            activeTab,
        }
    })
}

export function closeTab(index: number) {
    MainViewStore.update(value => {
        const newTabs = [...value.tabs]
        newTabs.splice(index, 1)
        let activeTab = value.activeTab > newTabs.length - 1 ? newTabs.length - 1 : value.activeTab

        return {
            ...value,
            tabs: newTabs,
            activeTab,
        }
    })
}

export function focusTab(index: number) {
    MainViewStore.update(value => ({...value, activeTab: index}))
}

export function closeFocusedTab() {
    MainViewStore.update(value => {
        if(value.activeTab !== -1) {
            const newTabs = [...value.tabs]
            newTabs.splice(value.activeTab, 1)
            let activeTab = value.activeTab > newTabs.length - 1 ? newTabs.length - 1 : value.activeTab
    
            return {
                ...value,
                tabs: newTabs,
                activeTab,
            }
        }

        return value
    })
}