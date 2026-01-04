import {dev} from "$app/environment";
import {client} from "$lib/dp/client.svelte.ts";
import {PUBLIC_VAPID_PUBLIC_KEY} from "$env/static/public";

type PushSubscriptionResult = {
  subscription: PushSubscription | null
  status: "granted" | "denied" | "already-subscribed" | "not-found"
}

export class PushService {
  private VAPID_PUBLIC_KEY = PUBLIC_VAPID_PUBLIC_KEY;

  public async subscribeUser(): Promise<PushSubscriptionResult> {
    const result = await Notification.requestPermission();
    if (result === 'denied') {
      console.error('The user explicitly denied the permission request.');
      return {
        subscription: null,
        status: 'denied'
      };
    }
    if (result === 'granted' && dev) {
      console.info('The user accepted the permission request.');
    }
    const registration = await navigator.serviceWorker.getRegistration();

    if (!registration) {
      console.error('Service worker registration not found.');
      return {
        subscription: null,
        status: 'not-found'
      };
    }

    const subscribed = await registration.pushManager.getSubscription();
    if (subscribed) {
      console.info('User is already subscribed.');
      return {
        subscription: subscribed,
        status: 'already-subscribed'
      };
    }
    const subscription = await registration.pushManager.subscribe({
      userVisibleOnly: true,
      applicationServerKey: this.VAPID_PUBLIC_KEY
    });

    return {
      subscription: subscription,
      status: 'granted'
    };
  }

  public async sendSubscriptionToServer(subscription: PushSubscription) {
    const response = await client.send("/api/push/subscribe", {
      method: "POST",
      body: subscription.toJSON()
    });

    if (!response.ok) {
      throw new Error('Bad status code from server.');
    }

    return response.json();
  }
}


