export interface INotificationsContext {
    sendNotification(notification: INotification): void
    sendInfoNotification(notification: ITypedNotification): void
    sendWarningNotification(notification: ITypedNotification): void
    sendSuccessNotification(notification: ITypedNotification): void
    sendErrorNotification(notification: ITypedNotification): void
    sendFatalNotification(notification: ITypedNotification): void
}

export interface INotification {
    message: string
    type?: ENotificationType
    icon?: TNotificationIconType
    iconProps?: any
    color?: string
    background?: string
}
export type ITypedNotification = Omit<INotification, "type" | "icon" | "iconProps" | "color">
export type TNotificationIconType = ENotificationType | string
export type Notifications = { [key: string]:  INotification}

export enum ENotificationType {
    Info = "info",
    Success = "check_circle",
    Warning = "warning",
    Error = "error",
    Fatal = "dangerous",
}

export const NotificationTypeColourMap = new Map<ENotificationType, {fg: string, bg: string}>([
    [ENotificationType.Info, {
        fg: "var(--info-notification-fg)",
        bg: "var(--info-notification-bg)",
    }],
    [ENotificationType.Success, {
        fg: "var(--success-notification-fg)",
        bg: "var(--success-notification-bg)",
    }],
    [ENotificationType.Warning, {
        fg: "var(--warning-notification-fg)",
        bg: "var(--warning-notification-bg)",
    }],
    [ENotificationType.Error, {
        fg: "var(--error-notification-fg)",
        bg: "var(--error-notification-bg)",
    }],
    [ENotificationType.Fatal, {
        fg: "var(--fatal-notification-fg)",
        bg: "var(--fatal-notification-bg)",
    }],
])