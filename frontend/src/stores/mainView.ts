import type { SvelteComponentTyped } from "svelte";
import { get, writable } from "svelte/store";
import type { IDisplayComponent } from "../types/componentTypes";

export interface ITab<T extends SvelteComponentTyped = SvelteComponentTyped> extends IDisplayComponent<T> {
    id: string
}

export interface IMainViewStore {
    tabs: ITab[],
    activeTab: number
}

export const MainViewStore = writable<IMainViewStore>({
    tabs: [],
    activeTab: -1,
})

export function openTabInBackground<T extends SvelteComponentTyped>(component: ITab<T>) {
    MainViewStore.update(value => ({
        ...value,
        tabs: [...value.tabs, component],
    }))
}

export function openTabFocused<T extends SvelteComponentTyped>(component: ITab<T>) {
    if (!component.id) {
        return
    }
    
    const storeValue = get(MainViewStore)
    if (Array.isArray(storeValue.tabs)) {
        let activeTab = storeValue.tabs[storeValue.activeTab]
        let firstFindIndex = undefined
        const existingTabIndex = storeValue.tabs.findIndex((t, i) => {
            if (t.id === component.id) {
                if (activeTab.id !== component.id) {
                    return true
                } else if(firstFindIndex === undefined) {
                    firstFindIndex = i
                    return false
                } else if (i > storeValue.activeTab) {
                    return true
                }
            }

            return false
        })
        if (existingTabIndex !== -1) {
            focusTab(existingTabIndex)
            return
        } else if (firstFindIndex !== undefined && firstFindIndex !== -1) {
            focusTab(firstFindIndex)
            return
        }
    }

    openNewTabFocused(component)
}

export function openNewTabFocused<T extends SvelteComponentTyped>(component: ITab<T>) {
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
    MainViewStore.update(value => {
        if (index < 0) {
            return value
        }

        if (index >= value?.tabs?.length ?? 0) {
            return value
        }

        return ({...value, activeTab: index})
    })
}

export function closeFocusedTab() {
    MainViewStore.update(value => {
        if(value.activeTab !== -1) {
            if (!value.tabs.length) {
                return {
                    ...value,
                    tabs: [],
                    activeTab: -1,
                }
            }

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