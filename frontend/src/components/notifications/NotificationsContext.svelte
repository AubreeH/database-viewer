<script lang="ts">
	import { setContext } from "svelte";
	import type { INotification, Notifications, ITypedNotification } from "./types";
	import { NotificationTypeColourMap, ENotificationType } from "./types";
	import Notification from "./Notification.svelte";

	let notifications: Notifications = {};

	setContext("notifications", {
		sendNotification(notification: ITypedNotification) {
			sendNotification(constructTypedNotification(notification, ENotificationType.Info));
		},
		sendInfoNotification(notification: ITypedNotification) {
			sendNotification(constructTypedNotification(notification, ENotificationType.Info, true));
		},
		sendSuccessNotification(notification: ITypedNotification) {
			sendNotification(constructTypedNotification(notification, ENotificationType.Success, true));
		},
		sendWarningNotification(notification: ITypedNotification) {
			sendNotification(constructTypedNotification(notification, ENotificationType.Warning, true));
		},
		sendErrorNotification(notification) {
			sendNotification(constructTypedNotification(notification, ENotificationType.Error, true));
		},
		sendFatalNotification(notification) {
			sendNotification(constructTypedNotification(notification, ENotificationType.Fatal, true));
		},
	});

	function sendNotification(notification: Required<INotification>) {
		const timestamp = new Date().getTime();
		if (notifications) {
			notifications[timestamp] = notification;
		} else {
			notifications = { [timestamp]: notification };
		}

		setTimeout(() => clearNotification(timestamp), 5000);
	}

	function clearNotification(timestamp: number) {
		if (notifications[timestamp]) {
			delete notifications[timestamp];
			notifications = notifications;
		}
	}

	function constructTypedNotification(notification: ITypedNotification, notificationType: ENotificationType, forced?: boolean): Required<INotification> {
		if (forced) {
			return {
				...notification,
				type: notificationType,
				icon: notificationType,
				iconProps: undefined,
				color: NotificationTypeColourMap.get(notificationType).fg,
				background: NotificationTypeColourMap.get(notificationType).bg,
			};
		} else {
			return {
				type: notificationType,
				icon: notificationType,
				iconProps: undefined,
				color: NotificationTypeColourMap.get(notificationType).fg,
				background: NotificationTypeColourMap.get(notificationType).bg,
				...notification,
			};
		}
	}
</script>

<slot />

<div class="notifications">
	{#each Object.entries(notifications) as [timestamp, notification] (timestamp)}
		<Notification
			id="notification-{timestamp}"
			{notification}
			on:click={() => clearNotification(Number(timestamp))}
			on:keypress={() => clearNotification(Number(timestamp))}
		/>
	{/each}
</div>

<style lang="scss">
	.notifications {
		pointer-events: none;
		display: flex;
		flex-direction: column-reverse;
		position: fixed;
		bottom: 0;
		right: 0;
		padding: 0 1em;
	}
</style>
