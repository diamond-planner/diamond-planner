import {dev} from "$app/environment";
import {client} from "$lib/dp/client.svelte.ts";
import {PUBLIC_VAPID_PUBLIC_KEY} from "$env/static/public";
import type {PushsubscriptionsCreate, PushsubscriptionsResponse} from "$lib/dp/types/pb-types.ts";
import {Collection} from "$lib/dp/enum/Collection.ts";

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

  public async sendSubscriptionToServer(subscription: PushSubscription, userID: string): Promise<PushsubscriptionsResponse | null> {
    const subscriptionData = this.convertToDatabaseFormat(subscription, userID);

    const response = await client
      .collection(Collection.PushSubscriptions)
      .create<PushsubscriptionsResponse>(subscriptionData);

    return response ?? null;
  }

  protected convertToDatabaseFormat(subscription: PushSubscription, userID: string): PushsubscriptionsCreate {
    const decoder = new TextDecoder("utf-8");
    // TODO: submit this correctly
    const authKey = subscription.getKey("auth");
    const p256dhKey = subscription.getKey("p256dh");

    if (!authKey || !p256dhKey) {
      throw new Error("Failed to extract keys from push subscription.");
    }

    return {
      user: userID,
      endpoint: subscription.endpoint,
      key_auth: decoder.decode(authKey),
      key_p256dh: decoder.decode(p256dhKey),
      encoding: PushManager.supportedContentEncodings.join(),
    };
  }
}


