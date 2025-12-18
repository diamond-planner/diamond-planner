import type {Toast} from "$lib/dp/types/Toast.ts";

class ToastController {
  public toastQueue: Toast[] = $state([]);
  public readonly toastDuration = 5000;

  public trigger(toast: Toast) {
    // Generate a UUID if one is not provided
    const toastWithId = {
      ...toast,
      id: toast.id || crypto.randomUUID(),
    };

    this.toastQueue.push(toastWithId);

    setTimeout(() => {
      const index = this.toastQueue.findIndex((t) => t.id === toastWithId.id);
      if (index !== -1) {
        this.toastQueue.splice(index, 1);
      }
    }, this.toastDuration);
  }

  public clear() {
    this.toastQueue = [];
  }

  public triggerGenericFormSuccessMessage(modelName: string) {
    this.trigger({
      message: `${modelName} data saved successfully.`,
      background: "preset-filled-success-500",
    });
  }

  public triggerGenericFormErrorMessage(modelName: string) {
    this.trigger({
      message: `An error occurred while saving ${modelName} data.`,
      background: "preset-filled-error-500",
    });
  }
}

export const toastController = new ToastController();
