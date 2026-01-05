<script lang="ts">
  import type {PageProps} from "./$types";
  import {PushService} from "$lib/dp/service/PushService.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import {client} from "$lib/dp/client.svelte.ts";

  let {params}: PageProps = $props();

  const pushService = new PushService();
  const userID = $derived(params.user);

  async function register() {
    const result = await pushService.subscribeUser();

    switch (result.status) {
      case "denied":
        toastController.trigger({
          message: "Push notifications are disabled on your browser.",
          background: "preset-filled-error-500"
        });
        return;

      case "not-found":
        toastController.trigger({
          message: "Technical error - service worker registration not found.",
          background: "preset-filled-error-500"
        });
        return;

      case "already-subscribed":
        toastController.trigger({
          message: "You are already subscribed to push notifications.",
          background: "preset-filled-warning-500"
        });
        return;

      case "granted":
        if (result.subscription instanceof PushSubscription) {
          const response = await pushService.sendSubscriptionToServer(result.subscription!, userID);
          if (response && response.id) {
            toastController.trigger({
              message: "Push notifications successfully registered.",
              background: "preset-filled-success-500"
            });
          }
        } else {
          toastController.trigger({
            message: "Technical error - subscription not submitted to server.",
            background: "preset-filled-error-500"
          });
        }
    }
  }

  async function sendTestNotification() {
    await client.send<string>(`/api/webpush/${userID}/notify-me`, {
      method: "POST"
    });
  }
</script>

<button
  class="btn preset-tonal-primary border"
  onclick={register}>
  Register for Push Notifications
</button>

<button
  class="btn preset-tonal-primary border"
  onclick={sendTestNotification}>
  Send Test Notification
</button>

<style>
  button {
    margin: 1rem;
  }
</style>
