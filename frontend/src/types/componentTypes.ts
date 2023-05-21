import type { ComponentProps, ComponentType, SvelteComponentTyped } from "svelte";

export interface IDisplayComponent<T extends SvelteComponentTyped> {
    name?: string,
    component: ComponentType<T>,
    props?: ComponentProps<T>
}