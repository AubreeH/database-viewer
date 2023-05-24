export interface IOption {
    display: Exclude<string, "">,
    value: Exclude<string, "">,
    title?: string
}

export type Options = IOption[]