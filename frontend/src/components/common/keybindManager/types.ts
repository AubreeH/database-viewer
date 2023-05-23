export interface IKeybindManagerContext {
    subscribe: (keybind: IKeybind) => void,
    unsubscribe: (keybind: IKeybind) => void,
}

export interface IKeybind {
    ctrl?: boolean,
    alt?: boolean,
    shift?: boolean,
    name: string,
    key: string,
    handler: (e: KeyboardEvent) => void
}

export enum Binary {
    Off,
    On,
}

export const BinaryValueToBooleanMap = {
    [Binary.Off]: false,
    [Binary.On]: true,
}

export type KeybindStore = { [key in Binary]?: { [key in Binary]?: { [key in Binary]?: { [key: string]: IKeybind[] }}}}