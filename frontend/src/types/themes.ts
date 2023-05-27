export enum EThemes {
    Dark = "dark",
    Light = "light"
}

export type ThemeValueObject = { [key in EThemes]: string }