import {afterEach, beforeEach, describe, expect, it, vi} from "vitest";
import type {Toast} from "$lib/dp/types/Toast.ts";
import {toastController} from "./ToastController.svelte.js";

describe("ToastController", () => {
  // Mock data
  const mockToast: Toast = {
    id: "test-id",
    message: "Test message",
    background: "bg-primary",
  };

  // Mock crypto.randomUUID
  const mockUUID = "generated-uuid";
  const originalRandomUUID = crypto.randomUUID;

  // Reset the controller before each test
  beforeEach(() => {
    toastController.clear();
    vi.useFakeTimers();
    // @ts-expect-error
    crypto.randomUUID = vi.fn(() => mockUUID);
  });

  afterEach(() => {
    vi.restoreAllMocks();
    // Restore original crypto.randomUUID
    crypto.randomUUID = originalRandomUUID;
  });

  describe("trigger", () => {
    it("adds a toast to the toasts array", () => {
      // Initial state should be empty
      expect(toastController.toastQueue.length).toBe(0);

      // Trigger a toast
      toastController.trigger(mockToast);

      // Should have one toast
      expect(toastController.toastQueue.length).toBe(1);
      expect(toastController.toastQueue[0]).toEqual(mockToast);
    });

    it("automatically removes the toast after the timeout", () => {
      // Trigger a toast
      toastController.trigger(mockToast);

      // Should have one toast
      expect(toastController.toastQueue.length).toBe(1);

      // Fast-forward time
      vi.advanceTimersByTime(5000);

      // Toast should be removed
      expect(toastController.toastQueue.length).toBe(0);
    });

    it("correctly identifies and removes the specific toast by id", () => {
      // Create multiple toasts with different IDs
      const toast1: Toast = {id: "id1", message: "Message 1"};
      const toast2: Toast = {id: "id2", message: "Message 2"};

      // Add both toasts
      toastController.trigger(toast1);
      vi.advanceTimersByTime(1000); // slight time offset between the two
      toastController.trigger(toast2);

      // Should have two toasts
      expect(toastController.toastQueue.length).toBe(2);

      // Fast-forward time to remove the first toast with some padding
      vi.advanceTimersByTime(toastController.toastDuration - 1000);

      // Should have one toast left (the second one)
      expect(toastController.toastQueue.length).toBe(1);
      expect(toastController.toastQueue.at(0))?.toEqual(toast2);
    });

    it("generates a UUID if one is not provided", () => {
      // Create a toast without an ID
      const toastWithoutId: Toast = {message: "Message without ID"};

      // Trigger the toast
      toastController.trigger(toastWithoutId);

      // Should have one toast with a generated ID
      expect(toastController.toastQueue.length).toBe(1);
      expect(toastController.toastQueue[0].id).toBe(mockUUID);
      expect(toastController.toastQueue[0].message).toBe(toastWithoutId.message);

      // Verify crypto.randomUUID was called
      expect(crypto.randomUUID).toHaveBeenCalledTimes(1);
    });

    it("uses the provided ID if one is given", () => {
      // Create a toast with an ID
      const toastWithId: Toast = {id: "provided-id", message: "Message with ID"};

      // Trigger the toast
      toastController.trigger(toastWithId);

      // Should have one toast with the provided ID
      expect(toastController.toastQueue.length).toBe(1);
      expect(toastController.toastQueue[0].id).toBe("provided-id");

      // Verify crypto.randomUUID was not called
      expect(crypto.randomUUID).not.toHaveBeenCalled();
    });
  });

  describe("clear", () => {
    it("removes all toasts from the array", () => {
      // Add multiple toasts
      toastController.trigger({message: "Message 1"});
      toastController.trigger({message: "Message 2"});

      // Should have two toasts
      expect(toastController.toastQueue.length).toBe(2);

      // Clear all toasts
      toastController.clear();

      // Should have no toasts
      expect(toastController.toastQueue.length).toBe(0);
    });
  });
});
