<script lang="ts">
  import {PushService, type PushSubscriptionResult} from "$lib/dp/service/PushService.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import Switch from "$lib/dp/components/formElements/Switch.svelte";
  import {authSettings} from "$lib/dp/client.svelte.ts";

  interface Props {
    userID: string;
  }

  let {userID}: Props = $props();
  const model = $derived(authSettings.record);

  const pushService = new PushService();
  let pushEnabled = $state(await pushService.isUserSubscribed());

  async function register(): Promise<PushSubscriptionResult> {
    const result = await pushService.subscribeUser();

    switch (result.status) {
      case "denied":
        toastController.trigger({
          message: "Push notifications are disabled on your browser.",
          background: "preset-filled-error-500"
        });
        return result;

      case "not-found":
        toastController.trigger({
          message: "Technical error - service worker registration not found.",
          background: "preset-filled-error-500"
        });
        return result;

      case "already-subscribed":
        toastController.trigger({
          message: "You are already subscribed to push notifications.",
          background: "preset-filled-warning-500"
        });
        return result;

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
        return result;
    }
  }

  async function onPushSettingChange(e: { checked: boolean }) {
    if (e.checked) {
      const result = await register();
      if (result.status === "denied" || result.status === "not-found") {
        pushEnabled = false;
      }
    } else {
      const unsub = await pushService.unsubscribeUser();

      let serverResult: boolean | undefined;
      if (unsub) {
        serverResult = await pushService.deleteSubscriptionFromServer(unsub);
      }

      if (unsub && serverResult) {
        toastController.trigger({
          message: "Push notifications successfully unregistered.",
          background: "preset-filled-success-500"
        });
      } else {
        toastController.trigger({
          message: "Technical error - push notification unregistration failed.",
          background: "preset-filled-error-500"
        });
      }
    }
  }
</script>

{#if model}
  <header class="card-header">
    <h2 class="h3">
      Push Notifications
    </h2>
  </header>

  <div class="card-body">
    <div class="push-label">
      <Switch
        bind:checked={pushEnabled}
        onCheckedChange={onPushSettingChange}
        name="pushEnabled"
      >
        Send Push Notifications
      </Switch>
    </div>

    <p class="push-hint">
      Upon checking this box, your browser will prompt you for confirmation to enable push messages on this device.
    </p>

    <label class="label push-label">
      <span>Test</span>
      <button
        class="btn preset-tonal-primary border"
        onclick={() => pushService.sendTestNotification(model.id)}>
        Send Test Notification
      </button>
    </label>
  </div>
{/if}

<style>
  .push-label {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 4);
  }

  .push-label, .push-hint {
    margin-block-end: calc(var(--spacing) * 4);
  }
</style>
