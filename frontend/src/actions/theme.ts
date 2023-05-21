import type { Unsubscriber } from "svelte/store";
import ViewStore, { type ThemeType } from "../stores/view";

let unsubscriber: Unsubscriber = undefined

function theme(node: HTMLElement) {
    let previousThemeValue: ThemeType = undefined

    if(unsubscriber) {
        unsubscriber()
        unsubscriber = undefined
    }

    unsubscriber = ViewStore.subscribe((store) => {
        if(previousThemeValue && node.classList.contains(previousThemeValue)) {
            node.classList.remove(previousThemeValue)
        }
        previousThemeValue = store.theme
        node.classList.add(previousThemeValue)
    })
}


export default theme