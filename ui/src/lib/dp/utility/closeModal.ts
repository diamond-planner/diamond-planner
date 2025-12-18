import { browser } from "$app/environment"

/**
 * Hacky solution until I find a smarter way.
 */
export function closeModal() {
  if (browser) {
    const closeButton: HTMLButtonElement | null = document.querySelector("[data-dialog-close]")
    if (closeButton) {
      closeButton.click()
    }
  }
}
