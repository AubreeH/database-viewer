import { getContext } from "svelte";
import type { INotificationsContext } from "./types";

export function getNotificationContext(): INotificationsContext {
    return getContext("notifications");
}