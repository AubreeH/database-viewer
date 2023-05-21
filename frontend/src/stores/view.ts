import { writable } from "svelte/store";

export type ThemeType = "dark" | "light";

export interface IViewStore {
    theme: ThemeType,
}

const ViewStore = writable<IViewStore>({
    theme: "dark"
})

export default ViewStore