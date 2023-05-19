import {WindowToggleMaximise} from "../../../wailsjs/runtime/runtime"

export function handleTitlebarDoubleClick(e: MouseEvent) {
    if (e.target === e.currentTarget) {
        WindowToggleMaximise()
    }
}