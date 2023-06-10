import type { EThemes, ThemeValueObject } from "./themes"

export enum EDatabaseDrivers {
    MariaDB,
    MySql,
    PostgreSQL,
    SQLite,
}

export const DatabaseDriverStringMap: Map<EDatabaseDrivers, {value: string, display: string}> = new Map<EDatabaseDrivers, {value: string, display: string}>([
    [EDatabaseDrivers.MariaDB, {value: "mysql???mariadb", display: "MariaDB"}],
    [EDatabaseDrivers.MySql, {value: "mysql", display: "MySQL"}],
    [EDatabaseDrivers.PostgreSQL, {value: "postgres", display: "PostgreSQL"}],
    [EDatabaseDrivers.SQLite, {value: "sqlite", display: "SQLite"}],
])

export const DatabaseDriverIconMap: Map<EDatabaseDrivers, ThemeValueObject> = new Map<EDatabaseDrivers, ThemeValueObject>([
    [EDatabaseDrivers.MariaDB, {
        dark: "images/svg/driver icons/mariadb/white_vertical.svg",
        light: "images/svg/driver icons/mariadb/black_vertical.svg"
    }],
    [EDatabaseDrivers.MySql, {
        dark: "images/svg/driver icons/mysql/mysql-icon.svg",
        light: "images/svg/driver icons/mysql/mysql-icon.svg"
    }],
    [EDatabaseDrivers.PostgreSQL, {
        dark: "images/svg/driver icons/postgresql/postgres_elephant.svg",
        light: "images/svg/driver icons/postgresql/postgres_elephant.svg"
    }],
    [EDatabaseDrivers.SQLite, {
        dark: "images/svg/driver icons/sqlite/sqlite-icon.svg",
        light: "images/svg/driver icons/sqlite/sqlite-icon.svg"
    }],
]
)